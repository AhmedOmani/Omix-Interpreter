// Package lexer implements the lexical analysis phase of the Omix interpreter.
// It converts source code into a sequence of tokens that can be processed by the parser.
package lexer

import (
	"fmt"
	"omix-interpreter/errors"
	"strconv"
)

// Token represents a single lexical unit in the source code.
// It contains information about the token's type, the original text (lexeme),
// the interpreted value (literal), and its position in the source code.
type Token struct {
	Type    TokenType   		// The type of token (e.g., Number, String, Identifier)
	Lexeme  string      		// The original text from the source code
	Literal interface{} 		// The interpreted value of the token
	Line    int         		// Line number where the token appears
	Column  int         		// Column number where the token appears
}

// Used for debugging and error reporting.
func (t *Token) String() string {
	return fmt.Sprintf("Type: %s, Lexeme: %s, Literal: %v, Line: %d, Column: %d", t.Type, t.Lexeme, t.Literal, t.Line, t.Column)
}

// Tokenizer is the main lexical analyzer that processes source code into tokens.
// It maintains state about the current position in the source code and collects
// tokens and any lexical errors encountered.
type Tokenizer struct {
	Source  string                		// The complete source code to be tokenized
	Current int                   		// Current position in the source code
	Start   int                   		// Start position of the current token
	Line    int                   		// Current line number
	Column  int                   		// Current column number
	Tokens  []Token               		// List of tokens found so far
	Errors  []errors.LexicalError 		// List of lexical errors encountered
}

// NewTokenizer creates and initializes a new Tokenizer instance.
// It sets up the initial state for tokenizing the given source code.
func NewTokenizer(source string) *Tokenizer {
	Tokenizer := &Tokenizer{
		Source:  source,
		Current: 0,
		Start:   0,
		Line:    1,
		Column:  1,
		Tokens:  []Token{},
		Errors:  []errors.LexicalError{},
	}
	return Tokenizer
}

// ScanTokens processes the entire source code and returns a list of tokens.
// It continues scanning until it reaches the end of the source code.
func (t *Tokenizer) ScanTokens() []Token {
	for !t.atEnd() {
		t.Start = t.Current
		t.ScanToken()
	}
	t.Tokens = append(t.Tokens, Token{Eof, "", nil, t.Line, t.Column})
	return t.Tokens
}

// ScanToken processes a single token from the current position.
// It identifies the type of token and handles it appropriately.
func (t *Tokenizer) ScanToken() {
	c := t.advance()

	//Currenlty the pointer of the source code "current" points to the next character after "c" to be read.
	switch c {

	case '(':
		t.addToken(LeftParen, nil)

	case ')':
		t.addToken(RightParen, nil)

	case '{':
		t.addToken(LeftBrace, nil)

	case '}':
		t.addToken(RightBrace, nil)

	case ',':
		t.addToken(Comma, nil)

	case '.':
		if t.isDigit(t.peek()) {
			t.addError(t.Line, t.Column, "Invalid Number", "LexicalPhase")
			return
		} else if t.isAlpha(t.peek()) {
			t.addError(t.Line, t.Column, "Invalid identifier: cannot start with a dot", "LexicalPhase")
			return
		} else {
			t.addToken(Dot, nil)
		}
	case '-':
		t.addToken(Minus, nil)

	case '+':
		t.addToken(Plus, nil)

	case '*':
		t.addToken(Star, nil)

	case ';':
		t.addToken(Semicolon, nil)

	case '!':
		if t.match('=') {
			t.addToken(BangEqual, nil)
		} else {
			t.addToken(Bang, nil)
		}

	case '=':
		if t.match('=') {
			t.addToken(EqualEqual, nil)
		} else {
			t.addToken(Equal, nil)
		}

	case '<':
		if t.match('=') {
			t.addToken(LessEqual, nil)
		} else {
			t.addToken(Less, nil)
		}

	case '>':
		if t.match('=') {
			t.addToken(GreaterEqual, nil)
		} else {
			t.addToken(Greater, nil)
		}

	case '/':
		if t.match('/') {
			for t.peek() != '\n' && !t.atEnd() {
				t.advance()
			}
		} else {
			t.addToken(Slash, nil)
		}

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		t.Line++
		t.Column = 1
	case '"':
		t.string()

	default:
		if t.isDigit(c) {
			t.number()
		} else if t.isAlpha(c) {
			t.identifier()
		} else {
			t.addError(t.Line, t.Column, fmt.Sprintf("Unexpected character: %c", c), "LexicalPhase")
			return
		}
	}
}

// atEnd checks if we've reached the end of the source code.
func (t *Tokenizer) atEnd() bool {
	return t.Current >= len(t.Source)
}

