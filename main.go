package main

import (
	"flag"
	"fmt"
	"io"
	"markov-chain/pkg/inputoutput"
	"math/rand/v2"
	"os"
)

// flags: Number Of Words(-w); Prefix(-p); Prefix Length(-l);

var (
	numWordsFlag = flag.Int("w", 100, "Number of maximum words")
	lenFlag      = flag.Int("l", 2, "Starting prefix")
	prefixFlag   = flag.String("p", "", "Prefix length (default first words from input text)")
	helpFlag     = flag.Bool("help", false, "Show this screen.")
)

func main() {
	flag.Parse()
	stat, _ := os.Stdin.Stat()
	switch {
	case *helpFlag:
		inputoutput.PrintHelp()
		os.Exit(0)
	case *numWordsFlag < 0:
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be negative.")
		os.Exit(1)
	case *numWordsFlag > 10000:
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be more than 10,000.")
		os.Exit(1)
	case *lenFlag < 1:
		fmt.Fprintln(os.Stderr, "The prefix length can't be less than 1.")
		os.Exit(1)
	case *lenFlag > 5:
		fmt.Fprintln(os.Stderr, "The prefix length can't be more than 5.")
		os.Exit(1)
	case (stat.Mode() & os.ModeCharDevice) != 0:
		fmt.Fprintln(os.Stderr, "Error. Empty input.")
		os.Exit(1)
	}

	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading input.")
		os.Exit(1)
	}

	WordsMap := map[string][]string{}
	readingQueue := NewQueue()
	writingQueue := NewQueue()
	reader := ""
	for _, char := range text {
		if char == '\n' || char == ' ' {
			if len(reader) != 0 {
				if readingQueue.Len() != *lenFlag {
					readingQueue.Push(reader)
					if *prefixFlag == "" {
						writingQueue.Push(reader)
					}
					reader = ""
					continue
				}
				WordsMap[readingQueue.GetString()] = append(WordsMap[readingQueue.GetString()], reader)
				_ = readingQueue.Pop()
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
		for _, char := range *prefixFlag {
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
	switch {
	case writingQueue.Len() > *lenFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix longer than prefix length.")
		os.Exit(1)
	case writingQueue.Len() < *lenFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix shorter than prefix length.")
		os.Exit(1)
	case writingQueue.Len() > *numWordsFlag:
		fmt.Fprintln(os.Stderr, "Error: Prefix longer than words limit.")
		os.Exit(1)
	case !MapContains(WordsMap, writingQueue.GetString()):
		fmt.Fprintln(os.Stderr, "Error: Prefix is not present in text")
		os.Exit(1)
	}

	fmt.Print(writingQueue.GetString(), " ")
	for i := 0; i < *numWordsFlag-*lenFlag; i++ {
		word := WordsMap[writingQueue.GetString()][rand.IntN(len(WordsMap[writingQueue.GetString()]))]
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
