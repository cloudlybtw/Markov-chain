package inputoutput

import "fmt"
import "os"

func PrintHelp() {
	fmt.Println("Markov Chain text generator.\n")
	fmt.Println("Usage:")
  	fmt.Println("markovchain [-w <N>] [-p <S>] [-l <N>]")
  	fmt.Println("markovchain --help\n")
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words")
	fmt.Println("  -p S    Starting prefix")
	fmt.Println("  -l N    Prefix length")
	os.Exit(0)
}