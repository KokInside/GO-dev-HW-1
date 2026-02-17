package types

type TokenType int

const (
	Number TokenType = iota
	Operator
	Parenthesis
)

type Token struct {
	Type  TokenType
	Value string
}
