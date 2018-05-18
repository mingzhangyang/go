package main

import (
	"flag"
	"fmt"
	//"net"
	"os"
	//"strconv"
	//"time"
)

func main() {
	var count int
	var timeout int64
	var size int
	var neverstop bool
	flag.Int64Var(&timeout, "w", 1000, "等待每次回复的超时时间(毫秒)。")
	flag.IntVar(&count, "n", 4, "要发送的回显请求数。")
	flag.IntVar(&size, "l", 32, "要发送缓冲区大小。")
	flag.BoolVar(&neverstop, "t", false, "Ping 指定的主机，直到停止。")
	flag.Parse()
	args := flag.Args()
	// fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Usage: ", os.Args[0], "host")
		flag.PrintDefaults()
		flag.Usage()
		os.Exit(1)
	}
}