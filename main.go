package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer f.Close()

	curr_line := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			curr_line += string(data[:i])
			fmt.Printf("read: %s\n", curr_line)
			data = data[i+1:]
			curr_line = ""
		}

		curr_line += string(data)
	}

	if len(curr_line) != 0 {
		fmt.Printf("read: %s\n", curr_line)
	}
}
