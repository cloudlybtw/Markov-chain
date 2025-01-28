package output

import (
	"fmt"
	"markov-chain/pkg/utils"
	"math/rand/v2"
	"os"
)

func GenerateText(writingQueue *utils.Queue, lenFlag int, numWordsFlag int, wordsMap map[string][]string) {
	switch {
	case writingQueue.Len() > lenFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix longer than prefix length.")
		os.Exit(1)
	case writingQueue.Len() < lenFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix shorter than prefix length.")
		os.Exit(1)
	case writingQueue.Len() > numWordsFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix longer than words limit.")
		os.Exit(1)
	case !utils.MapContains(wordsMap, writingQueue.GetString()):
		fmt.Fprintln(os.Stderr, "Error: Prefix is not present in text")
		os.Exit(1)
	}

	fmt.Print(writingQueue.GetString(), " ")
	for i := 0; i < numWordsFlag-lenFlag; i++ {
		word := wordsMap[writingQueue.GetString()][rand.IntN(len(wordsMap[writingQueue.GetString()]))]
		if word == " (end)" {
			fmt.Print("\n")
			os.Exit(0)
		}
		fmt.Printf("%s ", word)
		_ = writingQueue.Pop()
		writingQueue.Push(word)
	}
	fmt.Print("\n")
}
