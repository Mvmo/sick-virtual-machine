package types

import (
	"syscall"
)

func AnyToSickType(any interface{}) SickType {
	switch any.(type) {
	case string:
		return SickString{any.(string)}
	case int:
		return SickInt{any.(int)}
	case bool:
		return SickBool{any.(bool)}
	default:
		syscall.Exit(-1)
		return nil
	}
}

type SickType interface {
	ToHuman() string
}

type SickNum interface {
	AsInt() int
	AsFloat() float64
}

type SickString struct {
	Value string
}

func (self SickString) ToHuman() string {
	return "string"
}

type SickInt struct {
	Value int
}

func (self SickInt) ToHuman() string {
	return "int"
}

func (self SickInt) AsInt() int {
	return self.Value
}

func (self SickInt) AsFloat() float64 {
	return float64(self.Value)
}

type SickBool struct {
	Value bool
}

func (self SickBool) ToHuman() string {
	return "bool"
}
