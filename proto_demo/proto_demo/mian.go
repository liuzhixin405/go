package main

import (
	"fmt"
	"proto_demo/proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {

	u := &userService.Userinfo{
		Username: "张三",
		Age:      18,
		Hobby:    []string{"篮球", "足球"},
	}
	fmt.Printf("%#v\n", u) //#输出标题

	data, _ := proto.Marshal(u)
	fmt.Println(data)
	user := userService.Userinfo{}
	proto.Unmarshal(data, &user)
	fmt.Printf("%v\n", user)
	fmt.Println(user.GetType())
}

//proto生成go文件 protoc --go_out=./ *.proto
