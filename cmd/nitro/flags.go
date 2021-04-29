package main

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/dcaiafa/nitro"
)

type Flag struct {
	Name     string
	Desc     string
	Sys      bool
	Required bool
	Value    interface{}
	Set      bool
}

func (f *Flag) FullName() string {
	if f.Sys {
		return "-" + f.Name
	} else {
		return "--" + f.Name
	}
}

type Flags struct {
	Help bool

	flags map[string]*Flag
}

func NewFlags() *Flags {
	return &Flags{
		flags: make(map[string]*Flag),
	}
}

func (f *Flags) Print(w io.Writer) {
	flags := f.GetFlags()
	maxWidth := 0
	for _, f := range flags {
		l := len(f.FullName())
		if l > maxWidth {
			maxWidth = l
		}
	}
	fmtSpec := fmt.Sprintf("  %%-%ds %%s\n", maxWidth)
	for _, f := range flags {
		fmt.Fprintf(w, fmtSpec, f.FullName(), f.Desc)
	}
}

func (f *Flags) GetFlags() []*Flag {
	flags := make([]*Flag, 0, len(f.flags))
	for _, flag := range f.flags {
		flags = append(flags, flag)
	}
	sort.Slice(flags, func(i, j int) bool {
		return flags[i].Name < flags[j].Name
	})
	return flags
}

func (f *Flags) AddFlag(flag *Flag) *Flag {
	f.flags[flag.Name] = flag
	return flag
}

func (f *Flags) AddFlagsFromMetadata(md *nitro.Metadata) error {
	for _, param := range md.Params {
		flag := &Flag{
			Name: param.Name,
			Desc: param.Desc,
		}
		switch param.Type {
		case "bool":
			flag.Value = new(bool)
		case "string":
			flag.Value = new(string)
		case "[]string":
			flag.Value = new([]string)
		case "int":
			flag.Value = new(int64)
		case "float":
			flag.Value = new(float64)
		case "", "any":
			flag.Value = new(interface{})
		default:
			return fmt.Errorf("unsupported param type %q", param.Type)
		}
		f.AddFlag(flag)
	}
	return nil
}

func (f *Flags) GetNitroValues() map[string]nitro.Value {
	values := make(map[string]nitro.Value, len(f.flags))

	for _, flag := range f.flags {
		if !flag.Set {
			continue
		}
		switch v := flag.Value.(type) {
		case *bool:
			values[flag.Name] = nitro.NewBool(*v)
		case *string:
			values[flag.Name] = nitro.NewString(*v)
		case *[]string:
			a := nitro.NewArray()
			for _, s := range *v {
				a.Push(nitro.NewString(s))
			}
			values[flag.Name] = a
		case *int64:
			values[flag.Name] = nitro.NewInt(*v)
		case *float64:
			values[flag.Name] = nitro.NewFloat(*v)
		default:
			panic("unreachable")
		}
	}

	return values
}

func (f *Flags) Parse(args []string) ([]string, error) {
	var err error
	for len(args) > 0 {
		if !IsFlag(args[0]) {
			break
		}
		if args[0] == "-h" || args[0] == "--help" {
			f.Help = true
			args = args[1:]
			continue
		}
		args, err = f.parseFlag(args)
		if err != nil {
			return nil, err
		}
	}
	return args, nil
}

func (f *Flags) parseFlag(args []string) ([]string, error) {
	origFlag := args[0]
	flagName := origFlag
	args = args[1:]

	flagName = flagName[1:]
	if len(flagName) == 0 {
		return nil, fmt.Errorf("invalid flag without name")
	}

	isSys := true
	if flagName[0] == '-' {
		isSys = false
		flagName = flagName[1:]
	}

	hasValue := false
	flagValue := ""
	idxEq := strings.IndexByte(flagName, '=')
	if idxEq != -1 {
		flagValue = flagName[idxEq+1:]
		flagName = flagName[:idxEq]
		hasValue = true
	}

	flag := f.flags[flagName]
	if flag == nil || flag.Sys != isSys {
		return nil, fmt.Errorf("invalid flag %v", origFlag)
	}

	if !hasValue {
		if _, ok := flag.Value.(*bool); ok {
			flagValue = "true"
		} else {
			if len(args) != 0 && !IsFlag(args[0]) {
				flagValue = args[0]
				args = args[1:]
			} else {
				return nil, fmt.Errorf("flag %v requires a value", origFlag)
			}
		}
	}

	err := f.parseFlagValue(flagValue, flag.Value)
	if err != nil {
		return nil, err
	}

	flag.Set = true

	return args, nil
}

func (f *Flags) parseFlagValue(str string, v interface{}) error {
	switch v := v.(type) {
	case *bool, []*bool:
		b, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		switch v := v.(type) {
		case *bool:
			*v = b
		case *[]bool:
			*v = append(*v, b)
		}

	case *string:
		*v = str

	case *[]string:
		*v = append(*v, str)

	case *int64, *[]int64:
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		switch v := v.(type) {
		case *int64:
			*v = i
		case *[]int64:
			*v = append(*v, i)
		}

	case *float64, *[]float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		switch v := v.(type) {
		case *float64:
			*v = f
		case *[]float64:
			*v = append(*v, f)
		}

	default:
		return fmt.Errorf("invalid flag value type %q", reflect.TypeOf(v))
	}
	return nil
}

func IsFlag(arg string) bool {
	return len(arg) != 0 && arg[0] == '-'
}
