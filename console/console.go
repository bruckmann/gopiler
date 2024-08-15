package console

import (
	"bufio"
	"io"

	"fmt"

	"github.com/bruckmann/gopiler/enums"
	"github.com/bruckmann/gopiler/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for token := l.NextToken(); token.Type != enums.EOF; token = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", token)
		}
	}
}
