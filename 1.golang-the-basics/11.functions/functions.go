package main

import "fmt"

func main() {
	fmt.Println(greetings("Bob"))        // => Hello Bob
	fmt.Println(addition(3.2, 4.3, 5.5)) // => 13
	fmt.Println(split(12))               // => 4 8
	fmt.Println(swap(4, 7))              // => 7 4
	fmt.Println(sum(3, 4))               // => 7
	fmt.Println(sumNnumbers(3, 4, 5))    // => 12

	a := f()
	fmt.Println("Closure first run:", a()) // => Closure first run: 10
	a()
	fmt.Println("Closure second run:", a()) // => Closure second run: 30
	a()
	fmt.Println("Closure third run:", a()) // => Closure third run: 50

	// To confirm that the state is unique to that particular function
	// let's create and test a new one.
	b := f()
	fmt.Println("New closure state:", b()) // => New closure state: 10

}

// A basic function that takes a name of type string and returns a greeting message of type string
func greetings(name string) string {
	return "Hello " + name
}

// Still a basic function that takes multiple arguments
func addition(firstNumber, secondNumber, thirdNumber float64) float64 {
	return firstNumber + secondNumber + thirdNumber
}

// Multiple return: Let's write a function that takes a single integer
// and split it into 2 integers such that the sum gives us thesame input integer

func split(x int) (y, z int) {
	y = x * 1 / 3
	z = x - y

	// we could still say return y, z
	return
}

// Similar to above, A function can return any number of results.
func swap(x, y int) (int, int) {
	return y, x
}

// Assigning a function to a variable
// notice this function does not have a name, making it an anonymous function
var sum = func(x, y int) int {
	return x + y
}

// Variadic functions
// We use them when we don't know the number of arguments a function can take
// They are just like any other functions but can accept an unknown number of arguments
// syntax: func funcName(param ...paramType) returnType {}

func sumNnumbers(numbers ...float64) float64 {
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// closures
// A closure is a special type of anonymous function that references variables
// declared outside of the function itself
// It creates a function that has access to the data that persists even after the function exits

func f() func() int {
	x := 0
	return func() int {
		x += 10
		return x
	}
}
