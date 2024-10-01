package main

import (
	"bytes"
	"encoding/json"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
)

type Response struct {
	translatedText string
	error          *string
}

func Translate(cCtx *cli.Context, translateLink string, translateApiKey string, text string, fromLanguage string, toLanguage string) string {
	// Return if the text is empty to save on resources
	if len(text) == 0 {
		return ""
	}

	// Create the Http Client
	httpClient := &http.Client{}

	// Create the body and encode to reader
	resJson, reqBody := map[string]string{"q": text, "source": fromLanguage, "target": toLanguage}, new(bytes.Buffer)
	err := json.NewEncoder(reqBody).Encode(resJson)
	if err != nil {
		log.Fatal(err)
		return text
	}

	// Create the request
	req, err := http.NewRequest(http.MethodPost, translateLink, reqBody)
	if err != nil {
		log.Fatal(err)
		return text
	}

	// Set Header
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	res, postErr := httpClient.Do(req)
	if postErr != nil {
		log.Fatal(postErr)
		return text
	}

	// Read the response body
	resBody, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return text
	}

	// Convert response to json
	response := Response{}
	jsonErr := json.Unmarshal(resBody, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return text
	}

	// Display error if translation ran into error
	if response.error != nil {
		log.Fatal(response.error)
		return text
	}

	// Return final translation
	return response.translatedText
}
