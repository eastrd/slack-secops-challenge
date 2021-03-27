package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"
)

// Given a string, record total count of words (case-insensitive)
// Counted words will all be returned in lower-case
func computeFrequency(s string) map[string]int {
	var fullWord string

	// Use a buffer to track partial words
	wordBuffer := make([]string, 0)
	freqMap := make(map[string]int)

	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			if len(wordBuffer) > 0 {
				//	If it's not a letter
				//		then add word buffer into frequency map & increment its count
				fullWord = strings.Join(wordBuffer[:], "")
				if _, ok := freqMap[fullWord]; ok {
					freqMap[fullWord]++
				} else {
					freqMap[fullWord] = 1
				}
				wordBuffer = nil
			}
		} else {
			// 	If it's a letter
			//		then keep track of its lowercase letter into a word buffer
			wordBuffer = append(wordBuffer, string(unicode.ToLower(r)))
		}
	}

	// Flush the buffer after iterating
	if len(wordBuffer) > 0 {
		fullWord = strings.Join(wordBuffer[:], "")
		if _, ok := freqMap[fullWord]; ok {
			freqMap[fullWord]++
		} else {
			freqMap[fullWord] = 1
		}
	}

	return freqMap
}

// Convert a frequency map into struct
//	which includes additional data on total count of non-repeating words
func makeJSON(freqMap map[string]int) ([]byte, error) {
	wordFreq := WordFrequency{
		TotalWords: len(freqMap),
		FreqMap:    freqMap,
	}

	b, err := json.Marshal(wordFreq)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// API endpoint for calculating word frequency
// Only accept POST as the input can be greater than 2MB
func handleGetWordFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case "POST":
		var userInput UserInput
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userInput)
		if err != nil {
			// JSON Decode failed, i.e. User sent an invalid request (Not in JSON)
			//	Reply 400 Bad Request
			log.Printf("error converting request to JSON: %s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("received: %s\n", userInput.InputSentence)

		b, err := makeJSON(computeFrequency(userInput.InputSentence))
		if err != nil {
			log.Printf("error making JSON for input: %s, error detail: %s", userInput.InputSentence, err.Error())
		}
		w.Write(b)

	default:
		// Default to 404 Not Found for non-POST requests
		w.WriteHeader(http.StatusNotFound)
	}
}
