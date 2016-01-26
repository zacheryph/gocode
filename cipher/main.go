package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	// Alphabet soup for decoding
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXY"
)

var (
	shift = flag.Int("shift", 0, "shift parameter for Ceaser Cipher")
	key   = flag.String("key", "", "keyword for Vigenere Cipher")
)

func abort(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *shift == 0 && *key == "" {
		abort("Must give -shift (ceaser) or -key (vigenere) flag")
	}

	if *shift < 0 || *shift > 25 {
		abort("-shift option must range between 1 and 25")
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	var res string
	if *shift > 0 {
		res = ceaserCipher(input)
	} else {
		res = vigeneteCipher(input)
	}

	fmt.Println("Response:")
	fmt.Println(res)
}

func ceaserRune(char rune, shift int) rune {
	cipher := Alphabet[shift:]
	if char < 'A' || char > 'Z' {
		return ' '
	}

	return rune(cipher[int(char)-'A'])
}

func ceaserCipher(input string) string {
	result := ""

	for _, char := range input {
		result += string(ceaserRune(char, *shift))
	}

	return result
}

func vigeneteCipher(r string) string {
	return "Vigenete: Not Implemented"
}
