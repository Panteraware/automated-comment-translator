package main

import (
	"strings"
)

type CheckingStringResponse struct {
	Matches bool
	Value   *string
}

func CheckingString(text string) CheckingStringResponse {
	response := &CheckingStringResponse{}

	array := AcceptedCommentSyntax()

	text = strings.TrimSpace(text)

	for i := 0; i < len(array); i++ {
		if strings.HasPrefix(text, array[i]) || strings.Contains(text, array[i]) {
			response.Matches = true
			response.Value = &array[i]
			return *response
		}
	}

	response.Matches = false
	return *response
}
