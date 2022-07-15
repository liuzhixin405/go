package main

import (
	"fmt"
	//"main/stackArray"
	"main/arrayList"
)

/*测试arraylist*/
// func main() {
// 	list := arrayList.NewArrayList()
// 	list.Append("a")
// 	list.Append("b")
// 	list.Append("c")
// 	fmt.Println(list)

// 	var ilist arrayList.List = arrayList.NewArrayList()
// 	ilist.Append("d")
// 	ilist.Append("e")
// 	ilist.Append("f")
// 	fmt.Println(ilist)
// 	ilist.Insert(1, "x")
// 	fmt.Println(ilist)
// 	for i := 0; i < 10; i++ {
// 		ilist.Insert(1, "插入")
// 		fmt.Println(ilist)
// 	}
// 	fmt.Println("delete")
// 	ilist.Delete(3)
// 	fmt.Println(ilist)
// }
/*测试iterator*/
// func main() {
// 	list := arrayList.NewArrayList()
// 	list.Append("a")
// 	list.Append("b")
// 	list.Append("c")

// 	for it := list.Iterator(); it.HasNext(); {
// 		item, _ := it.Next()
// 		fmt.Println(item)
// 	}
// }

/*测试stack*/
// func main() {

// 	mystack := stackArray.NewStack()
// 	mystack.Push(1)
// 	mystack.Push(10)
// 	mystack.Push(12)

// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())

// }

//测试 arrayliststack
// func main() {

// 	mystack := arrayList.NewArrayListStack()
// 	mystack.Push(1)
// 	mystack.Push(10)
// 	mystack.Push(12)

// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())
// 	fmt.Println(mystack.Pop())

// }

//测试 arraylistiteratorstack
func main() {
	mystack := arrayList.NewArrayListStackX()
	mystack.Push(1)
	mystack.Push(10)
	mystack.Push(12)

	// fmt.Println(mystack.Pop())
	// fmt.Println(mystack.Pop())
	// fmt.Println(mystack.Pop())
	// fmt.Println(mystack.Pop())

	for it:=mystack.Myit;it.HasNext();{
		item,_:=it.Next()
	    fmt.Println(item)
	}
}
