package main

import (
	"fmt"
	"home-dir/greetings"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	defer fmt.Println(greetings.Evening_greet("sme-a"))
	fmt.Println(greetings.Hello("dukc"))

}
