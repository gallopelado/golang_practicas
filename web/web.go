package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	// respuesta, err := http.Get("http://google.com")
	respuesta, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err)
		return
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)
	// fmt.Println(respuesta.Body)
}
