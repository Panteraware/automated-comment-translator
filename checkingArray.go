package main

import "slices"

type CheckingArrayResponse struct {
	Matches bool
	Value   *string
}

func CheckingArray(text []string) CheckingArrayResponse {
	response := &CheckingArrayResponse{}

	array := AcceptedCommentSyntax()

	for i := 0; i < len(array); i++ {
		if slices.Contains(text, array[i]) {
			response.Matches = true
			response.Value = &array[i]
			return *response
		}
	}

	response.Matches = false
	return *response
}
