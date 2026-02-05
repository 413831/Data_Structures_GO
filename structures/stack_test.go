package structures

import "testing"

func TestStack_Push(t *testing.T) {
	testCases := []struct {
		name     string
		values   []int
		expected int64
	}{{
		name:     "push 1",
		values:   []int{1, 10, 60},
		expected: 3,
	}}

	for _, testCase := range testCases {
		stack := NewStack[int]()

		for _, value := range testCase.values {
			stack.Push(value)
		}

		if stack.Size() != testCase.expected {
			t.Fail()
		}
	}
}
