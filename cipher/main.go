package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	// Alphabet soup for decoding
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	shift   = flag.Int("shift", 0, "shift parameter for Ceaser Cipher")
	key     = flag.String("key", "", "keyword for Vigenere Cipher")
	decrypt = flag.Bool("decrypt", false, "Decrypt an encrypted string. Default is Encrypt")
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
	input = strings.ToUpper(input)

	var res string
	if *shift > 0 {
		if *decrypt {
			*shift = 26 - *shift
		}

		res = ceaserCipher(input)
	} else {
		res = vigenereCipher(input)
	}

	fmt.Println("Response:")
	fmt.Println(res)
}

func runeToShift(l rune) int {
	return int(l) - 'A'
}

func ceaserRune(char rune, shift int) rune {
	if char < 'A' || char > 'Z' {
		return ' '
	}

	offset := (int(char) + shift - 'A') % 26
	return rune(Alphabet[offset])
}

func ceaserCipher(input string) string {
	result := ""

	for _, char := range input {
		result += string(ceaserRune(char, *shift))
	}

	return result
}

func vigenereCipher(input string) string {
	result := ""
	keyLen := len(*key)

	for idx, char := range input {
		shift := runeToShift(rune((*key)[idx%keyLen]))

		if *decrypt {
			shift = 26 - shift
		}

		result += string(ceaserRune(char, shift))
	}

	return result
}
