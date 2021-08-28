package main

import (
	"os"

	"github.com/ca-std/lib"
)

func main() {
	challenge1()
	challenge2()
}

func challenge2() {
	out(lib.XorY(<-lib.DecodeHex([]byte("1c0111001f010100061a024b53535009181c")), <-lib.DecodeHex([]byte("686974207468652062756c6c277320657965"))))
}

func challenge1() {
	out(<-lib.DecodeHex([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")))
}

func out(b []byte) {
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
