package types

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"syscall"
)

func AnyToSickObject(any interface{}) SickObject {
	switch any := any.(type) {
	case SickObject:
		return any
	case string:
		return SickString{any}
	case int:
		return SickInt{any}
	case bool:
		return SickBool{any}
	default:
		log.Fatalf("Types: No parsing for type %v", reflect.TypeOf(any))
		syscall.Exit(0)
		return nil
	}
}

type SickObject interface {
	ToHuman() string
	TypeName() string
}

type SickNum interface {
	AsInt() int
	AsFloat() float64
}

type SickString struct {
	Value string
}

func (SickString) TypeName() string {
	return "sick::string"
}

func (sickString SickString) ToHuman() string {
	return strings.ReplaceAll(sickString.Value, "\\n", "\n") // this is super weird lol
}

type SickInt struct {
	Value int
}

func (SickInt) TypeName() string {
	return "sick::int"
}

func (sickInt SickInt) ToHuman() string {
	return strconv.Itoa(sickInt.Value)
}

func (sickInt SickInt) AsInt() int {
	return sickInt.Value
}

func (sickInt SickInt) AsFloat() float64 {
	return float64(sickInt.Value)
}

type SickBool struct {
	Value bool
}

func (SickBool) TypeName() string {
	return "sick::bool"
}

func (sickBool SickBool) ToHuman() string {
	return strconv.FormatBool(sickBool.Value)
}
