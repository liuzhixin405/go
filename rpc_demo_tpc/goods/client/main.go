package main

import (
	"fmt"
	"net/rpc"
)

type Goods struct {
}

//Add
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}
type AddGoodsResp struct {
	Success bool
	Message string
}

//get
type GetGoodsReq struct {
	Id int
}
type GetGoodsResp struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()

	var reply AddGoodsResp
	err2 := conn.Call("goods.AddGoods", AddGoodsReq{
		Id:      10,
		Title:   "商品标题",
		Price:   10.0,
		Content: "商品内容",
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Printf("addgoods => %#v", reply)

	var getGoods GetGoodsResp

	err3 := conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 10,
	}, &getGoods)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("getgoods =>%#v", getGoods)
}
