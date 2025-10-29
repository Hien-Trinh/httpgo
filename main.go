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

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer f.Close()

	curr_line := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			if len(curr_line) != 0 {
				fmt.Printf("read: %s\n", curr_line)
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
			fmt.Printf("read: %s\n", curr_line)
			curr_line = ""
		}
		curr_line += parts[len(parts)-1]
	}

}
