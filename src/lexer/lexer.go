package lexer

import "regexp"

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPatten struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPatten
	Tokens   []Token
	source   string
	pos      int
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	return lex.Tokens
}

func (lex *lexer) advanceN(n int) {
	// jump to a specific position in string
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	// append to existing tokens
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	// current postion (index)
	return lex.source[lex.pos]
}

func (lex *lexer) reminder() string {
	// remaining string to traverse
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	// check if it's endof the string (source)
	return lex.pos >= len(lex.source)
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		// advance the lexer's position past the value we just reached
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Tokens, 0),
		patterns: []regexPatten{
			{regexp.MustCompile(`\[`), defaultHandler(OEPN_BRAKET, "[")},
		},
	}
}
