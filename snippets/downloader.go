package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"net/url"
	"strings"
	"net/http"
	"io"
)

func readline(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		s := scanner.Text()
		_, err := url.ParseRequestURI(s)
		if err == nil {
			//fmt.Println(u)
			res = append(res, s)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(res)
	return res
}

func download(url string, c chan interface{}) {
	files := strings.Split(url, "/")
	file := files[len(files) - 1]
	out, err := os.Create(file)
	if err != nil {
		c <- err
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		c <- err
		return
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		c <- err
		return
	}
	c <- n
	return
}

func main() {
	p := os.Args[1]
	urls := readline(p)
	c := make(chan interface{}, len(urls))
	for i := 0; i < len(urls); i++ {
		go download(urls[i], c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
