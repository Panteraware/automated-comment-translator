package main

import "regexp"

func Format(text string) string {
	// Removes all special characters from string
	re, _ := regexp.Compile("/[^\\w\\s]/gi")

	text = re.FindString(text)

	// Still more replacements left

	return text
}
