package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Sex  string
}

type Item struct {
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("p.name ,%v, p.sex %v ", p.Name, p.Sex)
}

func (i Item) String() string {
	return fmt.Sprintf("item.name %v", i.Name)
}

func Parse(i interface{}) interface{} {
	switch i.(type) {
	case string:
		return &Item{
			Name: i.(string),
		}
	case []string:
		data := i.([]string)
		length := len(data)
		if length == 2 {
			return &Person{
				Name: data[0],
				Sex:  data[1],
			}
		} else {
			return nil
		}

	default:
		panic(errors.New("type match miss"))
	}
	return nil
}
func main() {
	itemEntity := Parse("jack").(*Item)
	fmt.Println(itemEntity)
	personEntity := Parse([]string{"rose", "female"}).(*Person)
	fmt.Println(personEntity)
}

/*
item.name jack
p.name ,rose, p.sex female
*/
