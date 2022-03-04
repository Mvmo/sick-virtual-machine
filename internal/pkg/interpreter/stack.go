package interpreter

import "mvmo.dev/sickvm/internal/pkg/types"

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the SickObjectStack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element) // Simply append the new value to the end of the SickObjectStack
}

// Remove and return top element of SickObjectStack. Return false if SickObjectStack is empty.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the SickObjectStack by slicing it off.
	return element
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	return (*s)[index]
}

type SickObjectStack []types.SickObject

func (s *SickObjectStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the SickObjectStack
func (s *SickObjectStack) Push(element interface{}) {
	*s = append(*s, types.AnyToSickObject(element)) // Simply append the new value to the end of the SickObjectStack
}

// Remove and return top element of SickObjectStack. Return false if SickObjectStack is empty.
func (s *SickObjectStack) Pop() types.SickObject {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the SickObjectStack by slicing it off.
	return element
}

func (s *SickObjectStack) Peek() types.SickObject {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	return (*s)[index]
}
