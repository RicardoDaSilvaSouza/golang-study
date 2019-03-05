package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com.br/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		var bytes []byte
		body := ""
		read, err := resp.Body.Read(bytes)
		for read > 0 && err != nil {
			body += string(bytes)
			read, err = resp.Body.Read(bytes)
		}
		fmt.Println(body)
	}
}
