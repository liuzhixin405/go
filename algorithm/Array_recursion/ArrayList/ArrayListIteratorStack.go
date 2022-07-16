package arrayList

type StackArrayX interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type StackX struct {
	myarray *ArrayList
	Myit Iterator

}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()
	mystack.Myit=mystack.myarray.Iterator()
	return mystack

}
func (st *StackX) Clear() {
	st.myarray.Clear()
    st.myarray.theSize=0
}

func (st *StackX) Size() int {
	return st.myarray.theSize
}
func (st *StackX) Pop() (data interface{}) { //弹出
	if !st.IsEmpty() {
		last := st.myarray.dataStore[st.myarray.theSize-1]
		st.myarray.Delete(st.myarray.theSize - 1)
		return last
	}
	return nil
}
func (st *StackX) Push(data interface{}) { //压入
	if !st.IsFull() {
		st.myarray.Append(data)
	}

}
func (st *StackX) IsFull() bool {
	return st.myarray.theSize >= 10

}
func (st *StackX) IsEmpty() bool {
	return st.myarray.theSize == 0
}
