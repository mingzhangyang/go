package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(f.Name())
	fmt.Fprint(f, "Hello world")
	go func() {
		time.Sleep(time.Second)
		foo, err := os.Create("world.txt")
		if err != nil {
			panic(err)
		}
		defer foo.Close()
		fmt.Println(foo.Name())
		fmt.Fprint(foo, "This is a test.")
	}()
	fmt.Println("Existing the main goroutine")
}
