package input

import (
	"flag"
	"fmt"
	"markov-chain/pkg/output"
	"os"
)

var (
	numWordsFlag = flag.Int("w", 100, "Number of maximum words")
	lenFlag      = flag.Int("l", 2, "Starting prefix")
	prefixFlag   = flag.String("p", "", "Prefix length (default first words from input text)")
	helpFlag     = flag.Bool("help", false, "Show this screen.")
)

func ParseFlags() (int, int, string) {
	flag.Parse()
	stat, _ := os.Stdin.Stat()
	switch {
	case *helpFlag:
		output.PrintHelp()
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
	return *numWordsFlag, *lenFlag, *prefixFlag
}
