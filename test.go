// A package is a collection of files and codes
// A main package means that our code should be compiled into
// an executable program at the end. Hence, go will create a stand
// alone executable file for us if we choose to build the app
// Running that file will start our program
// Since we will be writing a program that will run on
// our computer, we shall be using pacakge main
package main

//  Importing FMT
//  fmt is a package from the go std lib.
// The std lib contains packages for different functionality.
// fmt is used to formatting strings and printing messages.
import "fmt"

// Out function declaration starts with 'func'
// The name of this function is "main", this particular function is the entry
// point of our app.
// When when run our application, go will look for files in our program
// and it will look for this main function. Then, it will fire this function
// automatically.
// Law leeha esm tany, msh hat-et-run. We must have only one main func
// in our application
func main() {
	// fmt.<method_name>
	// A method starts with a capital letter
	// This method is called print line
	fmt.Println("Hello, ninjas!")

}