// advance moves the current position forward and returns the character at that position.
func (t *Tokenizer) advance() rune {
	c := rune(t.Source[t.Current])
	t.Current++
	t.Column++
	return c
}

// addToken creates a new token and adds it to the token list.
func (t *Tokenizer) addToken(tokenType TokenType, Litreal interface{}) {
	text := t.Source[t.Start:t.Current]
	t.Tokens = append(t.Tokens, Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: Litreal,
		Line:    t.Line,
		Column:  t.Column,
	})
}

// match checks if the next character matches the expected character.
// If it matches, advances the current position.
func (t *Tokenizer) match(expected rune) bool {
	if t.atEnd() {
		return false
	}
	if rune(t.Source[t.Current]) != expected {
		return false
	}
	t.Current++
	t.Column++
	return true
}

// peek returns the next character without advancing the current position.
func (t *Tokenizer) peek() rune {
	if t.atEnd() {
		return '\000'
	}
	return rune(t.Source[t.Current])
}

// peekNext returns the character after the next one without advancing.
func (t *Tokenizer) peekNext() rune {
	if t.atEnd() {
		return '\000'
	}
	return rune(t.Source[t.Current+1])
}

// peekToIdx returns the character at the specified offset without advancing.
func (t *Tokenizer) peekToIdx(idx int) rune {
	if t.atEnd() {
		return '\000'
	}
	return rune(t.Source[t.Current+idx])
}

// string processes a string literal from the current position.
// It handles the content between double quotes.
func (t *Tokenizer) string() {
	for t.peek() != '"' && !t.atEnd() {
		t.advance()
	}
	if t.atEnd() {
		t.addError(t.Line, t.Column, "Unterminated String.", "LexicalPhase")
		return
	}
	t.advance()
	value := t.Source[t.Start+1 : t.Current-1]
	t.addToken(String, value)
}

// isDigit checks if a character is a decimal digit.
func (t *Tokenizer) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// number processes a numeric literal from the current position.
// It handles both integer and floating-point numbers.
func (t *Tokenizer) number() {
	// First phase: Collect all digits before decimal point
	for t.isDigit(t.peek()) {
		t.advance()
	}

	// Check for invalid identifier starting with number (e.g., "123abc")
	if t.isAlpha(t.peek()) {
		t.addError(t.Line, t.Column, "Invalid identifier: cannot start with a number", "LexicalPhase")
		return
	}

	// Check for invalid decimal point (e.g., "123.")
	if t.peek() == '.' && !t.isDigit(t.peekNext()) {
		t.addError(t.Line, t.Column, "Invalid Number", "LexicalPhase")
		return
	}

	// Handle decimal point and digits after it
	if t.peek() == '.' && t.isDigit(t.peekNext()) {
		t.advance()
		for t.isDigit(t.peek()) {
			t.advance()
		}
	}

	// Check for invalid decimal point (e.g., "123.4.*")
	if t.peek() == '.' {
		t.addError(t.Line, t.Column, "Invalid Number", "LexicalPhase")
		return
	}

	// Convert the collected number to float64
	lexem := t.Source[t.Start:t.Current]
	number, err := strconv.ParseFloat(lexem, 64)

	if err != nil {
		t.addError(t.Line, t.Column, fmt.Sprintf("Invalid Number: %s", lexem), "LexicalPhase")
		return
	}

	t.addToken(Number, number)
}

// isAlpha checks if a character is a letter or underscore.
func (t *Tokenizer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlphaNumeric checks if a character is a letter, digit, or underscore.
func (t *Tokenizer) isAlphaNumeric(c rune) bool {
	return t.isAlpha(c) || t.isDigit(c)
}

// identifier processes an identifier or keyword from the current position.
// It checks if the identifier matches any reserved keywords.
func (t *Tokenizer) identifier() {
	for t.isAlphaNumeric(t.peek()) {
		t.advance()
	}

	lexem := t.Source[t.Start:t.Current]
	tokenType, ok := keywords[lexem]

	if ok == false {
		t.addToken(Identifier, nil)
	} else {
		t.addToken(tokenType, nil)
	}
}

// addError creates and adds a new lexical error to the error list.
func (t *Tokenizer) addError(line int, column int, message string, phase string) {
	err := &errors.LexicalError{
		BaseError: errors.BaseError{
			Line:    line,
			Column:  column,
			Message: message,
			Phase:   phase,
		},
	}
	t.Errors = append(t.Errors, *err)
}

// HasErrors checks if any lexical errors were encountered.
func (t *Tokenizer) HasErrors() bool {
	return len(t.Errors) > 0
}

// GetErrors returns the list of lexical errors encountered.
func (t *Tokenizer) GetErrors() []errors.LexicalError {
	return t.Errors
}
