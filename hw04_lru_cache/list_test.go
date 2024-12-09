package hw04lrucache

import (
	"testing"

	//nolint:depguard
	"github.com/bxcodec/faker/v3"
	//nolint:depguard
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

	t.Run("TestOfMove", func(t *testing.T) {
		l := NewList()

		for _, v := range [...]string{"first", "second", "third", "fourth"} {
			l.PushFront(v)
		} // [fourth, third, second, first]

		first := l.Back()
		l.MoveToFront(first) // [first, fourth, third, second]
		second := l.Back()
		l.MoveToFront(second) // [second, first, fourth, third]
		first2 := l.Front().Next
		l.MoveToFront(first2) // [first, second, fourth, third]
		third := l.Back()
		l.MoveToFront(third) // [third, first, second, fourth]

		elems := make([]string, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(string))
		}
		require.Equal(t, []string{"third", "first", "second", "fourth"}, elems)
	})
}

// тест на производительность.
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

type testStruct struct {
	Name string `faker:"name"`
}

func newStrFaker() testStruct {
	var strFaker testStruct
	faker.FakeData(&strFaker) // Генерация случайных данных для структуры
	return strFaker
}

func TestListOfFaker(t *testing.T) {
	maxCount := 5
	l := NewList()

	expected := make([]testStruct, maxCount)

	for idx := 0; idx < maxCount; idx++ {
		data := newStrFaker()
		l.PushBack(data)
		expected[idx] = data
	}

	elems := make([]testStruct, 0, l.Len())

	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(testStruct))
	}

	require.Equal(t, expected, elems)
}
