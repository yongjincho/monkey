package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, // 0
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"}, // 4

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"}, // 8
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="}, // 12
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"}, // 15
		{token.COMMA, ","},
		{token.IDENT, "y"},  // 18
		{token.RPAREN, ")"}, // 20
		{token.LBRACE, "{"},
		{token.IDENT, "x"}, // 22
		{token.PLUS, "+"},
		{token.IDENT, "y"},     // 25
		{token.SEMICOLON, ";"}, // 27
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"}, // 29

		{token.LET, "let"},
		{token.IDENT, "result"}, // 31
		{token.ASSIGN, "="},
		{token.IDENT, "add"}, // 33
		{token.LPAREN, "("},
		{token.IDENT, "five"}, // 35
		{token.COMMA, ","},
		{token.IDENT, "ten"},   // 37
		{token.RPAREN, ")"},    // 38
		{token.SEMICOLON, ";"}, // 39

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"}, // 43
		{token.INT, "5"},
		{token.SEMICOLON, ";"}, // 45

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"}, // 48

		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"}, // 51
		{token.LT, "<"},
		{token.INT, "10"},   // 53
		{token.RPAREN, ")"}, // 54

		{token.LBRACE, "{"},
		{token.RETURN, "return"}, // 56
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"}, // 58
		{token.RBRACE, "}"},    // 59

		{token.ELSE, "else"}, // 60

		{token.LBRACE, "{"},      // 61
		{token.RETURN, "return"}, // 62
		{token.FALSE, "false"},   // 63
		{token.SEMICOLON, ";"},   // 64
		{token.RBRACE, "}"},      // 65

		{token.INT, "10"}, // 66
		{token.EQ, "=="},
		{token.INT, "10"}, // 68
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
