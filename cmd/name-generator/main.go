package main

import (
	"fmt"
	namegenerator "github.com/plusalphagroup/namegenerator"
)

func main() {
	// Print random 1000 french first and last names
	for i := 0; i < 1000; i++ {
		firstName, lastName := namegenerator.GetRandomName(namegenerator.France)
		fmt.Println(firstName + " " + lastName)
	}
}
