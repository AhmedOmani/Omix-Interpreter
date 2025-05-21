This is the high-level modula architecture for Omix lang interpreter,
It divided to layers each layer will be responsible for one phase 

1- Frontend Layer:
    * Handles all source code processing.
    - Components:
        * Lexer: Tokentize source code.
        * Parser: Builds AST.
        * SemanticAnalyzer: Validates program semantics.

2- Core Engine Layer :
    - Components:
        * Interpreter: Main orchestrator that coordinates all components.
        * Runtime: Manages program state, memory, and execution context.
        * ErrorHandler: Centralized error management.

3- AST Layer:
    * Represent program structure.
    - Components:
        * Node: Base interface for all AST nodes.
        * Expression.
        * Statments
    * This is the most important layer to achieve extensibility.

4- Runtime Layer
    * Manages program execution
    -Components:
        * Environment: Variable scope management
        * MemoryManager: Memory allocation and garbage collection
        * Value: Represents runtime values

5- Standard Library Layer
    * Built-in functionality
    - Components:
        * Builtins: Core functions
        * IO: Input/output operations

--> Each layer communicates through well-defined interfaces.
--> Components dont know about each others implemntation.
--> Makes it easy to swap implementations.

Project architecture:

omix-interpreter/
├── cmd/
│   └── omix/
│       └── main.go             # Entry point 
│
├── internal/
│   ├── lexer/            
│   │   ├── token.go            # Token definition
│   │   └── tokenizer.go        # Tokenizer/Scanner implementaion 
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
|
├── examples/                   # Programs written in our language
│   ├── hello.omix              # Example: Hello World program
