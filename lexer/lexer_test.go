package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		got := l.NextToken()

		if got.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected %q got %q", i, tt.expectedType, got.Type)
		}

		if got.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected %q got %q", i, tt.expectedLiteral, got.Literal)
		}
	}

}
