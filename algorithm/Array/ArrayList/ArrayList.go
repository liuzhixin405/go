package arrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                  //大小
	Get(index int) (interface{}, error)         //获取指定元素
	Set(index int, newval interface{}) error    //修改数据
	Insert(index int, newval interface{}) error // 插入
	Append(newval interface{})                  //追加
	Clear()
	Delete(index int) error
	String() string
	Len() int
	Iterator() Iterator
}

type ArrayList struct {
	dataStore []interface{} //数组存储
	theSize   int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil

}

func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)
	list.theSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Clear() {

}
func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //0-index 和 index+1叠加
	list.theSize--
	return nil
}
func (list *ArrayList) Len() int {
	return list.theSize
}
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index > list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval
	return nil
}
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.checkFull()
	list.dataStore = list.dataStore[:list.theSize+1]
	for i := list.theSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.theSize++
	return nil
}
func (list *ArrayList) checkFull() {
	if list.theSize == cap(list.dataStore) {
		newdataSource := make([]interface{}, 2*list.theSize, 2*list.theSize) //注意
		copy(newdataSource, list.dataStore)
		// for i := 0; i < len(list.dataStore); i++ {
		// 	newdataSource[i] = list.dataStore[i]
		// }
		list.dataStore = newdataSource
	}
}
