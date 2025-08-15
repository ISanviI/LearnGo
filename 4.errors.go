package main

import (
	// "errors"
	"fmt"
)

// **Error Handling** - Go uses multiple return values to handle errors.

// func main() {
	// Example of error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Result:", result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, fmt.Errorf("DivisionError - %w", errors.New("Division by zero is not allowed"))
		// Alternatively, more correct way while using the Error interface is by creating a custom error struct type:
		return 0, DivisionError{Message: "Division by zero is not allowed."}
	}
	return a / b, nil
}

// DivisionError is a custom error type
type DivisionError struct {
	Message string
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("DivisionError: %s", e.Message)
}
