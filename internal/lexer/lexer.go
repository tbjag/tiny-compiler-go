package lexer

import (
	"bufio"
	"os"
	"tbjag/token-go/internal/model"
)

func Lex(file os.File) string {
	newReader := bufio.NewReader(&file)
	g := model.Generator{
		Reader: *newReader,
	}


	// move forloop into generator functions
	res := ""
	for {
        ch, err := g.Reader.ReadByte() // reads one byte
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            panic(err)
        }
        res += string(ch)
    }

	return res
}
