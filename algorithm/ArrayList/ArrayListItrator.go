package arrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

type Iterable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list         *ArrayList
	currentIndex int
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.theSize
}
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个元素")
	}
	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++

	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}
func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
