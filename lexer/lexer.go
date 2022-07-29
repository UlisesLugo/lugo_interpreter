package lexer

type Lexer struct {
	input 	 	 string
	position 	 int  // Curr pos in input (curr char)
	readPosition int  // Curr reading pos in input (after curr char)
	ch			 byte // current char examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}