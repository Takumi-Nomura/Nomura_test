package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("C:/Users/bat/go/src/github.com/Nomura_test/hello/data.txt")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
