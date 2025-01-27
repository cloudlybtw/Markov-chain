package main

import (
	"flag"
	"fmt"
	"os"
	"io"
)

// flags: Number Of Words(-w); Prefix(-p); Prefix Length(-l);

type Queue struct {
	queue []string
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]string, 0),
	}
}

func (q *Queue) Push(v string) {
	q.queue = append(q.queue, v)
}

func (q *Queue) Pop() string {
	if len(q.queue) == 0 {
		return "Empty queue"
	}
	element := q.queue[0]
	q.queue = q.queue[1:]
	return element
}

func (q *Queue) GetString() string {
	str := ""
	for i, a := range q.queue {
		str += a
		if i+1 != len(q.queue) {
			str += " "
		}
	}
	return str
}

func (q *Queue) Len() int {
	return len(q.queue)
}


var (
	numWordsFlag  = flag.Int("w", 100, "help him")
	lenFlag = flag.Int("l", 2, "nah dont help him")
	prefixFlag = flag.String("p", "", "never help anyone")
)

func main() {
	flag.Parse()
	if *numWordsFlag < 0 {
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be negative.")
		os.Exit(1)
	} else if *numWordsFlag > 10000 {
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be more than 10,000.")
		os.Exit(1)
	}
	if *lenFlag < 1 {
		fmt.Fprintln(os.Stderr, "The prefix length can't be less than 1.")
		os.Exit(1)
	} else if *lenFlag > 5 {
		fmt.Fprintln(os.Stderr, "The prefix length can't be more than 5.")
		os.Exit(1)
	}
	fmt.Println(*numWordsFlag)
	fmt.Println(*lenFlag)
	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while reading input.")
		os.Exit(1)
	}

	// text, err := os.ReadFile("temptext.txt")
	// if err != nil {
	// 	os.Exit(1)
	// }
	WordsMap := map[string][]string{}
	readingQueue := NewQueue()
	reader := ""
	for _, char := range text {
		if char == '\n' || char == ' ' {
			if len(reader) != 0 {
				if readingQueue.Len() != *lenFlag {
					readingQueue.Push(reader)
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
	WordsMap[readingQueue.GetString()] = append(WordsMap[readingQueue.GetString()], " (end)")
	
	fmt.Println(WordsMap)
}
