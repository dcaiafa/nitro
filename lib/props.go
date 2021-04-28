package lib

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/dcaiafa/nitro"
)

type fieldMapping struct {
	fieldIndex int
	typ        reflect.Type
}

type Value2Structer struct {
	once   sync.Once
	fields map[string]*fieldMapping
}

func (s *Value2Structer) Convert(from nitro.Value, to interface{}) error {
	s.once.Do(func() { s.init(to) })

	rv := reflect.Indirect(reflect.ValueOf(to))
	fromObj, ok := from.(*nitro.Object)
	if !ok {
		return fmt.Errorf("expected object; received %v", nitro.TypeName(from))
	}

	var err error
	fromObj.ForEach(func(k, v nitro.Value) bool {
		kstr, ok := k.(nitro.String)
		if !ok {
			err = fmt.Errorf("invalid option %v", k.String())
			return false
		}
		field := s.fields[kstr.String()]
		if field == nil {
			err = fmt.Errorf("invalid option %v", kstr.String())
			return false
		}

		rfield := rv.Field(field.fieldIndex)

		switch field.typ.Kind() {
		case reflect.Bool:
			vb, ok := v.(nitro.Bool)
			if !ok {
				err = fmt.Errorf("option %v expected bool; received %v",
					kstr.String(), nitro.TypeName(v))
				return false
			}
			rfield.SetBool(vb.Bool())
		}

		return true
	})

	return err
}

func (s *Value2Structer) init(o interface{}) {
	typ := reflect.Indirect(reflect.ValueOf(o)).Type()

	if typ.Kind() != reflect.Struct {
		panic("value is not a struct")
	}

	s.fields = make(map[string]*fieldMapping)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		nitroName := field.Tag.Get("nitro")
		if nitroName == "" {
			continue
		}

		s.fields[nitroName] = &fieldMapping{
			fieldIndex: i,
			typ:        field.Type,
		}
	}
}
