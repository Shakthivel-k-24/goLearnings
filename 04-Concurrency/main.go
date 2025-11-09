package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int, errchan chan error, donechan chan bool) (int, error) {
	if b == 0 {
		errchan <- errors.New("cant divicde my zero")
		donechan <- true
		return 0, errors.New("cant divicde my zero")
	}
	errchan <- nil
	donechan <- true
	return a / b, nil
}
func main() {
	errchan := make(chan error)
	donechan := make(chan bool)
	go divide(10, 5, errchan, donechan)
	fmt.Println("HElo")
	for range 2 {
		select {
		case err := <-errchan:
			if err != nil {
				fmt.Println("error")

			}
		case <-donechan:
			// if (done!=nil){
			fmt.Println("Done")
			// }
		}

	}

}
