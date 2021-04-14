package meta

type Metadata struct {
	Params []*Param
}

type Param struct {
	Name string
	Type string
	Desc string
}
