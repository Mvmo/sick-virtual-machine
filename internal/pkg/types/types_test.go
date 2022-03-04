package types_test

import (
	"testing"

	"mvmo.dev/sickvm/internal/pkg/types"
)

func TestAddition(t *testing.T) {
	testCases := []struct {
		first    types.SickObject
		second   types.SickObject
		expected types.SickObject
	}{
		{types.SickString{"World"}, types.SickString{"Hello, "}, types.SickString{"Hello, World"}},
		{types.SickInt{1}, types.SickString{"1 + "}, types.SickString{"1 + 1"}},
		{types.SickString{" + 1"}, types.SickInt{1}, types.SickString{"1 + 1"}},
		{types.SickBool{true}, types.SickString{"i'm "}, types.SickString{"i'm true"}},
		{types.SickString{" i'm"}, types.SickBool{true}, types.SickString{"true i'm"}},
	}

	for _, testCase := range testCases {
		switch first := testCase.first.(type) {
		case types.Addable:
			t.Logf("Run Test: (%v) %v + (%v) %v", testCase.first.TypeName(), testCase.first.ToHuman(), testCase.second.TypeName(), testCase.second.ToHuman())
			result, err := first.Add(testCase.second)

			if err != nil {
				t.Error(err)
			}

			if result != testCase.expected {
				t.Errorf("expected %v and got %v", result.ToHuman(), testCase.expected.ToHuman())
			}
		default:
			t.Errorf("first needs to be Addable")
		}
	}
}
