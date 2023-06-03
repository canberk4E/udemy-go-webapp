package variables_functions

import "fmt"

func Execute() {
	var whatToSay string
	var i int

	whatToSay = "Goodby, cruel world!"
	fmt.Println(whatToSay)

	i = 10
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThing := saySomething()

	fmt.Println("The function returned", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "Something", "else"
}

/*
func sayAnotherThing(s string) (string, string) {
	return s, "World"
}
*/
