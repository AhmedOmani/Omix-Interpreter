package main 

import (
	"strconv"
)
type Token struct {
	Type TokenType
	Lexeme string
	Literal interface{} 
	Line int
}

type Tokenizer struct {
	Source string
	Current int 
	Start int
	Line int
	Tokens []Token
}

func newTokenizer(source string) *Tokenizer {
	Tokenizer := &Tokenizer{
		Source: source,
		Current: 0,
		Start: 0,
		Line: 1,
		Tokens: []Token{},
	}
	return Tokenizer ;
}

func (t* Tokenizer) scanTokens() []Token {
	for !t.atEnd() {
		t.Start = t.Current;
		t.scanToken();
	}
	t.Tokens = append(t.Tokens , Token{Eof , "", nil , t.Line});
	return t.Tokens;
}

func (t *Tokenizer) scanToken() {
	c := t.advance();

	//Currenlty the pointer of the source code "current" points to the next character to be read.
	switch c {

	case '(' :
		t.addToken(LeftParen ,  nil);

	case ')' :
		t.addToken(RightParen , nil);

	case '{':
		t.addToken(LeftBrace , nil);

	case '}':
		t.addToken(RightBrace , nil);

	case ',':
		t.addToken(Comma , nil);

	case '.':
		t.addToken(Dot , nil);

	case '-':
		t.addToken(Minus , nil);
		
	case '+':
		t.addToken(Plus , nil);

	case '*':
		t.addToken(Star , nil);

	case ';':
		t.addToken(Semicolon , nil);
	
	case '!':
		if t.match('=') {
			t.addToken(BangEqual , nil);
		} else {
			t.addToken(Bang , nil);
		}

	case '=':
		if t.match('=') {
			t.addToken(EqualEqual , nil);
		} else {
			t.addToken(Equal , nil);
		}	

	case '<':
		if t.match('=') {
			t.addToken(LessEqual , nil);
		} else {
			t.addToken(Less , nil);
		}

	case '>':
		if t.match('=') {
			t.addToken(GreaterEqual , nil);
		} else {
			t.addToken(Greater , nil);
		}

	case '/':
		if t.match('/') {
			for t.peek() != '\n' && !t.atEnd() {
				t.advance();
			}
		} else {
			t.addToken(Slash , nil);
		}

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		t.Line++;
	case '"':
		t.string();

	default:
		if isDigit(c) {
			t.number();
		} else {
			//TODO: Handle error
		}
	}
}

func (t *Tokenizer) atEnd() bool {
	return t.Current >= len(t.Source);
}

func (t *Tokenizer) advance() rune {
	c := rune(t.Source[t.Current]);
	t.Current++;
	return c ;
}

func (t *Tokenizer) addToken(tokenType TokenType , Litreal interface{}) {
	text := t.Source[t.Start:t.Current];
	t.Tokens = append(t.Tokens , Token{
		Type: tokenType,
		Lexeme: text,
		Literal: Litreal,
		Line: t.Line,
	});
}

func (t *Tokenizer) match(expected rune) bool {
	if t.atEnd() {
		return false ;
	}
	if rune(t.Source[t.Current]) != expected {
		return false ;
	}
	t.Current++;
	return true ;
}

func (t *Tokenizer) peek() rune {
	if t.atEnd() {
		return '\000';
	}
	return rune(t.Source[t.Current]);
}

func (t *Tokenizer) peekNext() rune {
	if t.atEnd() {
		return '\000';
	}
	return rune(t.Source[t.Current + 1]);
}

func (t *Tokenizer) peekToIdx(idx int) rune {
	if t.atEnd() {
		return '\000';
	}
	return rune(t.Source[t.Current + idx]);
}

func (t *Tokenizer) string() {
	for t.peek() != '"' && !t.atEnd() {
		t.advance();
	}
	if t.atEnd() {
		//TODO: Handle error
	}
	t.advance();
	value := t.Source[t.Start + 1 : t.Current - 1];
	t.addToken(String , value);
}

func isDigit (c rune) bool {
	return c >= '0' && c <= '9';
}

//It scan the float numbers and the integet numbers 
func (t *Tokenizer) number() {
	for isDigit(t.peek()) {
		t.advance();
	}

	if t.peek() == '.' && isDigit(t.peekNext()) {
		t.advance();
		for isDigit(t.peek()) {
			t.advance();
		}
	}

	value := t.Source[t.Start:t.Current];
	number , err := strconv.ParseFloat(value , 64);

	if err != nil {
		//TODO: Handle error
	}
	t.addToken(Number , number);
}