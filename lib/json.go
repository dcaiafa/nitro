package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
)

func parsejson(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 0 {
		return nil, errNotEnoughArgs
	}
	input, err := ToReader(m, args[0])
	if err != nil {
		return nil, err
	}

	defer CloseReader(input)

	v, err := ParseJSON(input)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{v}, nil
}

func ParseJSON(r io.Reader) (nitro.Value, error) {
	parser := &jsonParser{
		dec: json.NewDecoder(r),
	}
	parser.dec.UseNumber()
	v, err := parser.parseValue()
	if err != nil {
		return nil, fmt.Errorf("couldn't parse JSON: %w", err)
	}
	return v, nil
}

type jsonParser struct {
	dec *json.Decoder
	tok json.Token
}

func (p *jsonParser) parseValue() (nitro.Value, error) {
	t, err := p.dec.Token()
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, nil
	}
	switch t := t.(type) {
	case json.Delim:
		if t == '[' {
			a := nitro.NewArray()
			for p.dec.More() {
				v, err := p.parseValue()
				if err != nil {
					return nil, err
				}
				a.Push(v)
			}
			_, err := p.dec.Token()
			if err != nil {
				return nil, err
			}
			return a, nil
		} else {
			o := nitro.NewObject()
			for p.dec.More() {
				k, err := p.parseValue()
				if err != nil {
					return nil, err
				}
				v, err := p.parseValue()
				if err != nil {
					return nil, err
				}
				o.Put(k, v)
			}
			_, err := p.dec.Token()
			if err != nil {
				return nil, err
			}
			return o, nil
		}

	case bool:
		return nitro.NewBool(t), nil

	case json.Number:
		i, err := t.Int64()
		if err == nil {
			return nitro.NewInt(i), nil
		}
		f, err := t.Float64()
		if err == nil {
			return nitro.NewFloat(f), nil
		}
		return nil, err

	case string:
		return nitro.NewString(t), nil

	default:
		return nil, fmt.Errorf("unexpected token %q", t)
	}
}

func tojson(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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

	jsonBytes, err := ToJSON(args[0], indent)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(jsonBytes))}, nil
}

func ToJSON(v nitro.Value, indent string) ([]byte, error) {
	marshaler := jsonMarshaler{}
	err := marshaler.marshal(v)
	if err != nil {
		return nil, err
	}

	jsonBytes := marshaler.buf.Bytes()

	if indent != "" {
		buf := bytes.Buffer{}
		err = json.Indent(&buf, jsonBytes, "", indent)
		if err != nil {
			return nil, err
		}
		jsonBytes = buf.Bytes()
	}

	return jsonBytes, nil
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
				return fmt.Errorf("while marshaling array: %w", err)
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
				err = fmt.Errorf("while marshaling key: %w", err)
				return false
			}
			m.buf.WriteByte(':')
			err = m.marshal(v)
			if err != nil {
				err = fmt.Errorf(
					"while marshaling value for key %q: %w",
					k, err)
				return false
			}
			return true
		})
		if err != nil {
			return fmt.Errorf("while marshaling object: %w", err)
		}
		m.buf.WriteByte('}')

	default:
		b, err := json.Marshal(v.String())
		if err != nil {
			return err
		}
		m.buf.Write(b)
	}

	return nil
}
