package meta

type Metadata struct {
	Params []*Param
}

type Param struct {
	Name       string
	Positional bool
	Type       string
	Desc       string
	Required   bool
}
