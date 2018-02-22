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
	"runtime"
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

func download(client *http.Client, url string, c chan interface{}) {
	// fmt.Println("begin to download " + url)
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

	resp, err := client.Get(url)
	if err != nil {
		c <- err
		return
	}
	defer resp.Body.Close()
	// fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		c <- fmt.Sprintf("Downloading %s failed, status code: %d", file, resp.StatusCode)
		return
	}

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
	mc := make(chan interface{}, len(urls)) // message channel

	n := runtime.NumCPU()
	// fmt.Println(n)
	// n := 4
	wc := make(chan *http.Client, n)
	for i := 0; i < n; i++ {
		go func () {
			wc <- &http.Client{Timeout: 30 * time.Second}
		}()
	}
	for i := 0; i < len(urls); i++ {
		client := <-wc
		go func(client *http.Client, url string) {
			download(client, url, mc)
			wc <- client
		}(client, urls[i])
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-mc)
	}
	fmt.Printf("Total time used: %.2f s\n", time.Since(begin).Seconds())
}
