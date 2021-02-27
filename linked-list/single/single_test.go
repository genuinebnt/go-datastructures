package singleLinkedList

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSingleLinkedList(t *testing.T) {
	t.Run("check if empty", func(t *testing.T) {
		l := List{head: nil, tail: nil}
		got := l.Show()
		want := []int{}

		Equals(t, got, want)
	})

	t.Run("check elements", func(t *testing.T) {
		l := List{head: nil, tail: nil}
		l.InsertEnd(22)
		l.InsertEnd(23)

		got := l.Show()
		want := []int{22, 23}

		Equals(t, got, want)
	})
}
