package main

import "fmt"

// Go support embedding of structs and interfaces
// to express a more seamless composition of types.

// In order word The Go programming language does not support inheritance
// but instead it supports composition via “embedding.”

// for example a structure can be embedded into another structure
// in such a way that the fields and Methods of the embedded structure
// can be access by it's parent structure directly

type Child struct {
	name string
	age  int
}

func (ch Child) sayHello() string {
	return "Hello my name is " + ch.name
}

// let's embed the child struct in the parent struct
// An embedding looks like a field without a name.

type Parent struct {
	Child
	married bool
}

func main() {

	// 	When creating structs with literals,
	// we have to initialize the embedding explicitly;
	// here the embedded type serves as the field name.

	p := Parent{
		Child: Child{
			name: "Mary",
			age:  3,
		},
		married: true,
	}
	// We can access the child’s fields directly on p, e.g. p.age
	fmt.Println("My child is", p.age, "years old")

	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("My child's name is", p.Child.name)

	// Since parent embeds child, the methods of child also become methods of parent.
	// Here we invoke a method that was embedded from child directly on p.

	fmt.Println(p.sayHello())

	// Since parent embeds child, the methods of child also become methods of parent.
	// Here we invoke a method that was embedded from child directly on p.

	// Notice that child already implements
	// the method sayHello() in the greeter interface
	type greeter interface {
		sayHello() string
	}

	var g greeter = p
	fmt.Println("Greet:", g.sayHello())
}
