package tests

import(
	"testing"
	"omix-interpreter/internal/lexer"
)

func TestTokenizerForSimpleExpression(t *testing.T) {
	source := "var xx = 5" ;

	tokenizer := lexer.NewTokenizer(source);
	tokens := tokenizer.ScanTokens();

	expectedTokens := []lexer.Token {
		{Type: lexer.Var, Lexeme: "var", Literal: nil, Line: 1, Column: 4},
		{Type: lexer.Identifier, Lexeme: "xx", Literal: nil, Line: 1, Column: 7},
		{Type: lexer.Equal, Lexeme: "=", Literal: nil, Line: 1, Column: 9},
		{Type: lexer.Number, Lexeme: "5", Literal: 5.0, Line: 1, Column: 11},
		{Type: lexer.Eof, Lexeme: "", Literal: nil, Line: 1, Column: 11},
	}

	if len(tokens) != len(expectedTokens) {
		t.Errorf("Expected %d tokens, got %d", len(expectedTokens), len(tokens))
		return;
	}

	for i, token := range tokens {
		if token.Type != expectedTokens[i].Type || token.Lexeme != expectedTokens[i].Lexeme || token.Literal != expectedTokens[i].Literal || token.Line != expectedTokens[i].Line || token.Column != expectedTokens[i].Column {
			t.Errorf("Token %d: Expected %v, got %v", i, expectedTokens[i], token)
		}
	}
}