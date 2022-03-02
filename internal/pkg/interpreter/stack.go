package interpreter

import "mvmo.dev/sickvm/internal/pkg/types"

type Stack []types.SickType

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, types.AnyToSickType(element)) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() types.SickType {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func (s *Stack) Peek() types.SickType {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	return (*s)[index]
}
