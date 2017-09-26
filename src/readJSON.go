package main
import (
	"os"
	"io"
	"log"
)
func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, input)
}
