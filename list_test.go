package datastructures_test

import (
	"testing"

	"github.com/crockerio/datastructures"
)

func TestAdd(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	expectSizeToBe(t, list, 1)
}

func TestAddm(t *testing.T) {
	list := &datastructures.List{}
	list.Addm([]interface{}{1, 2, 3})
	expectSizeToBe(t, list, 3)
}

func TestRemove(t *testing.T) {
	// Indivdual item
	list := &datastructures.List{}
	list.Add(1)
	expectSizeToBe(t, list, 1)
	list.Remove(1)
	expectSizeToBe(t, list, 0)

	// Multiple items
	multi := &datastructures.List{}
	multi.Add(1)
	multi.Add(2)
	multi.Add(3)
	multi.Add(1)
	expectSizeToBe(t, multi, 4)
	multi.Remove(1)
	expectSizeToBe(t, multi, 2)
}

func TestRemovem(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(1)
	expectSizeToBe(t, list, 4)
	list.Removem([]interface{}{1, 2})
	expectSizeToBe(t, list, 1)
}

func TestRemovei(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(1)
	expectSizeToBe(t, list, 4)
	list.Removei(3)
	expectSizeToBe(t, list, 3)
	// TODO Assert items
}

func TestRemovemi(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(1)
	expectSizeToBe(t, list, 4)
	list.Removemi([]int{1, 2})
	expectSizeToBe(t, list, 2)
	// TODO Assert items
}

func TestGet(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	item := list.Get(1)
	empty := list.Get(4)
	expectEquals(t, 2, item)
	expectEquals(t, nil, empty)
}

func TestHasIndex(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	expectEquals(t, true, list.HasIndex(0))
	expectEquals(t, false, list.HasIndex(1))
}

func TestHasItem(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	expectEquals(t, false, list.HasItem(0))
	expectEquals(t, true, list.HasItem(1))
}

func TestClear(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)
	list.Clear()

	expectSizeToBe(t, list, 0)
}

func TestAll(t *testing.T) {
	list := &datastructures.List{}
	list.Add(1)
	list.Add(2)

	expected := []interface{}{1, 2}

	for i, item := range list.All() {
		expectEquals(t, item, expected[i])
	}
}

func TestCount(t *testing.T) {
	list := &datastructures.List{}
	expectSizeToBe(t, list, 0)
	list.Add(1)
	expectSizeToBe(t, list, 1)
}

func expectEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func expectSizeToBe(t *testing.T, list datastructures.IList, size int) {
	if list.Count() != size {
		t.Errorf("list.Count() = %d; want %d", list.Count(), size)
	}
}
