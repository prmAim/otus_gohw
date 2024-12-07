package hw04lrucache

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
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

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.theFirst
}

func (l *list) Back() *ListItem {
	return l.theEnd
}

// TODO:
func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	// если l - пустой
	if l == nil {
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

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l == nil {
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

// TODO:
func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	// если это не первый элемент, то ..., иначе ...
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.theFirst.Next = i.Next
	}

	// если это не последний элемент

	l.size--

}

// TODO:
func (l list) MoveToFront(i *ListItem) {
}
