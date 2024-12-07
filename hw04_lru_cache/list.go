package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	StringAll(str string)
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	// Place your code here.
	theFirst *ListItem // первый элемент
	theEnd   *ListItem // последний элемент
	size     int       // размер списка
}

func NewList() List {
	return new(list)
}

// длина списка
func (l *list) Len() int {
	return l.size
}

// первый элемент списка
func (l *list) Front() *ListItem {
	return l.theFirst
}

// последний элемент списка
func (l *list) Back() *ListItem {
	return l.theEnd
}

// добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	if l == nil {
		return nil
	}

	newItem := &ListItem{Value: v}

	// если l - пустой
	if l.size == 0 {
		l.theFirst = newItem
		l.theEnd = newItem
	} else {
		newItem.Next = l.theFirst
		l.theFirst.Prev = newItem
		l.theFirst = newItem
	}

	l.size++

	return newItem
}

// добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	if l == nil {
		return nil
	}
	newItem := &ListItem{Value: v}

	if l.size == 0 {
		l.theFirst = newItem
		l.theEnd = newItem
	} else {
		l.theEnd.Next = newItem
		newItem.Prev = l.theEnd
		l.theEnd = newItem
	}

	l.size++

	return newItem
}

// удалить элемент
func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	// если это не первый элемент, то ..., иначе
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.theFirst.Next = i.Next
	}

	// если это не последний элемент, то ..., иначе
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.theEnd.Prev = i.Prev
	}

	l.size--
}

// переместить элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	// если у элемента ссылка на преддыдущий элемент = nil, то это и есть первый элемент
	if i.Prev == nil {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) StringAll(str string) {
	if l == nil || l.size == 0 {
		fmt.Println(str + ") = nil")
		return
	}

	for step := l.theFirst; step != nil; step = step.Next {
		fmt.Printf("%q) Element %+v \n", str, step.Value)
	}
}
