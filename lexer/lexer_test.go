package lexer

import (
	"testing"

	"go-monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	token.ASSIGN
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
	}
}
