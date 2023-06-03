package data_structures

import "log"

func Execute() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	log.Println(myMap["dog"])

	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	log.Println(mySlice)
}
