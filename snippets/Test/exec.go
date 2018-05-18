package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	// cmd := exec.Command("echo", "hello")
	cmd := exec.Command("ls", "-l")
	// err = cmd.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	out, err = cmd.Output()
	fmt.Println(string(out))
}
