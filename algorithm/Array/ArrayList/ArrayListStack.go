package arrayList

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	myarray *ArrayList
	capsize int //最大容量

}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList()
	mystack.capsize = 10
	return mystack

}
func (st *Stack) Clear() {
	st.myarray.Clear()
	st.capsize = 10
}

func (st *Stack) Size() int {
	return st.myarray.theSize
}
func (st *Stack) Pop() (data interface{}) { //弹出
	if !st.IsEmpty() {
		last := st.myarray.dataStore[st.myarray.theSize-1]
		st.myarray.Delete(st.myarray.theSize - 1)
		return last
	}
	return nil
}
func (st *Stack) Push(data interface{}) { //压入
	if !st.IsFull() {
		st.myarray.Append(data)
	}

}
func (st *Stack) IsFull() bool {
	return st.myarray.theSize >= st.capsize

}
func (st *Stack) IsEmpty() bool {
	return st.myarray.theSize == 0
}
