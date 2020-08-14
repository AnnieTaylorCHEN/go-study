package main //defines a package that can be complied and then executed, must have a func named main

import "fmt"//import other packages such as helper functions, fmt stands for format

func main(){
	fmt.Println("Hi there!") //declare functions, tell Go to do things
}

//Exercises

//what is the purpose of a package in Go?
// to group together code with a similar purpose 

//what is the special name for a package to tell Go that we want it to be turned into a fle that can be executed?
//main

//the one requirement of packages named "main" is that we...
//define a function named "main", which is ran automatically when the program runs

//why do we use "import" statements?
//to give our packages access to code written in another package