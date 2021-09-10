package meta

type Metadata struct {
	Params []*Param
}

type Param struct {
	Name     string
	IsFlag   bool
	Type     string
	Desc     string
	Required bool
}
