package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		checkList(t, l, []int{70, 80, 60, 40, 10, 30, 50})
	})

	t.Run("complex 2", func(t *testing.T) {
		l := NewList()

		el := l.PushBack(20)   // [20]
		l.Remove(el)           // []
		back := l.PushBack(30) // [30] 30
		l.PushFront(10)        // [10, 30]
		l.MoveToFront(back)    // [30, 10]
		checkList(t, l, []int{30, 10})

		for i := l.Front(); i != nil; i = i.Next {
			l.Remove(i) // []
		}
		checkList(t, l, []int{})

		l.PushFront(20) // [20]
		checkList(t, l, []int{20})
	})
}

func checkList(t *testing.T, l List, es []int) {
	require.Equal(t, len(es), l.Len())

	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	require.Equal(t, es, elems)
}
