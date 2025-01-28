# markov-chain

Text generator using Markov Chain algorithm

### Markov Chain Algorithm
An intelligent way to handle this type of text generation is with a technique called a Markov chain algorithm. This method views the input as a series of overlapping sequences, breaking each sequence into two parts: a multi-word prefix and a single suffix word that follows the prefix. The algorithm creates new sequences by randomly picking a suffix that follows each prefix, based on patterns in the original text.

### Usage:
cat input.text | ./markovchain [-w] [-l] [-p]

### Number of words (-w Int)
Accepts maximum number of words to be generated. Program prints generated text according to the Markov Chain algorithm limited by the given maximum number of words.

Default value: 100

#Constraints:

    Given number can't be negative.
    Given number can't be more 10,000.

#Example:

cat input.text | ./markovchain -w 7414

### Prefix length (-l Int)
Accepts the prefix length. Program prints generated text according to the Markov Chain algorithm with the given prefix length.

Default value: 2

#Constraints:

    Given prefix length can't be negative or 0.
    Given prefix length can't be greater than 5.

#Example:

cat input.text | ./markovchain -l 4

### Prefix (-p String)
Accepts the starting prefix. Program prints generated text according to the Markov Chain algorithm that starts with the given prefix.

By default starting prefix matches with provided text.

#Constraints:

    Given prefix must be present in the original text.
    Length of prefix should be the same as the prefix length value

#Example:

cat input.text | ./markovchain -p "Hello World!"