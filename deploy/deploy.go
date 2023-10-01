package main

import "fmt"

func incrementRelease() {
	// Define and initialize the variable Release with a default value of 1
	var Release = 1

	// Increase the value of Release by 1
	Release++

	// Print the new value of Release
	fmt.Printf("Release: update to %d\n")
}

func main() {
	// Call the function to increment the value of Release and print the new value
	incrementRelease()
}
