package main

import "fmt"

func main() {
	// Maps are Go’s built-in associative data type
	// sometimes called hashes or dicts in other languages.

	//you create a mapby using the make function followed by the map keyword
	//specifying the the key type in square braces and the value type out of the braces
	students := make(map[int]string)

	// Set key/value pairs using typical name[key] = val syntax.
	students[1] = "Jacob"
	students[2] = "Ralph"
	students[3] = "Perry"

	fmt.Println(students)                                   // => map[1:Jacob 2:Ralph 3:Perry]
	fmt.Println("The student with id = 1 is:", students[1]) // => The student with id = 1 is: Jacob

	// use the builtin len() function to return the length of the map

	fmt.Println(len(students)) // => 3

	// The builtin delete removes key/value pairs from a map.
	// here we remove the 1st student (Jacob) by specifying the map(student) and the id (key)
	// as parameters to the delete function

	delete(students, 1) // => The new students map is: map[2:Ralph 3:Perry]
	fmt.Println("The new students map is:", students)

	// Here we are trying to print a student that does not exist (student with id = 1)
	// We didn’t need the value itself, so we ignored it with the blank identifier _.
	_, stud := students[1]
	fmt.Println("The student with id 1 is:", stud) // => The student with id 1 is: false
}
