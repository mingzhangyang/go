package main

import (
	"fmt"
	"runtime"
	"io/ioutil"
	"strings"
	"time"
	"os"
)

type record struct {
	id int
	pid int
}

func main() {
	nums := runtime.NumCPU()
	var ws = make(chan record, nums)
	for i := 0; i < nums; i++ {
		go func(i int) {
			content, err := ioutil.ReadFile("urls.txt")
			if err != nil {
				panic(err)
			}
			lines := strings.Split(string(content), "\n")
			fmt.Printf("Reading from goroutine #%d, the first line is %s\n", i, string(lines[0]))
			fmt.Println(time.Now().UnixNano())
			<-time.After(5 * time.Second)
			ws <- record{i, os.Getpid()}
		}(i)
	}
	<-time.After(10 * time.Second)
	close(ws)
	for r := range ws {
		fmt.Printf("goroutine #%d, process id: %d\n", r.id, r.pid)
	}
}