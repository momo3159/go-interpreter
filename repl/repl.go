package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"momo3159/go-interpreter/lexer"
	"momo3159/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			fmt.Println("bye!")
			os.Exit(0)
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF;tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
