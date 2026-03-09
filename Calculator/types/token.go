package types

type TokenType int

const (
	Number TokenType = iota
	UnaryOp
	BinaryOp
	Parenthesis
)

type Token struct {
	Type  TokenType
	Value string
}
