# Omix Command Line Interface

This directory contains the main entry point for the Omix interpreter. The `main.go` file implements the command-line interface that allows users to interact with the interpreter in two ways:

## Main Components

### Omix Struct
```go
type Omix struct {
    hadError bool
}
```
The core struct that maintains the interpreter's state, particularly tracking whether any errors have occurred during execution.

### Command Line Interface
The program supports two modes of operation:

1. **REPL Mode** (Read-Eval-Print Loop)
   - Activated when no arguments are provided
   - Provides an interactive prompt (`>`) for entering code
   - Each line is executed immediately
   - Supports multi-line input
   - Exit with Ctrl+D (EOF)

2. **Script Mode**
   - Activated when a file path is provided as an argument
   - Reads and executes the entire file
   - Exits with appropriate status codes:
     - 65: Compilation error
     - 70: Runtime error

### Error Handling
- Centralized error reporting through the `error()` and `report()` methods
- Errors include line numbers for better debugging
- Non-zero exit codes for script execution errors

## Usage
```bash
# Start REPL
./omix

# Run a script
./omix path/to/script.omix
```

