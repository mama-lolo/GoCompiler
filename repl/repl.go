package repl

import (
	"bufio"
	"fmt"
	"io"
	"mama-lolo/GoCompiler/lexer"
	"mama-lolo/GoCompiler/token"
)

const prompt = "++>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Fprintf(out, "%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
