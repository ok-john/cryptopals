package main

import (
	"bufio"
	"bytes"
	"math/big"
	"os"

	"github.com/ca-std/lib"
)

func main() {
	challenge1()
	challenge2()
	challenge3()
	challenge4()
	challenge5()
	challenge6()
}

func challenge6() {
	turkey := ham([]byte("this is a test"), []byte("wokka wokka!!!"))
	if turkey != 37 {
		panic("not chill dude")
	}

}

func challenge5() {
	input := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	out(<-lib.EncodeHex(lib.XorY(expand([]byte("ICE"), len(input)), input)))
}

func challenge4() {
	raw := read("4.txt")
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

func read(path string) []byte {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return raw
}

func ham(t0, t1 []byte) int {
	x0, x1 := new(big.Int), new(big.Int)
	if len(t0) != len(t1) {
		panic("dont u know that hams must be equal length?")
	}
	x0.SetBytes(t0)
	x1.SetBytes(t1)
	hamming := 0

	s0, s1 := x0.Text(2), x1.Text(2)

	for i := 0; i < len(s0); i++ {
		if s0[i] != s1[i] {
			hamming += 1
		}
	}
	return hamming
}

func contract(k []byte, toSize int) ([][]byte, int) {
	result, n := [][]byte{}, len(k)
	leftover := n % toSize
	for i := 0; i < n-leftover; i += toSize {
		j := i + toSize
		result = append(result, k[i:j])
	}

	if leftover > 0 {
		result = append(result, k[n-leftover:])
	}

	return result, leftover
}

func expand(k []byte, toSize int) []byte {
	result := make([]byte, toSize)
	n := len(k)
	for i := 0; i < toSize; i++ {
		result[i] = k[i%n]
	}
	return result
}
