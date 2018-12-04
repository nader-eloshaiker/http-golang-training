package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Print("error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
