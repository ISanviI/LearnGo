package main

type car struct {
	Company string
	Price   int
}

func (c car) updatePrice(newPrice int) {
	c.Price = newPrice
}

func (c *car) updateCompany(newCompany string) {
	c.Company = newCompany
}

// func main() {
	// **Pointers**
	// Pointers in Go are similar to C/C++ pointers, but they are safer and easier to use.
	// Deferencing a nil pointer could cause severe bugs and panics.
	var a int = 42
	var p *int = &a // p is a pointer to the variable a
	println("Value of p (address of a):", p)
	println("Value pointed by p:", *p)
	*p = 100
	println("New value of a after modifying through pointer p:", a)

	c1 := car{
		Company: "Toyota",
		Price:   30000,
	}
	// No need to use `&` to pass the struct to a function, as it is passed by reference by default.
	c1.updatePrice(35000)     // This will not change the original car's price as it is passed by value
	c1.updateCompany("Honda") // This will change the original car's company as it is passed by reference
	println("Car after updatePrice:", c1.Company, "Price:", c1.Price)
}
