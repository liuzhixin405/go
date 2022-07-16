package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func mainX() {
	path := "D:\\GitHub\\go"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
