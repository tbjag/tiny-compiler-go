package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Value struct {
	IntVal    *int
	StringVal *string
}

type Token interface {
	Name() string
	CommonName() string
	Value() Value
}

func main() {
	filePath := filepath.Join("tests", "test1.t")
	fmt.Println(filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading file %v", err)

	}
	fmt.Println(string(content))

}
