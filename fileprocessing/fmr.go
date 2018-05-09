package fileprocessing

import (
	"regexp"
	"os"
	"bufio"
	"io"
	"errors"
	"fmt"
	"log"
)

// StreamEditor remove all the substring matched to the regexp pattern
func StreamEditor(path string, pat string, repl string, delm byte){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("error happens openning the file")
	}
	fr := bufio.NewReader(file)
	// output, err := os.Create(path + ".out")
	re := regexp.MustCompile(pat)
	var b []byte
	var replB = []byte(repl)
	for {
		b, err = fr.ReadBytes(delm)
		if err == io.EOF {
			_, err = os.Stdout.Write(re.ReplaceAll(b, replB))
			if err != nil {
				log.Fatal("error happens writting to stdout")
			}
			break
		}
		if err != nil {
			log.Fatal("error happens reading the file")
		}
		_, err = os.Stdout.Write(re.ReplaceAll(b, replB))
		if err != nil {
			log.Fatal("error happens writting to the stdout")
		}
	}
}

// StreamEditorToFile edit and save to file
func StreamEditorToFile(path string, pat string, repl string, delm byte)error{
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fr := bufio.NewReader(file)
	output, err := os.Create(path + ".out")
	re := regexp.MustCompile(pat)
	var b []byte
	var replB = []byte(repl)
	var count int
	for {
		b, err = fr.ReadBytes(delm)
		if err == io.EOF {
			n, err := output.Write(re.ReplaceAll(b, replB))
			if err != nil {
				return err
			}
			count += n
			break
		}
		if err != nil {
			return errors.New("Error happent in reading the file")
		}
		n, err := output.Write(re.ReplaceAll(b, replB))
		if err != nil {
			return err
		}
		count += n
	}
	fmt.Printf("%d bytes written.\n", count)
	return nil
}