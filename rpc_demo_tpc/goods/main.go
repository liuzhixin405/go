package main

import (
	"fmt"
	"net"

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

func (g Goods) AddGoods(req AddGoodsReq, res *AddGoodsResp) error {
	fmt.Printf("AddGoods , %#v", req)

	*res = AddGoodsResp{
		Success: true,
		Message: "success",
	}
	return nil
}

func (g Goods) GetGoods(req GetGoodsReq, res *GetGoodsResp) error {
	fmt.Printf("AddGoods , %#v", req)

	*res = GetGoodsResp{
		Id:      10,
		Title:   "商品10",
		Price:   12.80,
		Content: "内容",
	}
	return nil
}
func main() {

	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println("err1:", err1)
	}

	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println("err2:", err2)
	}

	defer listener.Close()

	for {
		fmt.Println("准备建立链接...")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println("err3:", err3)
		}
		rpc.ServeConn(conn)
	}
}
