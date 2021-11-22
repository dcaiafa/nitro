package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
	fatihcolor "github.com/fatih/color"
)

type printMods struct {
	NoNL  bool
	Color *fatihcolor.Color
}

type nonlMod struct{}

func (m *nonlMod) String() string { return "<nonl>" }
func (m *nonlMod) Type() string   { return "nonl" }

func (m *nonlMod) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, errors.New("nonl does not support this operation")
}

func getPrintMods(args []nitro.Value) (printMods, []nitro.Value) {
	hasMods := false
Loop:
	for _, arg := range args {
		switch arg.(type) {
		case *nonlMod, *colorMod:
			hasMods = true
			break Loop
		}
	}

	if !hasMods {
		// Shortcut: there are no mod arguments.
		return printMods{}, args
	}

	// Compute mods while simultaneously compiling the list of non-mod args.
	var mods printMods
	newArgs := make([]nitro.Value, 0, len(args))
	for _, arg := range args {
		switch m := arg.(type) {
		case *nonlMod:
			mods.NoNL = true
		case *colorMod:
			mods.Color = m.color
		default:
			newArgs = append(newArgs, arg)
		}
	}

	return mods, newArgs
}

var errNoNLUsage = errors.New(
	`invalid usage. Expected nonl()`)

func nonl(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 0 {
		return nil, errNoNLUsage
	}

	return []nitro.Value{&nonlMod{}}, nil
}

type colorMod struct {
	color *fatihcolor.Color
}

func (m *colorMod) String() string { return "<color>" }
func (m *colorMod) Type() string   { return "color" }

func (m *colorMod) EvalOp(op nitro.Op, operand nitro.Value) (nitro.Value, error) {
	return nil, errors.New("color does not support this operation")
}

var errColorUsage = errors.New(
	`invalid usage. Expected color(string+)`)

func color(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) == 0 {
		return nil, errColorUsage
	}

	attribs := make([]fatihcolor.Attribute, 0, len(args))

	for _, arg := range args {
		argStr, ok := arg.(nitro.String)
		if !ok {
			return nil, errColorUsage
		}

		attrib, err := colorAttribute(argStr.String())
		if err != nil {
			return nil, err
		}

		attribs = append(attribs, attrib)
	}

	c := &colorMod{
		color: fatihcolor.New(attribs...),
	}

	return []nitro.Value{c}, nil
}

func colorAttribute(v string) (attrib fatihcolor.Attribute, err error) {
	switch v {
	case "reset":
		attrib = fatihcolor.Reset
	case "bold":
		attrib = fatihcolor.Bold
	case "faint":
		attrib = fatihcolor.Faint
	case "italic":
		attrib = fatihcolor.Italic
	case "underline":
		attrib = fatihcolor.Underline
	case "blinkslow":
		attrib = fatihcolor.BlinkSlow
	case "blinkrapid":
		attrib = fatihcolor.BlinkRapid
	case "reversevideo":
		attrib = fatihcolor.ReverseVideo
	case "concealed":
		attrib = fatihcolor.Concealed
	case "crossedout":
		attrib = fatihcolor.CrossedOut
	case "black":
		attrib = fatihcolor.FgBlack
	case "red":
		attrib = fatihcolor.FgRed
	case "green":
		attrib = fatihcolor.FgGreen
	case "yellow":
		attrib = fatihcolor.FgYellow
	case "blue":
		attrib = fatihcolor.FgBlue
	case "magenta":
		attrib = fatihcolor.FgMagenta
	case "cyan":
		attrib = fatihcolor.FgCyan
	case "white":
		attrib = fatihcolor.FgWhite
	case "hiblack":
		attrib = fatihcolor.FgHiBlack
	case "hired":
		attrib = fatihcolor.FgHiRed
	case "higreen":
		attrib = fatihcolor.FgHiGreen
	case "hiyellow":
		attrib = fatihcolor.FgHiYellow
	case "hiblue":
		attrib = fatihcolor.FgHiBlue
	case "himagenta":
		attrib = fatihcolor.FgHiMagenta
	case "hicyan":
		attrib = fatihcolor.FgHiCyan
	case "hiwhite":
		attrib = fatihcolor.FgHiWhite
	case "bgblack":
		attrib = fatihcolor.BgBlack
	case "bgred":
		attrib = fatihcolor.BgRed
	case "bggreen":
		attrib = fatihcolor.BgGreen
	case "bgyellow":
		attrib = fatihcolor.BgYellow
	case "bgblue":
		attrib = fatihcolor.BgBlue
	case "bgmagenta":
		attrib = fatihcolor.BgMagenta
	case "bgcyan":
		attrib = fatihcolor.BgCyan
	case "bgwhite":
		attrib = fatihcolor.BgWhite
	case "bghiblack":
		attrib = fatihcolor.BgHiBlack
	case "bghired":
		attrib = fatihcolor.BgHiRed
	case "bghigreen":
		attrib = fatihcolor.BgHiGreen
	case "bghiyellow":
		attrib = fatihcolor.BgHiYellow
	case "bghiblue":
		attrib = fatihcolor.BgHiBlue
	case "bghimagenta":
		attrib = fatihcolor.BgHiMagenta
	case "bghicyan":
		attrib = fatihcolor.BgHiCyan
	case "bghiwhite":
		attrib = fatihcolor.BgHiWhite
	default:
		return fatihcolor.Reset, fmt.Errorf("invalid color attribute %v", v)
	}
	return attrib, nil
}

func print(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return basePrint(Stdout(vm), vm, args, nRet)
}

