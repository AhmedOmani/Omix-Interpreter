# Omix Interpreter

A modular, extensible interpreter implementation in Go, following the principles from "Crafting Interpreters". This project demonstrates the implementation of a programming language interpreter with a clean, layered architecture.

## Architecture Overview

The interpreter follows a modular design with clear separation of concerns. Here's how the components are organized:

### 1. Command Line Interface (`cmd/omix/`)
- Entry point for the interpreter
- Handles user interaction through:
  - REPL (Read-Eval-Print Loop) mode
  - Script file execution mode
- Manages error reporting and exit codes

### 2. Core Components (`internal/`)
#### Lexer (`internal/lexer/`)
- `token.go`: Defines token types and structures
- `tokenizer.go`: Implements source code tokenization
- Converts source code into a stream of tokens

#### Parser (`internal/parser/`)
- `ast.go`: Defines Abstract Syntax Tree structures
- `parser.go`: Implements parsing logic
- Converts tokens into an AST

#### Interpreter (`internal/interpreter/`)
- `interpreter.go`: Core interpretation logic
- Executes the AST
- Manages program flow

#### Runtime (`internal/runtime/`)
- `environment.go`: Manages variable scopes
- `values.go`: Defines runtime value types
- Handles memory management

### 3. Error Handling (`pkg/error/`)
- Centralized error management
- Consistent error reporting
- Error recovery strategies

### 4. Examples (`examples/`)
- Sample programs demonstrating language features
- Test cases and usage examples
- Learning resources

## Key Design Principles
- Each component is self-contained and has a single responsibility
- Clear interfaces between components
- Easy to test and extend
- Consistent error handling throughout

## Project Structure
```
omix-interpreter/
├── cmd/
│   └── omix/
│       └── main.go             # Entry point
│
├── internal/
│   ├── lexer/            
│   │   ├── token.go            # Token definition
│   │   └── tokenizer.go        # Tokenizer/Scanner implementation
│   │
│   ├── parser/             
│   │   ├── ast.go              # AST definitions
│   │   └── parser.go           # Parser implementation
│   │
│   ├── interpreter/            # The main interpreter
│   │   └── interpreter.go
│   │
│   └── runtime/                # Runtime features
│       ├── environment.go
│       └── values.go
│
├── pkg/
│   └── error/                  # Error handling
│
├── examples/                   # Example programs
│   └── hello.omix              # Hello World example
```

## Getting Started
```bash
# Run in REPL mode
./omix

# Execute a script file
./omix path/to/script.omix
```

## Development Status
🚧 Under active development - Currently implementing the lexer and parser components. 