package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
)

func fromjson(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	input, err := ToReader(ctx, args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid argument #1: %w", err)
	}

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
				a.Append(v)
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
