package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// WordCount :
//	- Total number of words
//	- A map that records individual words with their total frequency
type WordCount struct {
	TotalWords int            `json:"count"`
	FreqMap    map[string]int `json:"words"`
}

// TODO. Support other alphabets outside english but within unicode, e.g. French / Greek / Vietnamese
//	Excluding Eastern-Asian characters since it's more complicated to decide the word composition

// Given a string, record total count of words (case-insensitive)
// Counted words will all be returned in lower-case
func computeFrequency(s string) map[string]int {
	freqMap := make(map[string]int)
	// Use a word buffer to keep track of partial words
	wordBuffer := make([]string, 0)
	var fullWord string
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
	wordCount := WordCount{
		TotalWords: len(freqMap),
		FreqMap:    freqMap,
	}
	b, err := json.Marshal(wordCount)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	b, err := makeJSON(computeFrequency("aasdasda"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", string(b))
}
