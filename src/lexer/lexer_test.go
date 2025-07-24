package lexer

import (
	"testing"

	"itstejas.com/monkey-go/src/token"
)

func TestNextTokenRandomSymbols(t *testing.T) {
	input := "=+{}(),;"

	tests := []struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.LeftBrace, "{"},
		{token.RightBrace, "}"},
		{token.LeftParen, "("},
		{token.RightParen, ")"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
	}

	lexer := NewLexer(input)

	for i, tc := range tests {
		token := lexer.NextToken()

		if tc.expectedKind != token.Kind {
			t.Fatalf("tests[%d] - kind wrong. expected %q, got %q", i, tc.expectedKind, token.Kind)
		}

		if tc.expectedLiteral != token.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tc.expectedLiteral, token.Literal)
		}
	}
}

func TestNextTokenValidSyntax(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedKind    token.Kind
		expectedLiteral string
	}{
		{token.Let, "LET"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "LET"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Let, "LET"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "FUNCTION"},
		{token.LeftParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "LET"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LeftParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RightParen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.If, "IF"},
		{token.LeftParen, "("},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "RETURN"},
		{token.Boolean, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "ELSE"},
		{token.LeftBrace, "{"},
		{token.Return, "RETURN"},
		{token.Boolean, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Integer, "10"},
		{token.Equals, "=="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.NotEquals, "!="},
		{token.Integer, "9"},
		{token.Semicolon, ";"},
		{token.Eof, ""},
	}

	lexer := NewLexer(input)

	for i, tc := range tests {
		token := lexer.NextToken()

		if tc.expectedKind != token.Kind {
			t.Fatalf("tests[%d] - kind wrong. expected %q, got %q", i, tc.expectedKind, token.Kind)
		}

		if tc.expectedLiteral != token.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tc.expectedLiteral, token.Literal)
		}
	}
}
