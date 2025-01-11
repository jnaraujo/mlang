package repl

import (
	"bufio"
	"fmt"
	"io"
	"mlang/lexer"
	"mlang/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tkn := l.NextToken(); tkn.Type != token.EOF; tkn = l.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}

}
