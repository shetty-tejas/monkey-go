package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"itstejas.com/monkey-go/src/lexer"
	"itstejas.com/monkey-go/src/token"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprint(out, Prompt)
		if err != nil {
			log.Fatal("initialising prompt went wrong.")
			return
		}

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for t := lexer.NextToken(); t.Kind != token.Eof; t = lexer.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", t)

			if err != nil {
				log.Fatal("tokenisation went wrong.")
				return
			}
		}
	}
}
