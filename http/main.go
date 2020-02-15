package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lw := LogWriter{}

	io.Copy(lw, resp.Body)

	// resp.Body.Read(a)
	// fmt.Println(string(a))
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote his many bytes:", len(bs))
	return len(bs), nil
}
