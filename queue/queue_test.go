package queue

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, got, want []int) {

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v got %v", want, got)
	}
}

func TestQueue(t *testing.T) {
	t.Run("checking length elements", func(t *testing.T) {
		q1 := Queue{[]int{11, 55, 55, 33}}
		got := q1.Len()
		want := 4

		if got != want {
			t.Errorf("expected %d got %d", want, got)
		}
	})

	t.Run("checking for underflow", func(t *testing.T) {
		q1 := Queue{[]int{}}
		got := q1.Pop()
		want := []int{}

		Equals(t, got, want)
	})

	t.Run("checking elements", func(t *testing.T) {
		q1 := Queue{[]int{}}
		q1.Push(12)
		q1.Push(33)

		got := q1.Show()
		want := []int{12, 33}

		Equals(t, got, want)
	})
}
