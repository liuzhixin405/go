package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	//"main/stackArray"
)

func GetAllX(path string, files []string, level int) ([]string, error) {
	levelstr := ""
	if level == 1 {
		levelstr = "+"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "+"
	}
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read {
		if fi.IsDir() {
			fullDir := path + "\\" + fi.Name()
			files = append(files, fullDir)
			files, _ = GetAllX(fullDir, files, level+1)

		} else {
			fullDir := path + "\\" + fi.Name()
			files = append(files, levelstr+fullDir)
		}
	}
	return files, nil

}

func main() {
	path := "D:\\GitHub\\go"
	files := []string{}
	files, _ = GetAllX(path, files, 1)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
