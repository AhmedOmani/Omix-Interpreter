// Package lexer implements the lexical analysis phase of the Omix interpreter.
package lexer

// TokenType represents the different types of tokens that can be recognized by the lexer.
// Each token type corresponds to a specific lexical element in the source code.
type TokenType int

// Constants defining all possible token types in the language.
const (
	// Single-character tokens
	LeftParen  TokenType = iota // (
	RightParen                  // )
	LeftBrace                   // {
	RightBrace                  // }
	Comma                       // ,
	Dot                         // .
	Minus                       // -
	Plus                        // +
	Semicolon                   // ;
	Slash                       // /
	Star                        // *

	// One or two character tokens
	Bang         // !
	BangEqual    // !=
	Equal        // =
	EqualEqual   // ==
	Greater      // >
	GreaterEqual // >=
	Less         // <
	LessEqual    // <=

	// Literals
	Identifier // variable names, function names, etc.
	String     // string literals
	Number     // numeric literals

	// Keywords
	And    // and
	Class  // class
	Else   // else
	False  // false
	Fun    // fun
	For    // for
	If     // if
	Nil    // nil
	Or     // or
	Print  // print
	Return // return
	Super  // super
	This   // this
	True   // true
	Var    // var
	While  // while

	// Special tokens
	Eof // end of file
)

// String returns a human-readable representation of the token type.
// Used for debugging and error reporting.
func (t TokenType) String() string {
	switch t {
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"
	case Bang:
		return "Bang"
	case BangEqual:
		return "BangEqual"
	case Equal:
		return "Equal"
	case EqualEqual:
		return "EqualEqual"
	case Greater:
		return "Greater"
	case GreaterEqual:
		return "GreaterEqual"
	case Less:
		return "Less"
	case LessEqual:
		return "LessEqual"
	case Identifier:
		return "Identifier"
	case String:
		return "String"
	case Number:
		return "Number"
	case And:
		return "And"
	case Class:
		return "Class"
	case Else:
		return "Else"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case For:
		return "For"
	case If:
		return "If"
	case Nil:
		return "Nil"
	case Or:
		return "Or"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Super:
		return "Super"
	case This:
		return "This"
	case True:
		return "True"
	case Var:
		return "Var"
	case While:
		return "While"
	case Eof:
		return "Eof"
	default:
		return "Unknown"
	}
}

// keywords maps reserved words to their corresponding token types.
// This map is used to identify keywords during lexical analysis.
var keywords = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"fun":    Fun,
	"for":    For,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}
