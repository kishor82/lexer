package lexer

import (
	"fmt"
	"regexp"
)

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

	// iterate while we still have tokens
	for !lex.at_eof() {
		mached := false

		for _, pattern := range lex.patterns {

			loc := pattern.regex.FindStringIndex(lex.reminder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				mached = true
				break
			}
		}

		if !mached {
			panic(fmt.Sprintf("Lexer::Error -> unrecognize token near %s\n", lex.reminder()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
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
		Tokens: make([]Token, 0),
		patterns: []regexPatten{
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			// order is important, specifically for keywords like `==` and `=`
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.reminder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.reminder())
	lex.advanceN(match[1])
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.reminder())
	stringLiteral := lex.reminder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.reminder())

	if kind, exists := reserved_lookup[value]; exists {
		lex.push(NewToken(kind, value))
	} else {
		lex.push(NewToken(IDENTIFIER, value))
	}

	lex.advanceN(len(value))
}
