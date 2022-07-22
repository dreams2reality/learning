package main

import (
	"fmt"
	"time"

	"sandbox/animate"
)

func main() {
	animate.Clear()
	fmt.Println("Gun Firing")
	time.Sleep(time.Second * 3)
	animate.Gun()
	animate.Clear()
	fmt.Println("War Plane")
	time.Sleep(time.Second * 3)
	animate.WarPlane()
	animate.Clear()
	fmt.Println("War!!!")
	time.Sleep(time.Second * 3)
	animate.War()
}
