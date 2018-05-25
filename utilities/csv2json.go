package utilities

import (
	"bufio"
	"os"
	"log"
	"fmt"
	//"runtime"
	"encoding/json"
)

func combine(headers, values []string) string {
	if len(headers) != len(values) {
		fmt.Printf("header: %d, line: %d\n", len(headers), len(values))
		fmt.Println(values);
		// log.Panic("values in the line not matched with headers")
		return ""
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
	ch := make(chan []string, 10)
	done := make(chan bool)
	var counter int
	var line string
	var headers, fields []string
	var m = make(map[string]string)
	go func() {
		for fields = range ch {
			if len(headers) != len(fields) {
				log.Println("failed to match the header")
				fmt.Println(line)
				continue
			}
			for i := range headers {
				m[headers[i]] = fields[i]
			}
			b, err := json.Marshal(m)
			if err != nil {
				log.Println("failed to marshal the map")
				fmt.Println(line)
				continue
			}
			wr.Write([]byte(",\n"))
			wr.Write(b)
			// wr.WriteString(",\n" + combine(headers, fields))
		}
		close(done)
	}()
	go func() {
		for sc.Scan() {
			line = sc.Text()
			switch {
			case counter > 1:
				ch <- splitLine(line, ',')		
			default:
				switch counter {
				case 0:
					headers = splitLine(line, ',')
					wr.Write([]byte("[\n"))
					// fmt.Println(len(fields))
				case 1:
					fields = splitLine(line, ',')
					if len(headers) != len(fields) {
						log.Println("failed to match the header")
						fmt.Println(line)
						continue
					}
					for i := range headers {
						m[headers[i]] = fields[i]
					}
					b, err := json.Marshal(m)
					if err != nil {
						log.Println("failed to marshal the map")
						fmt.Println(line)
						continue
					}
					wr.Write(b)
					// wr.WriteString(combine(headers, fields))
				}
			}
			counter++
		}
		if err := sc.Err(); err != nil {
			log.Panic(err)
		}
		close(ch)
	}()
	<- done
	wr.WriteString("\n]")
	err = wr.Flush()
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to flush")
	}
	fmt.Printf("Convertion done, %d lines parsed.\n", counter-1)
}