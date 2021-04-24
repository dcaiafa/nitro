package lib

import (
	"fmt"
	"os"

	"github.com/dcaiafa/nitro"
)

func print(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	stdout := Stdout(m)
	iargs := valuesToInterface(args)
	fmt.Fprintln(stdout, iargs...)
	return nil, nil
}

func printf(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	format, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	stdout := Stdout(m)
	iargs := valuesToInterface(args[1:])
	fmt.Fprintf(stdout, format+"\n", iargs...)

	return nil, nil
}

func printall(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	stdout := Stdout(m)
	e, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}
	for {
		v, ok, err := nitro.Next(m, e, 1)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		fmt.Fprintln(stdout, valuesToInterface(v)...)
	}
	return nil, nil
}

func log(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	iargs := valuesToInterface(args)
	fmt.Fprintln(os.Stderr, iargs...)
	return nil, nil
}

func logf(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) < 1 {
		return nil, errNotEnoughArgs
	}

	format, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	iargs := valuesToInterface(args[1:])
	fmt.Fprintf(os.Stderr, format+"\n", iargs...)

	return nil, nil
}

func sprintf(m *nitro.VM, caps []nitro.ValueRef, args []nitro.Value, nRet int) ([]nitro.Value, error) {
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
