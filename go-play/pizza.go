package main

import "fmt"

func main() {
	fmt.Println("Do you like Pizza? \n")
	var likespizza string
	fmt.Scanln(&likespizza)

	if (likespizza) == "no" {
		fmt.Println("Thank you for your response")
	} else {
		fmt.Println("What kind of Piza do you like?")
		fmt.Println("i like pizza cheese")
	}

}
