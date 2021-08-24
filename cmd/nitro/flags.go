package main

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/dcaiafa/nitro"
)

type Flag struct {
	Name     string
	Desc     string
	Required bool
	Value    interface{}
	IsList   bool
	Set      bool
	Pos      *int
}

func (f *Flag) FullName() string {
	if f.Pos == nil {
		prefix := "-"
		if len(f.Name) > 1 {
			prefix = "--"
		}
		return prefix + f.Name
	} else {
		if f.Required {
			return fmt.Sprintf("<%v>", f.Name)
		} else {
			return fmt.Sprintf("[<%v>]", f.Name)
		}
	}
}

type Flags struct {
	Help bool

	flags   map[string]*Flag
	posArgs []*Flag
}

func NewFlags() *Flags {
	return &Flags{
		flags: make(map[string]*Flag),
	}
}

func (f *Flags) HasArgs() bool {
	return len(f.posArgs) != 0
}

func (f *Flags) PrintArgs(w io.Writer) {
	tabw := tabwriter.NewWriter(
		w,
		3,   /*minwidth*/
		0,   /*tabwidth*/
		2,   /*padding*/
		' ', /*padchar*/
		0,   /*flags*/
	)

	args := f.posArgs
	for _, arg := range args {
		fmt.Fprintf(tabw, "  %s\t%s\n", arg.FullName(), arg.Desc)
	}
	tabw.Flush()
}

func (f *Flags) GetArgs() []*Flag {
	return f.posArgs
}

func (f *Flags) HasFlags() bool {
	return len(f.flags) != 0
}

func (f *Flags) PrintFlags(w io.Writer) {
	tabw := tabwriter.NewWriter(
		w,
		3,   /*minwidth*/
		0,   /*tabwidth*/
		2,   /*padding*/
		' ', /*padchar*/
		0,   /*flags*/
	)

	flags := f.GetFlags()
	for _, f := range flags {
		fmt.Fprintf(tabw, "  %s\t%s\n", f.FullName(), f.Desc)
	}
	tabw.Flush()
}

func (f *Flags) GetFlags() []*Flag {
	flags := make([]*Flag, 0, len(f.flags))
	for _, flag := range f.flags {
		if flag.Pos == nil {
			flags = append(flags, flag)
		}
	}
	sort.Slice(flags, func(i, j int) bool {
		return flags[i].Name < flags[j].Name
	})
	return flags
}

func (f *Flags) AddFlag(flag *Flag) *Flag {
	if flag.Pos == nil {
		f.flags[flag.Name] = flag
	} else {
		f.posArgs = append(f.posArgs, flag)
	}
	return flag
}

func (f *Flags) AddFlagsFromMetadata(md *nitro.Metadata) error {
	posIndex := 0
	for _, param := range md.Params {
		flag := &Flag{
			Name:     param.Name,
			Desc:     param.Desc,
			Required: param.Required,
		}
		if param.Positional {
			flag.Pos = new(int)
			*flag.Pos = posIndex
			posIndex++
		}

		switch param.Type {
		case "bool":
			flag.Value = new(bool)
		case "", "string":
			flag.Value = new(string)
		case "[]string":
			flag.IsList = true
			flag.Value = new([]string)
		case "int":
			flag.Value = new(int64)
		case "float":
			flag.Value = new(float64)
		default:
			return fmt.Errorf("unsupported param type %q", param.Type)
		}

		f.AddFlag(flag)
	}
	sort.Slice(f.posArgs, func(i, j int) bool {
		return *f.posArgs[i].Pos < *f.posArgs[j].Pos
	})

	return nil
}

func (f *Flags) GetNitroValues() map[string]nitro.Value {
	values := make(map[string]nitro.Value, len(f.flags))
	processFlag := func(flag *Flag) {
		if !flag.Set {
			return
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

	for _, flag := range f.flags {
		processFlag(flag)
	}
	for _, posArg := range f.posArgs {
		processFlag(posArg)
	}

	return values
}

func (f *Flags) Parse(args []string) ([]string, error) {
	var err error

	posIndex := 0
	for len(args) > 0 {
		if IsFlag(args[0]) {
			if args[0] == "-h" || args[0] == "--help" {
				f.Help = true
				args = args[1:]
				continue
			}
			args, err = f.parseFlag(args)
			if err != nil {
				return nil, err
			}
		} else if posIndex < len(f.posArgs) {
			posArg := f.posArgs[posIndex]
			if !posArg.IsList {
				// Don't advance/consume argument if it is a list because all
				// remaining positional arguments, if available, will be added to it
				// (validation will have ensured that only the last argument is a
				// list).
				posIndex++
			}
			err = f.parseFlagValue(args[0], posArg.Value)
			if err != nil {
				return nil, fmt.Errorf(
					"invalid value for positional argument %v (%v): %w",
					posArg.Name, *posArg.Pos, err)
			}
			posArg.Set = true
			args = args[1:]
		} else {
			break
		}
	}

	for _, flag := range f.flags {
		if flag.Required && !flag.Set {
			return nil, fmt.Errorf(
				"missing %v", flag.FullName())
		}
	}
	for _, flag := range f.posArgs {
		if flag.Required && !flag.Set {
			return nil, fmt.Errorf(
				"missing %v", flag.FullName())
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

	// Single dash = short. E.g. -x
	// Double dash = long. E.g. --file
	short := true
	if flagName[0] == '-' {
		short = false
		flagName = flagName[1:]
	}

	if !short && len(flagName) <= 1 {
		return nil, fmt.Errorf("unknown flag %v", origFlag)
	}

	hasValue := false
	flagValue := ""

	// Both short and long formats can have a value specified via '='.
	// E.g. --foo=bar or -y=123

	idxEq := strings.IndexByte(flagName, '=')
	if idxEq != -1 {
		flagValue = flagName[idxEq+1:]
		flagName = flagName[:idxEq]
		hasValue = true
	}

	if short {
		// Multiple short flags can be specified in a single arg, and it can even
		// include a value for the last flag. All of the following lines are
		// equivalent.
		// -x -y -z=123
		// -xy -z 123
		// -xyz=123
		// -xyz123

		for i := 0; i < len(flagName); i++ {
			shortFlagName := flagName[i : i+1]
			flag := f.flags[shortFlagName]
			if flag == nil {
				return nil, fmt.Errorf("invalid flag -%v", shortFlagName)
			}

			if i < len(flagName)-1 {
				if _, ok := flag.Value.(*bool); !ok {
					hasValue = true
					flagValue = flagName[i+1:]
					i = len(flagName)
				}
			}

			if !hasValue {
				if _, ok := flag.Value.(*bool); ok {
					flagValue = "true"
				} else {
					if len(args) != 0 && !IsFlag(args[0]) {
						flagValue = args[0]
						args = args[1:]
					} else {
						return nil, fmt.Errorf("%v requires a value", flag.FullName())
					}
				}
			}

			err := f.parseFlagValue(flagValue, flag.Value)
			if err != nil {
				return nil, err
			}

			flag.Set = true
		}
	} else {
		flag := f.flags[flagName]
		if flag == nil {
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
	}

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
