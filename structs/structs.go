package structs

import (
	"log"
	"time"
)

type myStruct struct {
	FirstName string
}

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func Execute() {
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
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}
