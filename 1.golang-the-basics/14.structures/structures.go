package main

import "fmt"

func main() {
	// Go’s structs are typed collections of fields.
	// They’re useful for grouping data together to form records.

	// Here we create a new type called Animal, which is a structure
	// and give it some properties

	type Animal struct {
		name string
		age  int
	}

	// To create an instance of Animal we do as follows
	// We are creating a new Animal (dog) with name bobby and age 3

	dog := Animal{name: "Bobby", age: 3}
	fmt.Println(dog) // => {Bobby 3}

	// To get a specific value like the age we can use the dot (.)
	fmt.Println("Bobby is:", dog.age, "years old") // => Bobby is: 3 years old

	// we can equally use the new keyword to create a new instance of Animal

	var cat = new(Animal)
	cat.name = "Chippy"
	cat.age = 2

	fmt.Println(cat.name, "is", cat.age, "Years old")

}
