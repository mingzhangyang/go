package utilities

import (
	"bufio"
	"os"
	"log"
	"fmt"
	//"runtime"
	"encoding/json"
	"errors"
)

func parseLine(m map[string]string, line []byte, headers []string) ([]byte, error) {
	fields := splitBytes(line, byte(','))
	if len(headers) != len(fields) {
		log.Printf("headers: %d, fields: %d\n", len(headers), len(fields))
		fmt.Println(fields)
		return []byte(""), errors.New("not match")
	}
	for i := range headers {
		m[headers[i]] = fields[i]
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("failed to marshal the map")
		fmt.Println(fields)
		return []byte(""), errors.New("Marshal failed")
	}
	return b, nil
}

// CSV2JSON create JSON from CSV
func CSV2JSON(path string) {
	// fmt.Println(runtime.NumCPU())
	c, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to open file")
	}
	defer c.Close()
	j, err := os.Create(path + ".JSON")
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to create file")
	}
	defer j.Close()
	sc := bufio.NewScanner(c)
	wr := bufio.NewWriter(j)
	var counter int
	var line []byte
	var headers []string
	var m = make(map[string]string)

	for sc.Scan() {
		line = sc.Bytes()
		switch {
		case counter > 1:
			b, err := parseLine(m, line, headers)
			if err != nil {
				log.Printf("#%d line parsing failed\n", counter)
				fmt.Println(line)
			}
			wr.Write([]byte(",\n"))
			wr.Write(b)	
		default:
			switch counter {
			case 0:
				headers = splitBytes(line, byte(','))
				wr.Write([]byte("[\n"))
			case 1:
				b, err := parseLine(m, line, headers)
				if err != nil {
					log.Printf("#%d line parsing failed\n", counter)
					fmt.Println(line)
				}
				wr.Write([]byte(",\n"))
				wr.Write(b)	
			}
		}
		counter++
	}
	if err := sc.Err(); err != nil {
		log.Panic(err)
	}
	
	wr.Write([]byte("\n]"))
	err = wr.Flush()
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to flush")
	}
	fmt.Printf("\nConvertion done, %d lines parsed.\n", counter-1)
}