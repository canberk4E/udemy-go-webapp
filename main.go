package main

import (
	"time"

	"github.com/canberk4e/demoprogram/channels"
	"github.com/canberk4e/demoprogram/interfaces"
	"github.com/canberk4e/demoprogram/jsonhandling"
	"github.com/canberk4e/demoprogram/pointers"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	pointers.Execute()
	interfaces.Execute()
	channels.Execute()
	jsonhandling.Execute()

	/* 	var myType helpers.SomeType
	   	myType.TypeName = "TYPE"
	   	myType.TypeNumber = 1

	   	fmt.Println(myType.TypeName)
	   	fmt.Println(myType.TypeNumber)
	*/

	/*
		var myVar myStruct
		myVar.FirstName = "John"

		myVar2 := myStruct{
			FirstName: "Mary",
		}

		log.Println("myVar is set to", myVar.printFirstName())
		log.Println("myVar2 is set to", myVar2)

		user := User{
			FirstName:   "Trevor",
			LastName:    "Sawler",
			PhoneNumber: "1 555 555-1212",
		}

		log.Println(user.FirstName)
		log.Println(user.LastName)
		log.Println(user.BirthDate)

		myMap := make(map[string]string)
		myMap["dog"] = "Samson"
		log.Println(myMap["dog"])

		var mySlice []string
		mySlice = append(mySlice, "Trevor")
		mySlice = append(mySlice, "John")
		log.Println(mySlice)
	*/

	/*
		fmt.Println("Hello world!")

		var whatToSay string
		var i int

		whatToSay = "Goodby, cruel world!"
		fmt.Println(whatToSay)

		i = 10
		fmt.Println("i is set to", i)

		whatWasSaid, theOtherThing := saySomething()

		fmt.Println("The function returned", whatWasSaid, theOtherThing)


	*/
}

func saySomething() (string, string) {
	return "Something", "else"
}

func sayAnotherThing(s string) (string, string) {
	return s, "World"
}
