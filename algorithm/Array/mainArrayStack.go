package main

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
// func main() {
// 	mystack := arrayList.NewArrayListStackX()
// 	mystack.Push(1)
// 	mystack.Push(10)
// 	mystack.Push(12)

// 	// fmt.Println(mystack.Pop())
// 	// fmt.Println(mystack.Pop())
// 	// fmt.Println(mystack.Pop())
// 	// fmt.Println(mystack.Pop())

// 	for it:=mystack.Myit;it.HasNext();{
// 		item,_:=it.Next()
// 	    fmt.Println(item)
// 	}
// }

//递归
// func main() {
// 	mystack := stackArray.NewStack()
// 	mystack.Push(5)
// 	last := 0
// 	for !mystack.IsEmpty() {
// 		data := mystack.Pop()

// 		if data == 0 {
// 			last += 0
// 		} else {
// 			last += data.(int)
// 			mystack.Push((data.(int) - 1))
// 		}
// 	}
// 	fmt.Println(last)

// }

//"main/arrayList"

//斐波那契数列
// func FAB(num int) int {
// 	if num == 1 || num == 2 {
// 		return 1
// 	}
// 	return FAB(num-1) + FAB(num-2)
// }

// func main(){
// 	fmt.Println(FAB(7))
// }

//递归斐波那契  通过栈实现的递归
// func main() {
// 	mystack := stackArray.NewStack()
// 	mystack.Push(7)
// 	last := 0
// 	for !mystack.IsEmpty() {
// 		data := mystack.Pop()

// 		if data == 1 || data == 2 {
// 			last += 1
// 		} else {
// 			mystack.Push((data.(int) - 1))
// 			mystack.Push((data.(int) - 2))
// 		}
// 	}
// 	fmt.Println(last)

// }
