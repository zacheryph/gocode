package main

import "io"

const (
	// Alphabet soup for decoding
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Caeser the czar of ciphers
type Caeser struct {
	Shift int
	dst   io.Writer
}

func caeserRune(char rune, shift int) rune {
	if char < 'A' || char > 'Z' {
		return ' '
	}

	offset := (int(char) + shift - 'A') % 26
	return rune(Alphabet[offset])
}

// NewCaeser gives us a new caeser cipher writer
func NewCaeser(dst io.Writer, shift int, decrypt bool) *Caeser {
	if decrypt {
		shift = 26 - shift
	}

	return &Caeser{
		Shift: shift,
		dst:   dst,
	}
}

func (c *Caeser) Write(p []byte) (n int, err error) {
	res := ""
	for _, char := range string(p) {
		res += string(caeserRune(char, c.Shift))
	}
	return c.dst.Write([]byte(res))
}
