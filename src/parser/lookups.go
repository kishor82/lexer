package parser

import (
	"github.com/kishor82/lexer/src/ast"
	"github.com/kishor82/lexer/src/lexer"
)

type binding_power int

const (
	default_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type (
	stmt_handler func(p *parser) ast.Stmt
	nud_handler  func(p *parser) ast.Expr
	led_handler  func(p *parser, left ast.Expr, bp binding_power) ast.Expr
)

type (
	stmt_lookup map[lexer.TokenKind]stmt_handler
	nud_lookup  map[lexer.TokenKind]nud_handler
	led_lookup  map[lexer.TokenKind]led_handler
	bp_lookup   map[lexer.TokenKind]binding_power
)

var (
	bp_lu   = bp_lookup{}
	nud_lu  = nud_lookup{}
	led_lu  = led_lookup{}
	stmt_lu = stmt_lookup{}
)

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// Literal and Symbols
	nud(lexer.NUMBER, parse_primary_expr)
	nud(lexer.STRING, parse_primary_expr)
	nud(lexer.IDENTIFIER, parse_primary_expr)
}