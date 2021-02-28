package meta

import "fmt"

type Type interface {
	fmt.Stringer

	isType()
}

type String struct{}

func (s *String) String() string { return "String" }

type Int struct{}

func (i *Int) String() string {
	return "Int"
}

type Float struct{}

func (f *Float) String() string {
	return "Float"
}

type Bool struct{}

func (b *Bool) String() string {
	return "Bool"
}

type Any struct{}

func (a *Any) String() string {
	return "Any"
}

type Metadata struct {
	Params []Param
}
