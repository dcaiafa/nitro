package analysis

import "fmt"

type GoType struct {
	Package string
	Name    string
	Ref     bool
	Struct  bool
}

func (t GoType) String() string {
	ref := ""
	if t.Ref {
		ref = "*"
	}
	pkg := ""
	if t.Package != "" {
		pkg = fmt.Sprintf("(%s).", t.Package)
	}
	return ref + pkg + t.Name
}

func (t GoType) NoRef() GoType {
	t.Ref = false
	return t
}

type Type struct {
	Name    string
	GoType  GoType
	Nilable bool
	Iter    bool
}

func (t Type) String() string {
	nilable := ""
	if t.Nilable {
		nilable = "?"
	}
	if t.Iter {
		return fmt.Sprintf("Iter<%v>%v", t.GoType.Name, nilable)
	}
	return fmt.Sprintf("%v%v", t.GoType.Name, nilable)
}
