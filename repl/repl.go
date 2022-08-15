package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/uliseslugo/lugo_interpreter/evaluator"
	"github.com/uliseslugo/lugo_interpreter/lexer"
	"github.com/uliseslugo/lugo_interpreter/object"
	"github.com/uliseslugo/lugo_interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
