package runtime

import "log"

type StringID int32

type Strings struct {
	strsToID  map[string]StringID
	IDsToStrs []string
}

func NewStrings() *Strings {
	return &Strings{
		strsToID: make(map[string]StringID),
	}
}

func (s *Strings) GetID(str string) StringID {
	if id, ok := s.strsToID[str]; ok {
		return id
	}
	id := StringID(len(s.IDsToStrs))
	s.IDsToStrs = append(s.IDsToStrs, str)
	s.strsToID[str] = id
	return id
}

func (s *Strings) GetString(id StringID) string {
	ndx := int(id)
	if ndx < 0 || ndx >= len(s.IDsToStrs) {
		log.Panicf("No string with ID %v", id)
	}
	return s.IDsToStrs[ndx]
}

type Location struct {
	filename StringID
	lineNum  int
}

type Locations struct {
	strings   *Strings
	locations map[Location]int
}

func NewLocations(strings *Strings) *Locations {
	return &Locations{
		strings:   strings,
		locations: make(map[Location]int),
	}
}

type Program struct {
	globals   int
	fns       []Fn
	extFns    []ExternFn
	literals  []Value
	params    map[string]*Param
	reqParamN int
}
