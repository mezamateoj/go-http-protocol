package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := make([]byte, 8)

	test := ""
	for {
		chunk, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		parts := strings.Split(string(buf[:chunk]), "\n")
		for i := range parts {
			if i != len(parts)-1 {
				test += parts[i]
				fmt.Printf("read: %s\n", test)
				test = ""
			}

			if i == len(parts)-1 {
				test += parts[i]
			}
		}
	}

}
