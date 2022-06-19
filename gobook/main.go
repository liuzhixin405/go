package main

import (
	"fmt"
	"time"
)

type user struct {
	name      string
	email     string
	ext       int
	privleged bool
}
type admin struct {
	person user
	level  string
}

func main() {
	lisa := user{
		name:      "Lisa",
		email:     "1234@123.com",
		ext:       123,
		privleged: true,
	}
	lisa.noyify()
	fred := admin{
		person: user{
			name:      "Lisa",
			email:     "1234@123.com",
			ext:       123,
			privleged: true,
		},
		level: "super",
	}
	lzx := person("lzx")
	var dur Duration

	dur = Duration(100)

	fmt.Printf("%+v", lisa)
	fmt.Println()
	fmt.Printf("%+v", fred)
	fmt.Println()
	fmt.Printf("%+v", lzx)
	fmt.Println()
	fmt.Println(dur)

	var a int = 1
	changeValue(&a)
	fmt.Println("a=", a)

	plist := &user{
		name:      "pLisa",
		email:     "333@1667.com",
		ext:       123,
		privleged: true,
	}
	plist.noyify()
	fmt.Printf("sending user email to %s < %s> \n", plist.name, plist.email)

	plist.changeEmain("hhhhh.@hhhh.com")
	fmt.Printf("sending user email to %s < %s> \n", plist.name, plist.email)
}

func (u *user) changeEmain(email string) {
	u.email = email
}
func changeValue(p *int) {
	*p = 10
}
func (u *user) noyify() {
	u.email = "163@163.com"
	//fmt.Printf("sending user email to %s < %s> \n", u.name, u.email)
}

type person string
type Duration int64

//切片  通道  映射 函数类型 接口

type IP []byte

func (ip IP) MarshalText() ([]byte, error) {
	return []byte(ip), nil
}

func ipEmptyString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return string(ip)
}

type Time struct {
	sec  int64
	nsec int32
	loc  *time.Location
}

func Now() Time {
	return Time{time.Now().Unix(), 0, time.Local}
}

type File struct {
	*file
}
type file struct {
	fd      int
	name    string
	//dirinfo *dirInfo
	nepipe  int32
}

