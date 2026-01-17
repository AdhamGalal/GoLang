// This file that I have created is to get familiar with variables

package main

import "fmt"

func main() {

	// strings
	// To set a variable
	// 1. Write "var"
	// 2. Set the variable name
	// 3. Set the variable type.
	// 4. Set the value of the variable
	// Note: In go, if you are defining a variable that is not used it will print an error
	var nameOne string = "Adham"

	// Another way to set a variable
	// 1. Write "var"
	// 2. Write the name of the variable
	// 3. Give the variable a value and go will automatically define the type of the variable
	// Note: the type of this variable cannot be changed later on
	var nameTwo = "Galal"
	// Third way
	// Define the variable name along with the data type and assingn it later
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	// If we need to change the value of the variable, then it is fine.
	nameOne = "Whatever."
	nameTwo = "Do what ever you want"
	fmt.Println(nameOne, nameTwo, nameThree)
	// Another way to initialize the variable
	// 1. write the variable name.
	// 2. write ":="
	// 3. Set the variable value
	// This method is only to intilalize variables.
	// Note: this cannot be used outside of a "func"
	// Other methods can be used outside of the func
	nameFour := "Hi"
	fmt.Println(nameFour)

	// We have two different datatypes for numbers; one for ints and another for floats
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// We can specify the number of bits to for the integer to be stored.
	// It can be 8, 16 , 32, and 64
	// Note that it assumes that the int is signed, so int8 can store up to 2^7 -1 for +ve numbers and 2^7 for -ve numbers
	var numFour int8 = 50
	// By default int matches the architecture of the CPU
	fmt.Println(numFour)

	// An unsigned int is "uint" -> canot have a negative number. We can also specify the bit width
	// dont use the size unless you are certian about the size.

	// Floats

	//var scoreOne float32 = 25.98
	//var scoreTwo float64 = 3246534.321321

	// For most of the times we will be using floaat 64.
	// If we used the default initialization format it will use the 64 by default

}
