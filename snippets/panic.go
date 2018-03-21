package main

import "fmt"

func main() {
	fmt.Println("Hello Panic")
	panic("Show me the stacks ...") // this line will make the exit status to be 2
}
