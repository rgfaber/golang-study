package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	wrt := myWriter{}
	rsp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(wrt, rsp.Body)
}

type myWriter struct{}

func (w myWriter) Write(b []byte) (n int, err error) {
	fmt.Printf("MyWriter Output\n***********\n%+v bytes, Yeah!\n*************", len(b))
	return len(b), nil
}
