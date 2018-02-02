package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ReadLineObject struct {
	FileName    string
	CountOfChar int
	CountOfLine int
	FileInfo    os.FileInfo
	Size        int64
	Done        bool
	src         <-chan string
	currentLine string
}

func NewReadLineObject(name string) (*ReadLineObject, error) {
	var rlo *ReadLineObject
	f, err := os.Open(name)
	if err != nil {
		return rlo, err
	}
	rlo = &ReadLineObject{FileName: name}
	rlo.FileInfo, _ = f.Stat()
	rlo.Size = rlo.FileInfo.Size()
	rlo.src = func() <-chan string {
		c := make(chan string)
		go func() {
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				rlo.currentLine = scanner.Text()
				rlo.CountOfChar += len(rlo.currentLine)
				rlo.CountOfLine += 1
				c <- rlo.currentLine
			}
			rlo.Done = true
			close(c)
			f.Close()
		}()
		return c
	}()
	return rlo, nil
}

func (r ReadLineObject) ReadLine() string {
	return <-(r.src)
}

func main() {
	path := os.Args[1]
	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	rlo, err := NewReadLineObject(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rlo)
	fmt.Println(rlo.ReadLine())
	fmt.Println("--------------------")

	// This is not a good idea. Because user may just read several lines from the file.
	// If the ReadLineObject is not done, there will always be a goroutin waiting for being read from the channel.
	// So it is better to use bufio.Scanner to read a file line by line.

}
