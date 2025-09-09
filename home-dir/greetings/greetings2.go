package greetings

import (
	"fmt"
	"os"
)

func Evening_greet(name string) (evening_reset string) {
	evening_reset = fmt.Sprintf("good evening %v, your env var is %v", name, os.Getenv("RANDOM_NUMBER"))
	return
}
