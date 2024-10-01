package main

import (
	"strings"
)

func Format(text string) string {
	array := AcceptedCommentSyntax()

	for i := 0; i < len(text); i++ {
		text = strings.ReplaceAll(text, array[i], "")
	}

	return text
}
