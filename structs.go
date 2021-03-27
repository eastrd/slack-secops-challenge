package main

// WordFrequency :
// - Total number of words
// - A map that records individual words with their total frequency
type WordFrequency struct {
	TotalWords int            `json:"count"`
	FreqMap    map[string]int `json:"words"`
}

// UserInput :
// - Struct for user input JSON
type UserInput struct {
	InputSentence string `json:"input"`
}
