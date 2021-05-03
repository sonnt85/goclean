package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	gocleanClear()

	for {
		gocleanMoveTopLeft()

		w, h := gocleanSize()
		fmt.Printf("Width: %d Height: %d\n", w, h)

		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}
