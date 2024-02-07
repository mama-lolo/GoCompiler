package ast

type Node interface {
	TokenLiteral() string
}
type Satement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}
