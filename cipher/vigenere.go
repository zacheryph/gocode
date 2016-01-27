package main

import "io"

// Vigenere cipher
type Vigenere struct {
	Key     string
	keyLen  int
	Decrypt bool
	dst     io.Writer
}

// NewVigenere creates a new vigenere cipher
func NewVigenere(dst io.Writer, key string, decrypt bool) *Vigenere {
	return &Vigenere{
		Key:     key,
		keyLen:  len(key),
		Decrypt: decrypt,
		dst:     dst,
	}
}

func runeToShift(l rune) int {
	return int(l) - 'A'
}

func (c *Vigenere) Write(p []byte) (n int, err error) {
	res := ""

	for idx, char := range string(p) {
		shift := runeToShift(rune(c.Key[idx%c.keyLen]))

		if c.Decrypt {
			shift = 26 - shift
		}

		res += string(ceaserRune(char, shift))
	}
	return c.dst.Write([]byte(res))
}
