package main

import (
	"fmt"
	"os"
	"tbjag/token-go/internal/lexer"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	filePath := "tests/test1.t"
	fmt.Printf("reading in %s\n", filePath)
	file, err := os.Open(filePath)
    check(err)
	defer file.Close()

	ret := lexer.Lex(*file)
	fmt.Println(ret)

}