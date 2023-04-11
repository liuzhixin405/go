package main

import (
	"fmt"
	"go-event/eventbus"
)

type Actor struct{}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}
func main() {
	a := new(Actor)

	eventbus.RegisterEvent("OnSkill", a.OnEvent)
	eventbus.RegisterEvent("OnSkill", GlobalEvent)

	eventbus.CallEvent("OnSkill", 888)
}
