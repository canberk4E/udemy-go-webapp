package pointers

import "log"

func Execute() {
	// var colorGreen string
	// colorGreen = "Green"
	colorGreen := "Green"

	log.Println("colorGreen is set to", colorGreen)

	changeUsingPointer(&colorGreen)

	log.Println("after func call colorGreen is set to", colorGreen)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
