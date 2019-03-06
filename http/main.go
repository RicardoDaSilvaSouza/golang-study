package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://www.google.com.br/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		/*bytes := make([]byte, 1024)
		_, err := resp.Body.Read(bytes)
		body := string(bytes)
		for err == nil {
			_, err = resp.Body.Read(bytes)
			body += string(bytes)
		}
		fmt.Println(body)*/
		//io.Copy(os.Stdout, resp.Body)
		io.Copy(logWritter{}, resp.Body)
		resp.Body.Close()
	}
}

func (lw logWritter) Write(bs []byte) (written int, err error) {
	fmt.Println(string(bs))
	written = len(bs)
	fmt.Printf("%d bytes was written\n", written)
	return written, err
}
