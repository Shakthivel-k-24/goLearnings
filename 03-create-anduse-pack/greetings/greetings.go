package greetings

import "fmt"

func Hello(name string) (message string) {
	var int_a int
	var int_b int
	var int_c int
	int_a = 10
	int_b = 20
	int_c = int_a + int_b
	message = fmt.Sprintf("the sum of %d, %d is %d", int_a, int_b, int_c)
	return
}
