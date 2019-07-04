package main

import "fmt"

func main() {
	stream := make(chan interface{})
	fmt.Println("main stat")
	go func() {
		stream <- "hello"
	}()
	fmt.Println(<-stream)
	fmt.Println("main end")
}
