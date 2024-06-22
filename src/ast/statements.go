package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (b ExpressionStmt) stmt() {}

type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
}

func (b VarDeclStmt) stmt() {}
