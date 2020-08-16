package main

import "fmt"

type contactInfo struct {
	email string 
	zipCode int 
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	sheldon := person{
		firstName: "Sheldon",
		lastName: "Cooper",
		contact: contactInfo {
			email: "sheldon@cooper.com",
			zipCode: 12345,
		},
	}
	//& means to get the pointer of the variable, here we get pointer for sheldon then assign it to a new variable called sheldonPointer. We use pointer in Go because Go makes a copy of the value and make it available in another function that takes receiver. Once we define the pointer to type as below, we can also write code in a shorter way and let Go take care of the rest.
	// sheldonPointer := &sheldon
	// sheldonPointer.updateName("Amy")
	sheldon.updateName("Amy")
	sheldon.print()
}

//* before a type means this updateName function can only be called with a receiver of a pointer to a person.
func (pointerToPerson *person) updateName(newFirstName string) {
	//* before a pointer means to manipulate the value the pointer is pointing to, note this is different from the above *before the type!
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}


//NOTE IMPORTANT
//Value types: int, float, string, bool, structs ---> use pointer to change these things in functions 
//Reference Types: slices, maps, channels, pointers, functions ---> don't worry about pointers with those 