package main

import "io"

const (
	// Alphabet soup for decoding
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Ceaser the czar of ciphers
type Ceaser struct {
	Shift int
	dst   io.Writer
}

func ceaserRune(char rune, shift int) rune {
	if char < 'A' || char > 'Z' {
		return ' '
	}

	offset := (int(char) + shift - 'A') % 26
	return rune(Alphabet[offset])
}

// NewCeaser gives us a new ceaser cipher writer
func NewCeaser(dst io.Writer, shift int, decrypt bool) *Ceaser {
	if decrypt {
		shift = 26 - shift
	}

	return &Ceaser{
		Shift: shift,
		dst:   dst,
	}
}

func (c *Ceaser) Write(p []byte) (n int, err error) {
	res := ""
	for _, char := range string(p) {
		res += string(ceaserRune(char, c.Shift))
	}
	return c.dst.Write([]byte(res))
}
