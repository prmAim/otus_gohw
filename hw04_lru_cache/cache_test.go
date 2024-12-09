package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("push out an element", func(t *testing.T) {
		testData := []struct {
			input           string
			inputGetTest    string
			expectedGetTest bool
		}{
			{input: "first", inputGetTest: "first", expectedGetTest: true},   // [first]
			{input: "second", inputGetTest: "second", expectedGetTest: true}, // [second][first]
			{input: "third", inputGetTest: "third", expectedGetTest: true},   // [third][second][first]
			{input: "fourth", inputGetTest: "first", expectedGetTest: false}, // [fourth][third][second] вытолкнули [first]
			{input: "fifth", inputGetTest: "second", expectedGetTest: false}, // [fifth][fourth][third]  вытолкнули [second]
		}

		testCache := NewCache(3)

		for _, tc := range testData {
			tc := tc
			t.Run(tc.input, func(t *testing.T) {
				testCache.Set(Key(tc.input), tc.input)

				_, isOk := testCache.Get(Key(tc.inputGetTest))
				require.Equal(t, tc.expectedGetTest, isOk)
			})
		}
	})

	t.Run("testOfOldElement", func(t *testing.T) {
		testCache := NewCache(4)

		testCache.Set(Key("first"), "first")           // [first]
		testCache.Set(Key("second"), "second")         // [second][first]
		testCache.Set(Key("third"), "third")           // [third][second][first]
		testCache.Set(Key("oldElement"), "oldElement") // [oldElement][third][second][first]

		testCache.Get(Key("first"))  // [first][oldElement][third][second]
		testCache.Get(Key("second")) // [second][first][oldElement][third]
		testCache.Get(Key("first"))  // [first][second][oldElement][third]
		testCache.Get(Key("third"))  // [third][first][second][oldElement]

		testCache.Set(Key("fifth"), "fifth") // [fifth][third][first][second]

		_, isOk := testCache.Get(Key("oldElement")) // данного элемента нет, его вытолкнули, так как он самый старый
		require.Equal(t, false, isOk)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
