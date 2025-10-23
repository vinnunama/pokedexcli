
package main

import "strings"

func cleanInput(text string) []string {
	//The purpose of this function will be to split the user's input into "words" based on whitespace. 
	// It should also lowercase the input and trim any leading or trailing whitespace.
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	
	words := strings.Fields(text)
	return words
}