package main

import (
	"fmt"
	"time"
)

type cancatable interface {
	String() string
}

type comparable interface {
	~string | ~int | ~float64
}

// Parameterized interfaces - Life
// Individual satisfies ProperNoun because it has Name() string.
// Greeter satisfies Life[Individual] because it implements Action(Individual).
type Life[T ProperNoun] interface {
	Action(T)
}
type ProperNoun interface {
	Name() string
}
type Individual struct {
	name      string
	birthYear int
}

func (p Individual) Name() string {
	return p.name
}

type CalAge struct {
	age int
}

func (a *CalAge) Action(p Individual) {
	a.age = time.Now().Year() - p.birthYear
}

// func main() {
	// Generics allow you to write functions and data structures that can work with any data type.
	// This is useful for creating reusable code that can handle different types without duplication.
	// Very similar to templates in C++ or generics in Java.

}

func Print[T any](value T) {
	fmt.Printf("%v\n", value)
}

func concat[T cancatable](a, b T) string {
	return a.String() + b.String()
}

func compare[T comparable](a, b T) bool {
	return a == b
}
