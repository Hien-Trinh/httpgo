package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath = "messages.txt"

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)

		curr_line := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				if curr_line != "" {
					lines <- curr_line
					curr_line = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

			str := string(data[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				curr_line += parts[i]
				lines <- curr_line
				curr_line = ""
			}
			curr_line += parts[len(parts)-1]
		}
	}()
	return lines
}

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}
