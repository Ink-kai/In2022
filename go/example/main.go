package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", err)
	// Output:
	// some io.Reader stream to be read
}
