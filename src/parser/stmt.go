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

func parse_var_decl_stmt(p *parser) ast.Stmt {
	isConstant := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration, expected to find variable name").Value
	p.expect(lexer.ASSIGNMENT)
	assignedValue := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)

	return ast.VarDeclStmt{
		VariableName:  varName,
		IsConstant:    isConstant,
		AssignedValue: assignedValue,
	}
}
