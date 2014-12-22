package main

import "fmt"

//celsius to Fahrenheit converter
//takes a celsius in float64 as a paramete//rturning a float64 after calculations have been made
func CelToFah(value float64) float64 {
	var fah float64
	fah = value*9/5 + 32
	return fah
}
func main() {
	//declare the users input
	var userInput float64
	fmt.Print("Enter your Celsius to be converted to Fahrenheit ")
	//gets the input from the user
	fmt.Scanf("%f", &userInput)
	//format the output from the celsius to Fahrenheit conversion
	fmt.Printf("Your Fahrenheit is %g F\n", CelToFah(userInput))

}