func basePrint(out io.Writer, m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	mods, args := getPrintMods(args)

	fprint := fmt.Fprintln
	if mods.NoNL {
		fprint = fmt.Fprint
	}
	if mods.Color != nil {
		if mods.NoNL {
			fprint = mods.Color.Fprint
		} else {
			fprint = mods.Color.Fprintln
		}
	}

	if len(args) == 1 {
		switch arg := args[0].(type) {
		case nitro.Iterator:
			first := true
			for {
				if mods.NoNL && !first {
					_, err := fprint(out, " ")
					if err != nil {
						return nil, err
					}
				}
				nret := arg.IterNRet()
				v, err := m.IterNext(arg, nret)
				if err != nil {
					return nil, err
				}
				if v == nil {
					break
				}
				iargs := valuesToInterface(v)
				_, err = fprint(out, iargs...)
				if err != nil {
					return nil, err
				}
				first = false
			}
			return nil, nil

		case io.Reader:
			_, err := io.Copy(out, arg)
			if err != nil {
				return nil, err
			}
			if !mods.NoNL {
				fmt.Println("")
			}

			return nil, nil
		}
	}

	iargs := valuesToInterface(args)
	fprint(out, iargs...)

	return nil, nil
}

func printf(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return basePrintf(Stdout(vm), vm, args, nRet)
}

func basePrintf(out io.Writer, m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	mods, args := getPrintMods(args)

	fprintf := fmt.Fprintf
	if mods.Color != nil {
		fprintf = mods.Color.Fprintf
	}

	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	format, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	if !mods.NoNL {
		format = format + "\n"
	}

	iargs := valuesToInterface(args[1:])
	fprintf(out, format, iargs...)

	return nil, nil
}

func log(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return basePrint(Stderr(vm), vm, args, nRet)
}

func logf(vm *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	return basePrintf(Stderr(vm), vm, args, nRet)
}

func sprintf(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	format, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	iargs := valuesToInterface(args[1:])
	res := fmt.Sprintf(format, iargs...)

	return []nitro.Value{nitro.NewString(res)}, nil
}

type printTableOptions struct {
	AlignRight bool  `nitro:"alignright"`
	MinWidth   int64 `nitro:"minwidth"`
	Padding    int64 `nitro:"padding"`
	PadChar    int64 `nitro:"padchar"`
}

var printTableConv core.Value2Structer

var errPrintTableUsage = errors.New(
	`invalid usage. Expected print_table(iter, map?)`)

func printTable(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errPrintTableUsage
	}

	inIter, err := nitro.MakeIterator(m, args[0])
	if err != nil {
		return nil, errPrintTableUsage
	}
	defer m.IterClose(inIter)

	opts := printTableOptions{
		Padding: -1,
	}

	if len(args) == 2 {
		err = printTableConv.Convert(args[1], &opts)
		if err != nil {
			return nil, err
		}
	}

	var flags uint
	if opts.AlignRight {
		flags = flags | tabwriter.AlignRight
	}

	if opts.Padding == -1 {
		opts.Padding = 1
	}
	if opts.PadChar == 0 {
		opts.PadChar = ' '
	}

	tabw := tabwriter.NewWriter(
		Stdout(m),
		int(opts.MinWidth),
		0, /*tabwidth*/
		int(opts.Padding),
		byte(opts.PadChar),
		flags)

	defer tabw.Flush()

	buf := bytes.Buffer{}
	writeRecord := func(vs []nitro.Value) error {
		buf.Reset()
		for _, v := range vs {
			// TODO: replace \t in value.
			buf.WriteString(v.String())
			buf.WriteByte('\t')
		}
		buf.WriteByte('\n')
		_, err := tabw.Write(buf.Bytes())
		return err
	}

	first := true
	var headers []nitro.Value
	var values []nitro.Value
	for {
		vs, err := m.IterNext(inIter, 1)
		if err != nil {
			return nil, err
		}
		if vs == nil {
			break
		}

		rec := vs[0]

		if mrec, ok := rec.(*nitro.Object); ok {
			if first {
				headers = make([]nitro.Value, 0, mrec.Len())
				values = make([]nitro.Value, mrec.Len())
				err = nil
				mrec.ForEach(func(k, v nitro.Value) bool {
					kstr, ok := k.(nitro.String)
					if !ok {
						err = fmt.Errorf(
							"only string keys are supported, but map has key type %v",
							nitro.TypeName(k))
						return false
					}
					headers = append(headers, kstr)
					return true
				})
				writeRecord(headers)
			}
			for i, k := range headers {
				v, ok := mrec.Get(k)
				if ok {
					if vstr, ok := v.(nitro.String); ok {
						values[i] = vstr
					} else if str, ok := v.(fmt.Stringer); ok {
						values[i] = nitro.NewString(str.String())
					} else {
						values[i] = nitro.NewString("")
					}
				} else {
					values[i] = nitro.NewString("")
				}
			}
			writeRecord(values)
		} else if _, ok := rec.(*nitro.Array); ok {
			panic("not impl")
		} else {
			return nil, fmt.Errorf(
				"expected iterator values of array or map; received %v",
				nitro.TypeName(rec))
		}

		first = false
	}

	err = tabw.Flush()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func valuesToInterface(values []nitro.Value) []interface{} {
	ivalues := make([]interface{}, len(values))
	for i, v := range values {
		switch v := v.(type) {
		case nitro.Int:
			ivalues[i] = v.Int64()
		case nitro.Float:
			ivalues[i] = v.Float64()
		case nitro.String:
			ivalues[i] = v.String()
		default:
			ivalues[i] = v
		}
	}
	return ivalues
}
