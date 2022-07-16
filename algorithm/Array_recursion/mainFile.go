package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"main/stackArray"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read {
		if fi.IsDir() {
			fullDir := path + "\\" + fi.Name()
			files = append(files, fullDir)
			files, _ = GetAll(fullDir, files)

		} else {
			fullDir := path + "\\" + fi.Name()
			files = append(files, fullDir)
		}
	}
	return files, nil
}

func main1x() {
	path := "D:\\GitHub\\go"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
func main2x() {
	path := "D:\\GitHub\\go"
	files := []string{}
	mystack := stackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path = mystack.Pop().(string)
		files = append(files, path)
		read, _ := ioutil.ReadDir(path)

		for _, fi := range read {
			if fi.IsDir() {
				fullDir := path + "\\" + fi.Name()
				files = append(files, fullDir)
				mystack.Push(fullDir)
			} else {
				fullDir := path + "\\" + fi.Name()
				files = append(files, fullDir)
			}
		}
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

}
