package circularsinglylinkedlist

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

	t.Run("check InsertFront", func(t *testing.T) {
		l := List{head: nil, tail: nil}
		l.InsertFront(44)
		l.InsertFront(55)
		l.InsertEnd(100)
		l.InsertFront(1)
		l.InsertEnd(33)

		got := l.Show()
		want := []int{1, 55, 44, 100, 33}

		Equals(t, got, want)
	})

	t.Run("searching elements", func(t *testing.T) {
		l := List{head: nil, tail: nil}
		for i := 1; i <= 5; i++ {
			l.InsertEnd(i * 10)
			l.InsertFront(i * 5)
		}

		got := l.Search(25)
		want := "Found"

		if got != want {
			t.Errorf("got %s want %s and elements are %v", got, want, l.Show())
		}

	})
}
