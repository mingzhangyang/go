package main

import (
	"fmt"
	"io"
	"net/http"
)

type m struct {
	msg string
}

func (m M) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, m.msg)
	io.WriteString(w, "\n")
	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)
	fmt.Println(r.UserAgent())
}

func main() {
	m := M{"Test Server"}
	err := http.ListenAndServe(":8000", m)
	if err != nil {
		fmt.Println("Error")
	}
}
