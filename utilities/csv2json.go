package utilities

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

func combine(headers, values []string) string {
	if len(headers) != len(values) {
		log.Panic("values in the line not matched with headers")
	}
	res := "{"
	for i := range headers {
		res += "\"" + headers[i] + "\":\"" + values[i] + "\","
	
	}
	res = string(res[:len(res)-1])
	return res + "}"
}

// CSV2JSON create JSON from CSV
func CSV2JSON(path string) {
	c, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()
	j, err := os.Create(path + ".JSON")
	if err != nil {
		log.Panic(err)
	}
	defer j.Close()
	sc := bufio.NewScanner(c)
	wr := bufio.NewWriter(j)
	var counter int
	var line string
	var headers, fields []string
	for sc.Scan() {
		line = sc.Text()
		switch {
		case counter > 1:
			fields = splitLine(line, ",")
			wr.WriteString(",\n" + combine(headers, fields))
		default:
			switch counter {
			case 0:
				headers = splitLine(line, ",")
				wr.WriteString("[\n")
			case 1:
				fields = splitLine(line, ",")
				wr.WriteString(combine(headers, fields))
			}
		}
		counter++
	}
	if err := sc.Err(); err != nil {
		log.Panic(err)
	}
	wr.WriteString("\n]")
	err = wr.Flush()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Convertion done, %d lines parsed.\n", counter-1)
}