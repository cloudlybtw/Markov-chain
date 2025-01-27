package main

import (
	// "fmt"
	// "os"
	"flag"
	"fmt"
	"os"
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

var (
	numWordsFlag  = flag.Int("w", 100, "help him")
	prefixLenFlag = flag.Int("l", 2, "nah dont help him")
)

func main() {
	// var numWordsFlag = flag.Int("w", 100, "help him")
	// // var prefixFlag = flag.String("")
	// var prefixLenFlag = flag.Int("l", 2, "nah dont help him")
	flag.Parse()
	if *numWordsFlag < 0 {
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be negative.")
		os.Exit(1)
	} else if *numWordsFlag > 10000 {
		fmt.Fprintln(os.Stderr, "The maximum number of words can't be more than 10,000.")
		os.Exit(1)
	}
	if *prefixLenFlag < 1 {
		fmt.Fprintln(os.Stderr, "The prefix length can't be less than 1.")
		os.Exit(1)
	} else if *prefixLenFlag > 5 {
		fmt.Fprintln(os.Stderr, "The prefix length can't be more than 5.")
		os.Exit(1)
	}
	fmt.Println(*numWordsFlag)
	fmt.Println(*prefixLenFlag)

	// text, err := os.ReadFile("temptext.txt")
	// if err != nil {
	// 	os.Exit(1)
	// }
	// words := []string{}
	// reader := ""
	// for _, char := range text {
	// 	if char == '\n' || char == ' ' {
	// 		if len(reader) != 0 {
	// 			words = append(words, reader)
	// 			reader = ""
	// 		}
	// 		continue
	// 	}
	// 	reader += string(char)
	// }
	// if len(reader) != 0 {
	// 	words = append(words, reader)
	// }

	// WordsMap := map[string][]string{}

	// queue := NewQueue()
	// for i := 0; i < 5; i++ { //len(words)-2
	// 	queue.Push(words[i])
	// }
	// fmt.Println(queue.GetString())
	// fmt.Println(WordsMap)
}
