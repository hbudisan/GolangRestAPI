package main

import "fmt"

func variable_declaration() {
	// Declaring Variable
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	// Multiple Declarations
	var (
		employeeId          int    = 5
		firstName, lastName string = "Uzumaki", "Naruto"
	)
	fmt.Println(employeeId, firstName, lastName)

	// Short Declarations
	name := "Bill Gates"
	age, salary, isProgrammer := 35, 50000.0, true
	fmt.Println(name, age, salary, isProgrammer)
}

func type_inference() {
	// Type inference
	var name = "Steve Jobs"
	fmt.Printf("Variable 'name' typenya %T\n", name)

	// multiple variable declaration
	var firstName, lastName, age, salary = "James", "Bond", 30, 70000.0
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T \n", firstName, lastName, age, salary)
}

func constant_declaration() {
	// Untyped Constant
	const myFavLanguage = "Kotlin"
	const sunRisesInTheEast = true

	// Typed Constant
	const typedInt int = 100
	const typedStr string = "Hi"

	// Multiple Declaration
	const country, code = "Indonesia", 62
	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	fmt.Println(myFavLanguage, sunRisesInTheEast, typedInt, typedStr, country, code, employeeId, salary)
	//salary = 60000
}

func ioata_declaration() {
	const (
		first = iota
		second
	)

	const (
		third = iota
		four
	)
	fmt.Println(first, second, third, four)
}

func pointer_declaration() {
	var a = 5.67
	var p1 = &a
	var p2 *float64 = &a

	fmt.Println("Nilai yg disimpan variable a: ", a)
	fmt.Println("Alamat memory a: ", &a)
	fmt.Println("Nilai yg disimpan variable p1: ", p1)
	fmt.Println("Nilai yg disimpan variable p2: ", p2)
	a = 7.77
	fmt.Println("Nilai yg disimpan variable a: ", *p2)
}

func main() {
	fmt.Println("Primitive Data Types")
	// variable_declaration()
	// type_inference()
	// constant_declaration()
	// ioata_declaration()
	pointer_declaration()
}
