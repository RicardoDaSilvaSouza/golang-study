package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if file, err := os.Open(os.Args[1]); err == nil {
		io.Copy(os.Stdout, file)
	} else {
		fmt.Println("Something went wrong trying read file:", err)
	}
}
