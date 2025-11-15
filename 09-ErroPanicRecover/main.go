package main

import (
	"fmt"
)

func main() {
	fmt.Println("Entering main")
	defer func() {

		// err := recover()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Println("Exited cleanly")
	}()
	panic("I am panicking")

}
