package main ;

import (
	"bufio"
	"fmt"
	"os"
)

type Omix struct {
	hadError bool;
};

func NewOmix() *Omix {
	return &Omix{
		hadError: false,
	}
}

func main() {
	omix := NewOmix();
	
	if len(os.Args) > 2 {
		fmt.Println("Usage: omix [script]")
		os.Exit(64);
	} else if len(os.Args) == 2 {
		omix.runFile(os.Args[1]);
	} else {
		omix.runPrompt();
	}
}

func (o *Omix) runFile(path string) error {
	bytes, error := os.ReadFile(path);
	if error != nil {
		return error;
	}
	content := string(bytes);
	o.run(content);

	if o.hadError {
		os.Exit(65);
	}

	return nil;
}

func (o *Omix) runPrompt() {
	reader := bufio.NewReader(os.Stdin);

	for {
		fmt.Print("> ");
		line, err := reader.ReadString('\n');
		if err != nil {
			break;
		}
		o.run(line);
		o.hadError = false;
	}
}

func (o *Omix) run(source string) {
	//TODO: Create a tokenizer and tokenize the source code.
	// For now just print the source code
	fmt.Println(source);
	o.hadError = false;
}

func (o *Omix) error(line int , message string) {
	o.report(line , "" , message);
}

func (o *Omix) report(line int , where string , message string){
	fmt.Printf("[line %d] Error %s: %s\n" , line , where , message);
	o.hadError = true ;
}