package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	expr := `a == 1 && b == 2`
	fileSet := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Print(fileSet, exprAst)
}

