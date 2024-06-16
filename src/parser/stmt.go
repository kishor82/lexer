package parser

import (
	"github.com/kishor82/lexer/src/ast"
	"github.com/kishor82/lexer/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exist := stmt_lu[p.currentTokenKind()]

	if exist {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}
