package main

import (
	"fmt"

	"github.com/canberk4e/demoprogram/channels"
	"github.com/canberk4e/demoprogram/data_structures"
	"github.com/canberk4e/demoprogram/helpers"
	"github.com/canberk4e/demoprogram/interfaces"
	"github.com/canberk4e/demoprogram/json_handling"
	"github.com/canberk4e/demoprogram/pointers"
	"github.com/canberk4e/demoprogram/structs"
	"github.com/canberk4e/demoprogram/testing"
	"github.com/canberk4e/demoprogram/variables_functions"
)

func main() {
	fmt.Println("Hello world!")

	// using packages -> helpers
	var myType helpers.SomeType
	myType.TypeName = "TYPE"
	myType.TypeNumber = 1

	fmt.Println(myType.TypeName)
	fmt.Println(myType.TypeNumber)
	// ---------------------------- //

	variables_functions.Execute()
	structs.Execute()
	data_structures.Execute()
	pointers.Execute()
	interfaces.Execute()
	channels.Execute()
	json_handling.Execute()
	testing.Execute()
}
