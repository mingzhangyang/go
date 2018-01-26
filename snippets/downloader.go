package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	// "strings"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"
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
	start := time.Now()
	// files := strings.Split(url, "/")
	// file := files[len(files) - 1]
	file := path.Base(url)
	out, err := os.Create(file)
	if err != nil {
		c <- err
		return
	}
	defer out.Close()

	client := &http.Client{Timeout: 20 * time.Second}

	resp, err := client.Get(url)
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
	var s string
	s += (file + " [" + strconv.Itoa(int(n)) + " bytes] [" + fmt.Sprintf("%.2f s]", time.Since(start).Seconds()))
	c <- s
	return
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("a file containing urls is missing ...")
	}
	p := os.Args[1]
	urls := readline(p)
	begin := time.Now()
	c := make(chan interface{}, len(urls))
	for i := 0; i < len(urls); i++ {
		go download(urls[i], c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
	fmt.Printf("Total time used: %.2f s\n", time.Since(begin).Seconds())
}
