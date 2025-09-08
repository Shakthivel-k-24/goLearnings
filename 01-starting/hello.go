package main

import (
	"fmt"
)

func main() {
	a, b := 10, 10
	formatted_atr := fmt.Sprintf("Hello, World %d, %d", a, b)
	fmt.Println(formatted_atr)
}
