package utils

import (
	"fmt"
)

func User_confirmation(msg string) bool {
	var i string
	fmt.Println(msg, " [y/n]")
	fmt.Scan(&i)
	if i != "y" && i != "n" {
		fmt.Println("Please answer with \"y\" or \"n\"")
		return User_confirmation(msg)
	}

	return i == "y"
}

func Trim_note(note []string) []string {
	var start_index int

	for start_index = range note {
		if note[start_index] != "" {
			break
		}
	}

	// Ensure no empty lines at start of note text
	if start_index > 0 {
		note = note[start_index:]
	}

	var end_index int

	for end_index = len(note); end_index >= 1; end_index-- {
		if note[end_index-1] != "" {
			break
		}
	}

	// Ensure no empty lines at end of note text
	if end_index < len(note) {
		note = note[:end_index]
	}

	return note
}
