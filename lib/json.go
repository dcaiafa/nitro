package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

func parseJSON(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 0 {
		return nil, errNotEnoughArgs
	}
	input, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, err
	}

	defer core.CloseReader(input)

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
				a.Add(v)
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

type toJSONOptions struct {
	Indent nitro.Value `nitro:"indent"`
	Prefix nitro.Value `nitro:"prefix"`
}

var toJSONOptionConv core.Value2Structer

var errToJSONUsage = errors.New(
	`invalid usage. Expected to_json(bool|int|float|string|array|map, map?)`)

func toJSON(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errToJSONUsage
	}

	indent := ""
	prefix := ""
	if len(args) == 2 {
		var opt toJSONOptions
		err := toJSONOptionConv.Convert(args[1], &opt)
		if err != nil {
			return nil, err
		}
		indent, err = getStringOrSpaceSize(opt.Indent)
		if err != nil {
			return nil, fmt.Errorf("invalid indent option: %w", err)
		}
		prefix, err = getStringOrSpaceSize(opt.Prefix)
		if err != nil {
			return nil, fmt.Errorf("invalid prefix option: %w", err)
		}
	}

	jsonBytes, err := ToJSON(args[0], prefix, indent)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(string(jsonBytes))}, nil
}

func getStringOrSpaceSize(v nitro.Value) (string, error) {
	if v == nil {
		return "", nil
	}

	indent := ""
	switch v := v.(type) {
	case nitro.Int:
		indentN := int(v.Int64())
		if indentN < 0 || indentN >= 20 {
			return "", fmt.Errorf("invalid length %v", v.Int64())
		}
		indentB := make([]byte, indentN)
		for i := range indentB {
			indentB[i] = ' '
		}
		indent = string(indentB)

	case nitro.String:
		indent = v.String()
		if len(indent) > 20 {
			return "", fmt.Errorf("invalid length %v", len(indent))
		}

	default:
		return "", fmt.Errorf(
			"must be an int or string. It was %v",
			nitro.TypeName(v))
	}

	return indent, nil
}

func ToJSON(v nitro.Value, prefix, indent string) ([]byte, error) {
	marshaler := jsonMarshaler{}
	err := marshaler.marshal(v)
	if err != nil {
		return nil, err
	}

	jsonBytes := marshaler.buf.Bytes()

	if indent != "" {
		buf := bytes.Buffer{}
		err = json.Indent(&buf, jsonBytes, prefix, indent)
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
