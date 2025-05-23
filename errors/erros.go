package errors;

import "fmt" ;

type BaseError struct {
	Line 		int
	Column		int
	Message 	string
	Phase 		string
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("[%s]: ErrorMessage: %s at line %d, column %d", e.Phase, e.Message, e.Line, e.Column)
}
