package datastructures

// List - List Data Structure.
//
// Provides a dynamically sized list which can store
// any number of any type of items.
type List struct {
	pointer int
	size    int
	items   []interface{}
}

// IList - List Interface
type IList interface {
	Add(item interface{})
	Addm(items []interface{})
	Remove(item interface{})
	Removem(items []interface{})
	Removei(index int)
	Removemi(indexes []int)
	Get(index int) interface{}
	HasIndex(index int) bool
	HasItem(item interface{}) bool
	Clear()
	All() []interface{}
	Count() int
	resize(newSize int)
}

// Add - Add an item to the List
func (l *List) Add(item interface{}) {
	l.resize(l.size + 1)
	l.items[l.pointer] = item
	l.pointer++
}

// Addm - Add multiple items to the List
func (l *List) Addm(items []interface{}) {
	l.resize(l.size + len(items))
	for _, item := range items {
		l.items[l.pointer] = item
		l.pointer++
	}
}

// Remove - Remove all occurrences of an item from the List
//
// This method will remove all occurrences of the given item
func (l *List) Remove(item interface{}) {
	counter := 0
	newItems := make([]interface{}, l.size)

	for _, i := range l.items {
		if i != item {
			newItems[counter] = i
			counter++
		} else {
			l.pointer--
		}
	}

	l.items = newItems
}

// Removem - Removes all occurrences of the given items from the List
//
// This method will remove all occurrences of the given items
func (l *List) Removem(items []interface{}) {
	for _, item := range items {
		l.Remove(item)
	}
}

// Removei - Remove the item at the given index
func (l *List) Removei(index int) {
	counter := 0
	newItems := make([]interface{}, l.size)

	for i, item := range l.items {
		if i != index {
			newItems[counter] = item
			counter++
		} else {
			l.pointer--
		}
	}

	l.items = newItems
}

// Removemi - Remove the items at the given indexes
func (l *List) Removemi(indexes []int) {
	for _, index := range indexes {
		l.Removei(index)
	}
}

// Get - Retrieve the item at the given index
//
// Returns nil if the index isn't present in the list.
func (l *List) Get(index int) interface{} {
	if !l.HasIndex(index) {
		return nil
	}

	return l.items[index]
}

// HasIndex - Check if the given index exists within the List
func (l *List) HasIndex(index int) bool {
	return l.size > index
}

// HasItem - Check if the given item exists within the List
func (l *List) HasItem(item interface{}) bool {
	for _, i := range l.items {
		if i == item {
			return true
		}
	}

	return false
}

// Clear - Empty the list
func (l *List) Clear() {
	l.items = make([]interface{}, 0)
	l.size = 0
	l.pointer = 0
}

// All - Get all items in the list
func (l *List) All() []interface{} {
	return l.items
}

// Count - Get the number of elements of the List
func (l *List) Count() int {
	return l.pointer
}

func (l *List) resize(newSize int) {
	newItems := make([]interface{}, newSize)

	for i, item := range l.items {
		newItems[i] = item
	}

	l.items = newItems
	l.size = newSize
}
