package ddd

import (
	"EDA_GO/internal/registry"
	"fmt"
)

type IDSetter interface {
	SetID(string)
}

func SetID(id string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(IDSetter); ok {
			e.SetID(id)
			return nil
		}
		return fmt.Errorf("%T does not have method SetID(string)", v)
	}
}

type NameSetter interface {
	SetName(string)
}

func SetName(name string) registry.BuildOption {
	return func(v interface{}) error {
		if e, ok := v.(NameSetter); ok {
			e.SetName(name)
			return nil
		}
		return fmt.Errorf("%T does not have method SetName(string)", v)
	}
}
