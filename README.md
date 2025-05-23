# Omix Interpreter

A modern interpreter implementation in Go, following the principles from "Crafting Interpreters" by Robert Nystrom.

## Project Status

### Completed Features
- ✅ Lexical Analysis (Tokenizer)
  - Token recognition for all language elements
  - Support for numbers (integers, decimals, scientific notation)
  - String literals with error handling
  - Keywords and identifiers
  - Operators and punctuation
  - Comprehensive error reporting
  - Interactive REPL mode
  - File execution mode

### Current Phase
- 🔄 Parser Implementation (In Progress)
  - Expression parsing
  - Statement parsing
  - Abstract Syntax Tree (AST) construction

## Project Structure

```
omix-interpreter/
├── cmd/
│   └── omix/             # Command-line interface
├── internal/
│   ├── lexer/            # Lexical analysis
│   ├── parser/           # Syntax analysis (coming soon)
│   ├── ast/              # Abstract Syntax Tree (coming soon)
│   
└── examples/             # The code of omix language
└── errors/               # Error handling
└── test/                 # Unit Testing
```

## Getting Started

### Prerequisites
- Go 1.16 or higher
- Git

### Installation
```bash
git clone https://github.com/yourusername/omix-interpreter.git
cd omix-interpreter
go build -o omix cmd/omix/main.go
```

### Usage

1. Run a file:
```bash
./omix path/to/script.omix
```

2. Start interactive REPL:
```bash
./omix
```

## Language Features

### Current Support (Lexer Phase)
- Numbers:
  - Integers: `123`
  - Decimals: `123.456`, `.123`
- Strings: `"hello world"`
- Keywords: `if`, `else`, `while`, `for`, `var`, etc.
- Operators: `+`, `-`, `*`, `/`, `==`, `!=`, etc.
- Identifiers: Variable and function names

### Coming Soon (Parser Phase)
- Expression parsing
- Statement parsing
- Control flow structures
- Function declarations
- Class definitions

## Development

### Running Tests
```bash
go test ./...
```

## Acknowledgments
- "Crafting Interpreters" by Robert Nystrom
- Go programming language community 