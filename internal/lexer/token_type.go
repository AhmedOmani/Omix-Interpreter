package main 

type TokenType int ;

const (
	//One signle-character tokens
	LeftParen TokenType = iota; 
	RightParen 
	LeftBrace 
	RightBrace  
	Comma  
	Dot  
	Minus  
	Plus 
	Semicolon 
	Slash 
	Star 

	//One or two character tokens
	Bang 
	BangEqual 
	Equal 
	EqualEqual 
	Greater 
	GreaterEqual 
	Less 
	LessEqual 

	//Literals
	Identifier 
	String 
	Number 

	//Keywords
	And 
	Class
	Else 
	False 
	Func
	For 
	If 
	Nil 
	Or 
	Var 
	This
	True 
	While 
	Super		
	Return 
	Print

	//EOF
	Eof 
)