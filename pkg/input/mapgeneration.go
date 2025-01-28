package input

import (
	"fmt"
	"io"
	"markov-chain/pkg/utils"
	"os"
)

func Initialize(lenFlag int, prefixFlag string) (map[string][]string, *utils.Queue) {
	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading input.")
		os.Exit(1)
	}

	WordsMap := map[string][]string{}
	readingQueue := utils.NewQueue()
	writingQueue := utils.NewQueue()
	reader := ""
	for _, char := range text {
		if char == '\n' || char == ' ' {
			if len(reader) != 0 {
				if readingQueue.Len() != lenFlag {
					readingQueue.Push(reader)
					if prefixFlag == "" {
						writingQueue.Push(reader)
					}
					reader = ""
					continue
				}
				WordsMap[readingQueue.GetString()] = append(WordsMap[readingQueue.GetString()], reader)
				readingQueue.Pop()
				readingQueue.Push(reader)
				reader = ""
			}
			continue
		}
		reader += string(char)
	}
	if len(reader) != 0 {
		WordsMap[readingQueue.GetString()] = append(WordsMap[readingQueue.GetString()], reader) //" (end)"
		_ = readingQueue.Pop()
		readingQueue.Push(reader)
	}
	if readingQueue.Len() == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}
	WordsMap[readingQueue.GetString()] = append(WordsMap[readingQueue.GetString()], " (end)")
	reader = ""
	if writingQueue.Len() == 0 {
		for _, char := range prefixFlag {
			if char == '\n' {
				fmt.Fprintln(os.Stderr, "Error. '\n' is improper character, use space instead.")
				os.Exit(1)
			}
			if char == ' ' {
				if len(reader) != 0 {
					writingQueue.Push(reader)
					reader = ""
				}
				continue
			}
			reader += string(char)
		}
	}
	if len(reader) != 0 {
		writingQueue.Push(reader)
	}
	return WordsMap, writingQueue
}
