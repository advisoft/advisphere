package lox

import "fmt"

type Expr interface {
	Accept(ExprVisitor) interface{}
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (expr BinaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinaryExpr(expr)
}

type GroupingExpr struct {
	Expression Expr
}

func (expr GroupingExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitGroupingExpr(expr)
}

type LiteralExpr struct {
	Value interface{}
}

func (expr LiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteralExpr(expr)
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (expr UnaryExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnaryExpr(expr)
}

type ExprVisitor interface {
	visitBinaryExpr(expr BinaryExpr) string
	visitGroupingExpr(expr GroupingExpr) string
	visitLiteralExpr(expr LiteralExpr) string
	visitUnaryExpr(expr UnaryExpr) string
}

type AstPrinter struct {
	
}

func (printer AstPrinter) parenthesize(name string, exprs ...Expr) string {
	str := "("
	str += name
	for _, v := range exprs {
		str += " "
		str += v.Accept(printer).(string)
	}
	str += ")"
	return str
}

func (printer AstPrinter) Print(expr Expr) string {
	return expr.Accept(printer).(string)
}

func (printer AstPrinter) visitBinaryExpr(expr BinaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (printer AstPrinter) visitGroupingExpr(expr GroupingExpr) string {
	return printer.parenthesize("group", expr.Expression)
}

func (printer AstPrinter) visitLiteralExpr(expr LiteralExpr) string {
	str, ok := expr.Value.(string)
	if ok {
		return str
	}
	flt, ok := expr.Value.(float64)
	if ok {
		return fmt.Sprintf("%v", flt)
	}
	num, ok := expr.Value.(int)
	if ok {
		return fmt.Sprintf("%v", num)
	}

	return "nil"
}

func (printer AstPrinter) visitUnaryExpr(expr UnaryExpr) string {
	return printer.parenthesize(expr.Operator.Lexeme, expr.Right)
}
