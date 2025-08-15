package main

import "fmt"

// The `fmt` package in Go provides functions for formatted I/O, similar to C's printf and scanf.

// func main() {
	// **Variables**
	// Type of variables in go are declared after the variable unline C/C++.
	// 'rune' is an alias for 'int32' and is used to represent Unicode code points or characters.
	var b float64 = 3.14

	// Println vs Printf
	// Println automatically adds a newline at the end
	// Formatting output requires Printf
	fmt.Println("Hello, World!")

	// **Format Specifiers**
	// Format specifier for pointers: %p
	fmt.Printf("Format specifier if you are unsure of the type: %v\n", 42)
	fmt.Printf("Format specifier for integers: %d\n", 23)
	fmt.Printf("Format specifier for floating point numbers: %f\n", b)
	fmt.Printf("Format specifier for strings: %s\n", "Hello, World!")
	fmt.Printf("Format specifier for booleans: %t\n", true)
	// Others
	fmt.Printf("Format specifier for hexadecimal: %x\n", 255)
	fmt.Printf("Format specifier for octal: %o\n", 255)
	fmt.Printf("Format specifier for scientific notation: %e\n", 123456789.0)
	fmt.Printf("Format specifier for percentage: %%\n")
	fmt.Printf("Format specifier for strings with width: %10.5s\n", "Excuse Mee !!")        // 10 characters wide, 5 characters long
	fmt.Printf("Format specifier for floating point numbers with width: %10.2f\n", 3.14345) // 10 characters wide, 2 decimal places
	fmt.Printf("Format specifier for booleans with width: %5t\n", true)

	// **Code Shortening**
	// Using `:=` for short variable declaration, type is automatically inferred by Go compiler but is still statically typed
	// `a` is a block scoped variable if declared inside a block like `if`, `for`, etc.
	if a := 5; a > 0 {
		fmt.Println("a is positive")
	} else {
		fmt.Println("a is not positive")
	}

	// **Arrays**
	// IMP Note - Slices are passed by reference in Go, so modifying a slice inside a function will modify the original slice/array.
	var arr [5]int // Fixed-size array of integers
	arr[0] = 1
	primes := [5]int{2, 3, 5, 7, 11} // Array with initial values
	slice := primes[1:3]             // Slicing the array to get a sub-array (slice) from index 1 to 2 (exclusive of 3)
	fmt.Println("Array:", arr)
	fmt.Println("Slice of Primes:", slice)
	// Creating a slice
	newSlice := make([]int, 5, 10) // Creates a slice of integers with length 5 and capacity 10
	fmt.Println("New Slice:", newSlice)
	newSlice = append(newSlice, 6, 7, 8) // Appends elements to the slice
	fmt.Println("Appended Slice:", newSlice)
	// Length vs Capacity
	fmt.Printf("New Slice Length: %d, Capacity: %d\n", len(newSlice), cap(newSlice)) // Length of the slice
	for i, v := range newSlice {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	// Matrices
	matrix1 := make([][]int, 0)
	for i := 0; i < 3; i++ {
		row := make([]int, 0)
		for j := 0; j < 3; j++ {
			row = append(row, i*j)
		}
		matrix1 = append(matrix1, row)
	}
	matrix2 := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix2[i][j] = i + j
		}
	}
	fmt.Printf("Matrix 1: %v\nMatrix 2: %v", matrix1, matrix2)

	// Tricky Slices
	// When you create a slice from an array, it shares the same underlying array.
	// IMP - Do not do the following as it can lead to unexpected behavior.
	// Both otherSlice and someSlice will point to the same underlying array. Hence they most be the same slices and not different.
	// otherSlice = append(someSlice, 8) 

	// Example 1
	fmt.Println("\n\nExample 1: Slices and their addresses")
	r1 := make([]int, 3)
	r2 := append(r1, 4)
	r3 := append(r1, 5)
	fmt.Printf("\nLength and Capacity:\nr1: %d, %d\nFor r2: %d, %d\nFor r3: %d, %d", len(r1), cap(r1), len(r2), cap(r2), len(r3), cap(r3))
	fmt.Println("\nArrays:\nr1:", r1, "\nr2:", r2, "\nr3:", r3)
	fmt.Printf("Initial address of r1: %p, r2: %p, r3: %p", &r1[0], &r2[0], &r3[0])
	// Example 2
	fmt.Println("\n\nExample 2: Slices with different lengths and capacities")
	t1 := make([]int, 3, 5)
	t2 := append(t1, 4)
	t3 := append(t1, 5)
	fmt.Printf("Length and Capacity:\nt1: %d, %d\nFor t2: %d, %d\nFor t3: %d, %d", len(t1), cap(t1), len(t2), cap(t2), len(t3), cap(t3))
	fmt.Println("\nArrays:\nf1: ", t1, "\nt2: ", t2, "\nt3: ", t3)
	fmt.Printf("Initial address of t1: %p, t2: %p, t3: %p", &t1[0], &t2[0], &t3[0])
	fmt.Println("\n\nIf you observe carefully in example 1 there is no issue with the addresses of the slices as they are not sharing the same underlying array because the capacity of initial slice t1 is exceeded which requires copying the slice to a new location on creating r2 and r3 from it using append(). And as append function returns the new copied slice, it doesn't cause errors.\nHowever in case of example 2, there was an issue because the capacity of t1 is already more than what t2 and t3 requires so t3 overrides t2. And both t2 and t3 point to the same original slice t1.")

	// **Maps** Like dictionaries in python
	// IMP - Maps are passed by reference in functions
	// Keys is maps should be comparable like integers, strings, structs etc and not slices, maps, functions etc.
	// If a key is not present in the map, it returns the zero value of the value type.
	// If the map doesn't exist, it panics.
	ages := map[string]int{
		"Arjun":  46,
		"Trisha": 23,
	}
	fmt.Println(ages)
	ages["John"] = 30
	age, ok := ages["Arjun"] // ok is a boolean that indicates if the key exists in the map
	if ok {
		fmt.Println("Arjun's age is:", age)
	} else {
		fmt.Println("Arjun's age is not found in the map.")
	}
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	delete(ages, "Trisha")

	// **Conditionals**
	// For loop
	for i := range 5 {
		fmt.Println("Iteration:", i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}
	// While loop
	i := 0
	for i < 5 {
		fmt.Println("While loop iteration:", i)
		i++
	}
	// A for loop with initial and after statements but without the condition is also equivalent to a `while(1):`
	for i := 0; ; i++ {
		if i >= 5 {
			break // Breaks out of the loop
		}
		fmt.Println("Infinite for loop iteration:", i)
	}
	// Switch statement
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	default:
		fmt.Println("It's some other day!")
	}
}