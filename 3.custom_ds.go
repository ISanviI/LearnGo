package main

import "fmt"

// **Structs** - Group related data together.
// Nested structs (another structure as a member in the current struct while specifying both a different variable name for it and the name of another struct) are also possible in Go.
type Person struct {
	Name string
	Age  int
}
// Methods on structs
func (p *Person) updateAge(age int) { // Method with receiver type Person
	p.Age = age
	Greet(*p)
}

func Greet(p Person) int {
	fmt.Println("Hello, my name is " + p.Name)
	fmt.Println("I am", p.Age, "years old")
	return 1
}

// **Interfaces** - Define a contract that types can implement, more like a blueprint for methods.
// For example like abstract classes in C++ (methods must be redefined in child classes during inheritance, provides a blueprint)
type Shape interface {
	Area() float64 // Method signature
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius // Implementing the Area method for Circle
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height // Implementing the Area method for Rectangle
}
func PrintArea(s Shape) { // Function that takes an interface type
	// Type Assertions
	if circle, ok := s.(Circle); ok {
		fmt.Println("Circle Area:", circle.Area()) // Calls the Area method of the Shape interface
	} else if rectangle, ok := s.(Rectangle); ok {
		fmt.Println("Rectangle Area:", rectangle.Area())
	} else {
		fmt.Println("Unknown shape type")
	}
}

// func main() {
	// **Structs**
	// Creating a new Person instance
	person := Person{Name: "Alice", Age: 30}
	Greet(person)
	person.updateAge(35)

	// Anonymous Structs - Useful for quick, one-off data structures. (Useful for nested structs)
	anonymous := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  25,
	}
	fmt.Println("Anonymous Struct - Name:", anonymous.Name, ", Age:", anonymous.Age)

	type CarT1 struct {
		Company string
		Model   string
		Price   int
		Wheel   struct { // Anonymous Nested struct
			Radius   float32
			Material string
		}
	}

	type Wheel_A struct {
		Width    float32
		Diameter float32
	}

	type CarT2 struct {
		Publish_Year int
		Mileage      int
		Wheel        Wheel_A // Nested - Width of wheel of an instance of CarT2 (car2) called as `car.Wheel.Width`
	}

	// Embedding Structs - Embedding allows you to include one struct within another, directly inheriting its fields and methods.
	type CarT3 struct {
		Fuel    string
		Type    string
		Wheel_A // Embedded - Width of wheel of an instance of carT3 (car3) called as `car.Width`
	}

	car := CarT3{
		Fuel: "Diesel",
		Type: "SUV",
		Wheel_A: Wheel_A{ // Both parameter and struct name should be same
			Width:    225.5,
			Diameter: 17.0,
		},
	}
	fmt.Println("Car wheel's width (using car.width):", car.Width)

	// **Interfaces**
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}
	// PrintArea() can have any struct as an argument that implements the Shape interface
	// This is polymorphism in Go, where different types can be treated as the same type (interface) if they implement the same methods.
	PrintArea(circle)    // Calls the Area method for Circle
	PrintArea(rectangle) // Calls the Area method for Rectangle

}
