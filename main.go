package main

import (
	"fmt"
	"log"
	"time"

	"github.com/canberk4e/demoprogram/helpers"
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

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}

func main() {
	var myType helpers.SomeType
	myType.TypeName = "TYPE"
	myType.TypeNumber = 1

	fmt.Println(myType.TypeName)
	fmt.Println(myType.TypeNumber)

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 32,
	}

	PrintInfo(&gorilla)
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

		var colorGreen string
		colorGreen = "Green"

		log.Println("colorGreen is set to", colorGreen)
		changeUsingPointer(&colorGreen)
		log.Println("after func call colorGreen is set to", colorGreen)
	*/
}

func saySomething() (string, string) {
	return "Something", "else"
}

func sayAnotherThing(s string) (string, string) {
	return s, "World"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
