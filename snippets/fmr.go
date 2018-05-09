package main

import (
	f "go-learning/fileprocessing"
	_ "os"
)

func main() {
	f.StreamEditor("/home/mingzhang/Downloads/test.txt", "aa", "AA", ',')
	f.StreamEditorToFile("/home/mingzhang/Downloads/test.txt", "aa", "AA", ',')
}