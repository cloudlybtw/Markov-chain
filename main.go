package main

import "fmt"
import "os"
// flags: Number Of Words(-w); Prefix(-p); Prefix Length(-l);


func main() {
	text, err := os.ReadFile("temptext.txt")
	if err != nil {
		os.Exit(1)
	}
	words := []string {}
	reader := ""
	for _, char := range text {
		if char == '\n' || char == ' ' {
			if len(reader) != 0 {
				words = append(words, reader)
				reader = ""
			}
			continue
		}
		reader += string(char)
	}
	if len(reader) != 0 {
		words = append(words, reader)
	}
	
	WordsMap := map[string][]string {}

	for i := 0; i < len(words)-2; i++ {
		WordsMap[words[i]+" "+words[i+1]] = append(WordsMap[words[i]+" "+words[i+1]], words[i+2])
	}

	fmt.Println(WordsMap)
}
