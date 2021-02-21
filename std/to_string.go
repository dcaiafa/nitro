package std

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/dcaiafa/nitro"
)

func ToString(ctx context.Context, args []nitro.Value) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errors.New("argument is missing")
	}

	v := args[0]
	f := ""
	if len(args) > 1 {
		if s, ok := args[1].(nitro.String); ok {
			f = string(s)
		} else {
			return nil, errors.New("format argument must be a string")
		}
	}

	switch v := v.(type) {
	case nitro.String:
		return []nitro.Value{v}, nil
	case nitro.Int, nitro.Float, nitro.Bool:
		if f == "" {
			f = "%v"
		} else {
			f = "%" + f
		}
		s := fmt.Sprintf(f, v)
		return []nitro.Value{nitro.String(s)}, nil
	case *nitro.Object:
		s, err := formatObject(v)
		if err != nil {
			return nil, err
		}
		return []nitro.Value{nitro.String(s)}, nil
	default:
		return []nitro.Value{nitro.String("<" + reflect.TypeOf(v).String() + ">")}, nil
	}
}

type objectFormatter struct {
	visited map[nitro.Value]bool
	w       *strings.Builder
	indent  string
}

func formatObject(o *nitro.Object) (string, error) {
	of := &objectFormatter{
		visited: make(map[nitro.Value]bool),
		w:       &strings.Builder{},
	}
	of.format(o)
	return of.w.String(), nil
}

func (of *objectFormatter) format(v nitro.Value) {
	switch v := v.(type) {
	case nitro.String:
		// TODO: escape special characters.
		of.str(`"` + string(v) + `"`)
	case nitro.Int:
		of.str(strconv.FormatInt(int64(v), 10))
	case nitro.Float:
		of.str(strconv.FormatFloat(float64(v), 'g', -1, 64))
	case nitro.Bool:
		if bool(v) {
			of.str("true")
		} else {
			of.str("false")
		}
	case *nitro.Object:
		if of.visited[v] {
			of.str("<cycle>")
			return
		}
		of.visited[v] = true
		of.str("{")
		if of.indent != "" {
			of.str("\n" + of.indent)
		}
		first := true
		v.ForEach(func(k, v nitro.Value) bool {
			if of.indent == "" && !first {
				of.str(", ")
			}
			first = false
			if ks, ok := k.(nitro.String); ok {
				of.str(string(ks))
			} else {
				of.str("[")
				of.format(k)
				of.str("]")
			}
			of.str(": ")
			of.format(v)
			return true
		})
		of.str("}")

	default:
		of.str("<" + reflect.TypeOf(v).String() + ">")
	}
}

func (of *objectFormatter) str(s string) {
	of.w.WriteString(s)
}
