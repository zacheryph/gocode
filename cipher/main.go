package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	shift   = flag.Int("shift", 0, "shift parameter for Caeser Cipher")
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
		abort("Must give -shift (caeser) or -key (vigenere) flag")
	}

	if *shift < 0 || *shift > 25 {
		abort("-shift option must range between 1 and 25")
	}

	fmt.Println("Response:")
	var cipher io.Writer
	if *shift > 0 {
		cipher = NewCaeser(os.Stdout, *shift, *decrypt)
	} else {
		cipher = NewVigenere(os.Stdout, *key, *decrypt)
	}

	io.Copy(cipher, os.Stdin)
	fmt.Println("")
}
