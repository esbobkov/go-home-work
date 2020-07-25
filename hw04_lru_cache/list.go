package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Next, Prev *listItem
	Value      interface{}
}

type list struct {
	firstItem, lastItem *listItem
	length              int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.firstItem
}

func (l *list) Back() *listItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *listItem {
	newItem := listItem{Value: v}

	if l.firstItem == nil {
		l.firstItem = &newItem
		l.lastItem = &newItem
		newItem.Next = nil
		newItem.Prev = nil
		l.length++
	} else {
		l.insertBefore(l.firstItem, &newItem)
	}

	return &newItem
}

func (l *list) PushBack(v interface{}) *listItem {
	if l.lastItem == nil {
		return l.PushFront(v)
	}
	newItem := listItem{Value: v}
	l.insertAfter(l.lastItem, &newItem)
	return &newItem
}

func (l *list) Remove(i *listItem) {
	if i.Prev == nil {
		l.firstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.lastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}

func (l *list) insertAfter(item, newItem *listItem) {
	newItem.Prev = item
	if item.Next == nil {
		l.lastItem = newItem
	}
	item.Next = newItem
	l.length++
}

func (l *list) insertBefore(item, newItem *listItem) {
	newItem.Next = item
	if item.Prev == nil {
		l.firstItem = newItem
	}
	item.Prev = newItem
	l.length++
}
