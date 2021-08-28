package main

import (
	"bufio"
	"bytes"
	"os"

	"github.com/ca-std/lib"
)

func main() {
	challenge1()
	challenge2()
	challenge3()
	challenge4()
	challenge5()
}

func expand(k []byte, toSize int) []byte {
	result := make([]byte, toSize)
	n := len(k)
	for i := 0; i < toSize; i++ {
		result[i] = k[i%n]
	}
	return result
}

func challenge5() {
	input := `Burning 'em, if you ain't quick and nimble
	I go crazy when I hear a cymbal`
	key := "ICE"
	expanded := expand([]byte(key), len(input))

	out(<-lib.EncodeHex(lib.XorY(expanded, []byte(input))))
}

func challenge4() {
	raw, err := os.ReadFile("4.txt")
	if err != nil {
		panic(err)
	}
	_, inverse, top := lib.Counter(raw)

	reader := bytes.NewReader(raw)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		b := scanner.Bytes()
		out(<-lib.EncodeHex(lib.XorK(b, inverse[top])))
	}
}

func challenge3() {
	b := <-lib.DecodeHex([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	_, inverse, top := lib.Counter(b)
	out(lib.XorK(b, inverse[top]))
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
