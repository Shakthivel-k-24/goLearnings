package greetings

import (
	"fmt"
)

func Hello(name string) (ret string) {
	ret = fmt.Sprintf("this is a hello %v", name)
	return
}
