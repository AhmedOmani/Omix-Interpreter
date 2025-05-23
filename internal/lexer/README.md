# Lexer Package

The lexer package implements the lexical analysis phase of the Omix interpreter. It converts source code into a sequence of tokens that can be processed by the parser.

## Features

### Token Types
- **Single-character tokens**: `(`, `)`, `{`, `}`, `,`, `.`, `-`, `+`, `;`, `/`, `*`
- **Two-character tokens**: `!=`, `==`, `>=`, `<=`
- **Literals**:
  - Numbers (integers, decimals, scientific notation)
  - Strings (with escape sequence support)
  - Identifiers
- **Keywords**: `and`, `class`, `else`, `false`, `fun`, `for`, `if`, `nil`, `or`, `print`, `return`, `super`, `this`, `true`, `var`, `while`

### Error Handling
- Comprehensive error reporting with line and column information
- Specific error messages for:
  - Invalid number formats
  - Unterminated strings
  - Invalid identifiers
  - Unexpected characters

### Number Handling
- Supports multiple number formats:
  - Integers: `123`
  - Decimals: `123.456`, `.123`
  - Scientific notation: `1.23e4`
- Validates number formats:
  - Rejects multiple decimal points
  - Rejects trailing decimal points
  - Rejects invalid scientific notation
  - Rejects letters after numbers

### String Handling
- Supports string literals enclosed in double quotes
- Handles escape sequences:
  - `\n` for newline
  - `\t` for tab
  - `\\` for backslash
  - `\"` for quote

## Usage

```go
// Create a new tokenizer
tokenizer := NewTokenizer(source)

// Scan for tokens
tokens := tokenizer.ScanTokens()

// Check for errors
if tokenizer.HasErrors() {
    for _, err := range tokenizer.GetErrors() {
        fmt.Println(err)
    }
    return
}

// Process tokens
for _, token := range tokens {
    // Handle each token
    fmt.Println(token)
}
```

## Token Structure

```go
type Token struct {
    Type    TokenType    // The type of token
    Lexeme  string       // The original text
    Literal interface{}  // The interpreted value
    Line    int          // Line number
    Column  int          // Column number
}
```

## Error Types

The lexer reports several types of errors:
1. **Invalid Number**: When a number format is invalid
2. **Unterminated String**: When a string literal is not properly closed
3. **Invalid Identifier**: When an identifier starts with a number
4. **Unexpected Character**: When an unrecognized character is encountered

## Testing

The lexer includes comprehensive tests for:
- Token recognition
- Error handling
- Number parsing
- String handling
- Keyword recognition

Run the tests with:
```bash
go test ./internal/lexer
```

## Next Steps

The lexer is now complete and ready for the parser phase. The parser will:
1. Take the sequence of tokens
2. Build an Abstract Syntax Tree (AST)
3. Validate the syntax of the program
4. Prepare for interpretation 