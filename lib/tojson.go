package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dcaiafa/nitro"
)

func fnToJSON(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	indent := ""
	var err error
	if len(args) >= 2 {
		var opts *nitro.Object
		opts, err = getObjectArg(args, 1)
		if err != nil {
			return nil, err
		}
		opts.ForEach(func(opt, optValue nitro.Value) bool {
			optStr, ok := opt.(nitro.String)
			if !ok {
				err = fmt.Errorf("invalid option %v", opt)
				return false
			}
			switch optStr.String() {
			case "indent":
				val, ok := optValue.(nitro.String)
				if !ok {
					err = fmt.Errorf("option 'indent' must be a String")
					return false
				}
				indent = val.String()

			default:
				err = fmt.Errorf("invalid option %v", optStr.String())
				return false
			}
			return true
		})
	}
	if err != nil {
		return nil, err
	}

	marshaler := jsonMarshaler{}
	err = marshaler.marshal(args[0])
	if err != nil {
		return nil, err
	}

	buf := marshaler.buf.Bytes()

	if indent != "" {
		newBuf := bytes.Buffer{}
		err = json.Indent(&newBuf, buf, "", indent)
		if err != nil {
			return nil, err
		}
		buf = newBuf.Bytes()
	}

	return []nitro.Value{nitro.NewString(string(buf))}, nil
}

type jsonMarshaler struct {
	buf bytes.Buffer
}

func (m *jsonMarshaler) marshal(v nitro.Value) error {
	if v == nil {
		m.buf.WriteString("null")
		return nil
	}

	switch v := v.(type) {
	case nitro.Int:
		b, err := json.Marshal(v.Int64())
		if err != nil {
			return err
		}
		m.buf.Write(b)
	case nitro.Float:
		b, err := json.Marshal(v.Float64())
		if err != nil {
			return err
		}
		m.buf.Write(b)
	case nitro.String:
		b, err := json.Marshal(v.String())
		if err != nil {
			return err
		}
		m.buf.Write(b)
	case nitro.Bool:
		b, err := json.Marshal(v.Bool())
		if err != nil {
			return err
		}
		m.buf.Write(b)
	case *nitro.Array:
		m.buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				m.buf.WriteByte(',')
			}
			err := m.marshal(v.Get(i))
			if err != nil {
				return err
			}
		}
		m.buf.WriteByte(']')
	case *nitro.Object:
		m.buf.WriteByte('{')
		var err error
		var i int
		v.ForEach(func(k, v nitro.Value) bool {
			if i != 0 {
				m.buf.WriteByte(',')
			}
			i++
			err = m.marshal(k)
			if err != nil {
				return false
			}
			m.buf.WriteByte(':')
			err = m.marshal(v)
			if err != nil {
				return false
			}
			return true
		})
		if err != nil {
			return err
		}
		m.buf.WriteByte('}')
	default:
		return fmt.Errorf("cannot marshal value %v to JSON", nitro.TypeName(v))
	}
	return nil
}
