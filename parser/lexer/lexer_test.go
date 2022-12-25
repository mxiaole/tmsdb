package lexer

import (
	"testing"

	"github.com/mxiaole/myprometheus/parser/token"
)

func TestNextToken(t *testing.T) {
	input := "SELECT ="

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SELECT, "SELECT"},
		{token.EQUAL, "="},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
