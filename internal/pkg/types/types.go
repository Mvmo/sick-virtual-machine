package types

import (
	"fmt"
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

type Addable interface {
	Add(SickObject) (SickObject, error)
}

type Subtractable interface {
	Subtract(SickObject) (SickObject, error)
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

func (sickString SickString) Add(toAdd SickObject) (SickObject, error) {
	switch toAdd := toAdd.(type) {
	case SickInt:
		return SickString{toAdd.ToHuman() + sickString.Value}, nil
	case SickString:
		return SickString{toAdd.Value + sickString.Value}, nil
	case SickBool:
		return SickString{toAdd.ToHuman() + sickString.Value}, nil
	}
	return nil, fmt.Errorf("can't do %v + %v", sickString.TypeName(), toAdd.TypeName())
}

func (sickString SickString) Subtract(toSubtract SickObject) (SickObject, error) {
	switch toSubtract := toSubtract.(type) {
	case SickInt:
		return SickString{sickString.Value[:len(sickString.Value)-toSubtract.Value]}, nil
	}
	return nil, fmt.Errorf("can't do %v - %v", sickString.TypeName(), toSubtract.TypeName())
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

func (sickInt SickInt) Add(toAdd SickObject) (SickObject, error) {
	switch toAdd := toAdd.(type) {
	case SickInt:
		return SickInt{toAdd.Value + sickInt.Value}, nil
	case SickString:
		return SickString{toAdd.Value + sickInt.ToHuman()}, nil
	}
	return nil, fmt.Errorf("can't do %v + %v", sickInt.TypeName(), toAdd.TypeName())
}

func (sickInt SickInt) Subtract(toSubtract SickObject) (SickObject, error) {
	switch toSubtract := toSubtract.(type) {
	case SickInt:
		return SickInt{toSubtract.Value - sickInt.Value}, nil
	}
	return nil, fmt.Errorf("can't do %v - %v", sickInt.TypeName(), toSubtract.TypeName())
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

func (sickBool SickBool) Add(toAdd SickObject) (SickObject, error) {
	switch toAdd := toAdd.(type) {
	case SickString:
		return SickString{toAdd.Value + sickBool.ToHuman()}, nil
	}
	return nil, fmt.Errorf("can't do %v + %v", toAdd.TypeName(), sickBool.TypeName())
}
