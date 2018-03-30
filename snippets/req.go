package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	c := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.google.com/search?q=machine+learning", nil)
	// fmt.Println(req.URL)
	// fmt.Println(req.Header)
	res, _ := c.Do(req)
	// n, _ := io.Copy(ioutil.Discard, res.Body)
	n, _ := io.Copy(os.Stdout, res.Body)
	fmt.Println(n)
	res.Body.Close()
	fmt.Println()
}
