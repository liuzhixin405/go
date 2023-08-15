package main

import (
	"booktest/adapter"
	"booktest/service"
	"fmt"
)

func main() {
	cat := service.Cat{}
	dog := service.Dog{}

	fmt.Println(Show(cat))
	fmt.Println(Show(dog))

}

func Show(animal adapter.Animal) string {
	return animal.Speak()
}
