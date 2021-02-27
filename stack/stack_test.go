package stack

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected  %v got %v", got, want)
	}
}

func TestStack(t *testing.T) {

	t.Run("checking length", func(t *testing.T) {
		s := Stack{[]int{1, 2, 3}}
		got := s.Len()
		want := 3

		if got != want {
			t.Errorf("expected %d, want %d", got, want)
		}
	})

	t.Run("checking underflow", func(t *testing.T) {
		s := Stack{[]int{}}
		got := s.Pop()
		want := -1

		if got != want {
			t.Errorf("expected %d want %d", got, want)
		}
	})

	t.Run("checking pop", func(t *testing.T) {
		s := Stack{[]int{22, 22, 11, 67}}

		got := s.Pop()
		want := 67

		if got != want {
			t.Errorf("expected  %d got %d", got, want)
		}
	})

	t.Run("Checking push funtion", func(t *testing.T) {
		s := Stack{[]int{}}

		got := s.Push(22)
		want := []int{22}

		Equals(t, got, want)
	})
}
