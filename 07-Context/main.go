package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello there")
	ctx, cancel := context.WithCancel(context.Background())
	ctrl := make(chan bool)
	for i := 1; i < 5; i++ {
		go foreverWaiter(ctx, ctrl)
	}
	time.Sleep(2 * time.Second)
	cancel()
	fmt.Print("This is called after go routine ")
	// exampleTimeout()

}
func foreverWaiter(ctx context.Context, ctrlchan chan bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Terminating")
			return
		case vak := <-ctrlchan:
			fmt.Println(strconv.FormatBool(vak))
		default:
			// time.Sleep(5 * time.Second)
			fmt.Println("Starting to sleep")
		}
	}
}

// func exampleTimeout() {
// 	ctx := context.Background()
// 	ctxwithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
// 	defer cancel()
// 	done := make(chan struct{})
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		close(done)
// 	}()
// 	select {
// 	case <-done:
// 		fmt.Println("Called the api")
// 	case <-ctxwithTimeout.Done():
// 		fmt.Println("Oh no my timeout expired")

// 	}

// }
