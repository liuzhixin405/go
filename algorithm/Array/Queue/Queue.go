package queue

type MyQueue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	EnQueue(data interface{})
	Dequeue() interface{}
	Clear()
}

type Queue struct {
	dataSource []interface{}
	theSize    int
}

func NewQueue() *Queue {
	myqueue := new(Queue)
	myqueue.dataSource = make([]interface{}, 0)
	myqueue.theSize = 0
	return myqueue
}

func (myqueue *Queue) Size() int {
	return myqueue.theSize
}
func (myqueue *Queue) Front() interface{} {
	if myqueue.Size() == 0 {
		return nil
	}
	return myqueue.dataSource[0]
}
func (myqueue *Queue) End() interface{} {
	if myqueue.Size() == 0 {
		return nil
	}
	return myqueue.dataSource[myqueue.Size()-1]
}
func (myqueue *Queue) IsEmpty() bool {
	return myqueue.theSize == 0
}
func (myqueue *Queue) EnQueue(data interface{}) {
	myqueue.dataSource = append(myqueue.dataSource, data)
	myqueue.theSize++
}
func (myqueue *Queue) Dequeue() interface{} {
	if myqueue.Size() == 0 {
		return nil
	}
	data := myqueue.dataSource[0]
	if myqueue.Size() > 1 {
		myqueue.dataSource = myqueue.dataSource[1:myqueue.Size()]
	} else {
		myqueue.dataSource = make([]interface{}, 0)
	}
	myqueue.theSize--
	return data
}
func (myqueue *Queue) Clear() {
	myqueue.dataSource = make([]interface{}, 0)
	myqueue.theSize = 0
}
