package main

import (
	"fmt"
	"time"

	"github.com/sonnt85/goclean"
)

func main() {
	goclean.Clear()

	for {
		goclean.MoveTopLeft()

		w, h := goclean.Size()
		fmt.Printf("Width: %d Height: %d\n", w, h)

		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}
