package analysis

type Struct struct {
	Name   string
	Fields []*StructField
}

type StructField struct {
	Name string
	Type Type
}
