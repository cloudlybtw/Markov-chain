package main

import (
	"markov-chain/pkg/input"
	"markov-chain/pkg/output"
)

func main() {
	numWordsFlag, lenFlag, prefixFlag := input.ParseFlags()         // Number Of Words; Prefix Length; Prefix
	wordsMap, writingQueue := input.Initialize(lenFlag, prefixFlag) // Map of combinations; Map key queue
	output.GenerateText(writingQueue, lenFlag, numWordsFlag, wordsMap)
}
