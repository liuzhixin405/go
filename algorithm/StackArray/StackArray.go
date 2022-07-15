package stackArray

type StackArray interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataStore   []interface{}
	capsize     int //最大容量
	currentsize int //实际使用大小
}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataStore = make([]interface{}, 0, 10)
	mystack.capsize = 10
	mystack.currentsize = 0
	return mystack

}
func (st *Stack) Clear() {
	st.dataStore = make([]interface{}, 0, 10)
	st.currentsize = 0
	st.capsize = 10
}

func (st *Stack) Size() int {
	return st.currentsize
}
func (st *Stack) Pop() (data interface{}) { //弹出
	if !st.IsEmpty() {
		last := st.dataStore[st.currentsize-1]
		st.dataStore = st.dataStore[:st.currentsize-1]
		st.currentsize--
		return last
	}
	return nil
}
func (st *Stack) Push(data interface{}) { //压入
	if !st.IsFull() {
		st.dataStore = append(st.dataStore, data)
		st.currentsize++
	}

}
func (st *Stack) IsFull() bool {
	return st.currentsize >= st.capsize 
		
}
func (st *Stack) IsEmpty() bool {
	return st.currentsize == 0 
}
