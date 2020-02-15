package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type writer struct{}

func main() {
	filename := os.Args[1]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	w := writer{}
	io.Copy(w, file)

}

func (writer) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
