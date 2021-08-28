package main

import (
	"bufio"
	"bytes"
	"math"
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
	turkey := lib.Ham([]byte("this is a test"), []byte("wokka wokka!!!"))
	if turkey != 37 {
		panic("not chill dude")
	}

	shortestKeySize, shortestDistance := 0, math.MaxInt
	distancesToFrequency := map[int]int{}
	distancesToKeySizes := map[int][]int{}
	for keysize, reader, distance := int64(2), rd("6.txt"), 0; keysize < reader.Size()/2; keysize++ {
		for i := int64(0); i < reader.Size(); i += keysize {
			p0, p1 := make([]byte, keysize), make([]byte, keysize)
			reader.ReadAt(p0, i)
			reader.ReadAt(p1, i+keysize)
			distance += lib.Ham(lib.Expand(lib.ToBase(p0, 2), int(keysize)), lib.Expand(lib.ToBase(p1, 2), int(keysize)))
		}
		distance = distance / int(keysize)
		distancesToFrequency[distance]++
		distancesToKeySizes[distance] = append(distancesToKeySizes[distance], int(keysize))
		shortestDistance, shortestKeySize = lib.MinEq(shortestDistance, distance), lib.MinEq(int(shortestKeySize), int(keysize))
	}

}

func challenge5() {
	input := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	out(<-lib.EncodeHex(lib.XorY(lib.Expand([]byte("ICE"), len(input)), input)))
}

func challenge4() {
	u8, max, set := lib.UniformDistribution(8), 0, []byte{}
	for scanner := read("4.txt"); scanner.Scan(); {
		s := heaviest(<-lib.DecodeHex(scanner.Bytes()), u8)
		w := weigh(s)
		if w > max {
			max, set = w, s
		}
	}
	out(set)
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

func rd(path string) *bytes.Reader {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(raw)
}

func read(path string) *bufio.Scanner {
	return bufio.NewScanner(rd(path))
}


func heaviest(from []byte, u *lib.Universe) []byte {
	max, r := 0, []byte{}
	_, sets := u.PCollect()
	for _, s := range sets {
		if s != nil {
			for _, b := range s.V.Bytes() {
				x := lib.XorK(from, b)
				if w := weigh(x); w > max {
					max, r = w, x
				}
			}
		}
	}
	return r
}

func weigh(raw []byte) int {
	freq, w := map[string]float64{}, 0.0
	freq["e"], freq["t"], freq["a"], freq["o"], freq["i"], freq["n"], freq["s"], freq["r"], freq["h"], freq["d"], freq["l"], freq["u"], freq["c"], freq["m"], freq["f"], freq["y"], freq["w"], freq["g"], freq["p"], freq["b"], freq["v"], freq["k"], freq["z"], freq["q"], freq["j"], freq["z"] = 12.02, 9.10, 8.12, 7.68, 7.31, 6.95, 6.28, 6.02, 5.92, 4.32, 3.98, 2.88, 2.71, 2.61, 2.30, 2.11, 2.09, 2.03, 1.82, 1.49, 1.11, 0.69, 0.17, 0.11, 0.10, 0.07
	for _, b := range raw {
		if v, exists := freq[string(b)]; exists {
			w += v
		}
	}
	return int(w * 10.0000000005)
}

