package symbol

import "fmt"

type Type interface {
	fmt.Stringer

	isType()
}

type StringType struct{}

func (s *StringType) String() string { return "String" }
func (s *StringType) isType()        {}

type IntType struct{}

func (i *IntType) String() string { return "Int" }
func (i *IntType) isType()        {}

type FloatType struct{}

func (f *FloatType) String() string { return "Float" }
func (f *FloatType) isType()        {}

type BoolType struct{}

func (b *BoolType) String() string { return "Bool" }
func (b *BoolType) isType()        {}

type AnyType struct{}

func (a *AnyType) String() string { return "Any" }
func (a *AnyType) isType()        {}
