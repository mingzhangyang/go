package utilities

import (
	"bufio"
	"os"
	"log"
	"fmt"
	//"runtime"
	"encoding/json"
	"errors"
	"strings"
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
					fmt.Println(string(line))
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

func parseLine1(m map[string]string, fields []string, headers []string) ([]byte, error) {
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

// CSV2JSON1 use splitLine1
func CSV2JSON1(path string) {
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
	var line string
	var headers []string
	var fields []string
	var m = make(map[string]string)

	for sc.Scan() {
		line = sc.Text()
		switch {
		case counter > 1:
			splitLine1(&fields, line, ',')
			b, err := parseLine1(m, fields, headers)
			if err != nil {
				log.Printf("#%d line parsing failed\n", counter)
				fmt.Println(line)
			}
			wr.Write([]byte(",\n"))
			wr.Write(b)	
		default:
			switch counter {
			case 0:
				headers = splitLine(line, ',')
				fields = make([]string, len(headers))
				wr.Write([]byte("[\n"))
			case 1:
				splitLine1(&fields, line, ',')
				b, err := parseLine1(m, fields, headers)
				if err != nil {
					log.Printf("#%d line parsing failed\n", counter)
					fmt.Println(line)
				}
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



// CSV2JSON2 try
func CSV2JSON2(path string) {
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
	var line string
	var headers []string
	var fields []string
	var b strings.Builder
	var cols int

	for sc.Scan() {
		line = sc.Text()
		switch {
		case counter > 1:
			splitLine1(&fields, line, ',')
			b.Reset()
			b.Write([]byte(",\n{"))
			for i := 0; i < cols-1; i++ {
				fmt.Fprintf(&b, "\"%s\": %s, ", headers[i], fields[i])
			}
			fmt.Fprintf(&b, "\"%s\": %s}", headers[cols-1], fields[cols-1])
			wr.WriteString(b.String())	
		default:
			switch counter {
			case 0:
				headers = splitLine(line, ',')
				fields = make([]string, len(headers))
				cols = len(headers)
				wr.Write([]byte("["))
			case 1:
				splitLine1(&fields, line, ',')
				b.Write([]byte("\n{"))
				for i := 0; i < cols-1; i++ {
					fmt.Fprintf(&b, "\"%s\": %s, ", headers[i], fields[i])
				}
				fmt.Fprintf(&b, "\"%s\": %s}", headers[cols-1], fields[cols-1])
				wr.WriteString(b.String())	
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