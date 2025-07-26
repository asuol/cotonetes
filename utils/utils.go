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
