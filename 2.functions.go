package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// **Functions**
	// Note that functions by default in GO are 'call by value' and not 'call by reference' unless you explicitly pass address using `&` like in C/C++.
	add(3, 4) // Calling the add function

	// Using `_` to ignore return values as Go doesn't allow unused variables
	_, num := ignoreReturn()
	fmt.Println("Ignored return value, got:", num)

	// Inline Functions
	result := func(a, b int) (sum int) {
		sum = a + b
		return sum
	}
	fmt.Println("Inline function result:", result(3, 4))

	// Calling the naked return function
	x, y := nakedReturn(3, 4)
	fmt.Println("Naked return values:", x, y)

	// Calling Variadic Functions
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", total)
	// Using spread operator (...) to pass a slice as variadic arguments
	randomSlice := []int{10, 20, 30}
	totalFromSlice := sum(randomSlice...)
	fmt.Println("Sum from slice:", totalFromSlice)

	// Calling Higher Order Functions
	ans := higherOrderFunc(add, 5, 10)
	fmt.Println("Higher Order Function result:", ans)

	// Calling Curried Function
	doubleFunc := curriedDouble(add)
	fmt.Println("Curried Function result:", doubleFunc(5))

	// Using Closures
	countTotal, countSum := makeCounter(), makeCounter()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("Counter 1:", countTotal(1))
			fmt.Println("Counter 2:", countSum(i+j))
		}
	}
	fmt.Printf("Final Counter 1 Value: %d, Final Counter 2 Value: %d\n", countTotal(0), countSum(0))
}

// Functions (not allowed inside main function)
func add(a int, b int) int {
	return a + b
}

func ignoreReturn() (string, int) {
	return "Hello", 5
}

func nakedReturn(a, b int) (x, y int) {
	x = a + b
	y = a - b
	return // naked return, no need to specify return values, returns x and y automatically (called as implicit), not preferred.
}

// Variadic Functions
func sum(numbers ...int) int {
	// ... only means that the function can take a variable number of arguments but inside function it can be treated as a slice.
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Higher Order Functions
// Used in HTTP API handlers, Pub/Sub handlers, onClick callbacks, etc.
func higherOrderFunc(fn func(int, int) int, a, b int) int {
	return fn(a, b) // Calls the passed function with a and b as arguments
}

// Currying
// Used mostly in middleware functions in web frameworks, like Express.js in Node.js.
func curriedDouble(f func(int, int) int) func(int) int {
	return func(x int) int {
		return f(x, x) // Returns a function (called currying)
	}
}

// Defer - Used to ensure that a function call is performed later in a program's execution, usually for purposes like cleanup before a function exits.
func copyFile(src, dst string) error {
	// Defer is used to ensure that the file is closed after the function exits
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.New("Couldn't open file.")
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return errors.New("Couldn't create a file.")
	}

	defer dstFile.Close()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return nil
	} else {
		return errors.New("Couldn't copy file.")
	}
}

// Closures
// Closures are functions that capture the variables from their surrounding context.
func makeCounter() func(x int) int {
	count := 0               // This variable is captured by the closure
	return func(x int) int { // If it hadn't been a function, it would have been a normal variable and not a closure.
		count += x
		return count
	}
}
