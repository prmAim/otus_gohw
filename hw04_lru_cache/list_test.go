package hw04lrucache

import (
	"github.com/bxcodec/faker/v3"
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

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

// тест на производительность
func BenchmarkList(b *testing.B) {
	l := NewList()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			l.PushFront(b.N + i)
		} else {
			l.PushBack(b.N + i)
		}
		if i%3 == 0 {
			l.Remove(l.Front())
		} else {
			l.Remove(l.Back())
		}
	}
}

type StrFaker struct {
	Name string `faker:"name"`
}

func NewStrFaker() StrFaker {
	var strFaker StrFaker
	faker.FakeData(&strFaker) // Генерация случайных данных для структуры
	return strFaker
}

func TestListFaker(t *testing.T) {
	l := NewList()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			l.PushFront(NewStrFaker())
		} else {
			l.PushBack(NewStrFaker())
		}
	}

	require.Equal(t, 10, l.Len())
	require.Equal(t, NewStrFaker(), l.Front().Value)
}
