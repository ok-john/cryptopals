package main

import (
	"bufio"
	"bytes"
	"log"
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
	turkey := lib.Sb([]byte("wokka wokka!!!")).Ham(lib.Sb([]byte("this is a test")))
	if !turkey.Eq(lib.Si(37)) {
		panic("not chill dude")
	}

	keySpace := lib.PointDistribution(1)
	ham := q6().Apply(func(x, y *lib.Set) *lib.Set {
		q6().Apply(func(o, l *lib.Set) *lib.Set {
			o = o.Ham(l)
			return o
		}, x)
		return x
	}, lib.Si(1))

	for keysize, reader := int64(2), rd("6.txt"); keysize < reader.Size(); keysize++ {
		for i := int64(0); i < reader.Size(); i += keysize {
			buff := make([]byte, keysize)
			reader.ReadAt(buff, i)
			keySpace = keySpace.Insert(lib.Sb(buff))
		}
	}

	ham = ham.Intersection(keySpace)
	viewSpace("total key-space", keySpace)
	viewSpace("intersection with hamming-space", ham)
}

func q6() *lib.Universe {
	u := lib.UniformDistribution(1)
	for scanner := read("6.txt"); scanner.Scan(); {
		u = u.Insert(lib.Sb(<-lib.DecodeBase64(scanner.Bytes())))
	}
	return u
}

func challenge5() {
	input := []byte("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	out(<-lib.EncodeHex(lib.XorY(lib.Expand([]byte("ICE"), len(input)), input)))
}

func challenge4() {
	u8, max, set := lib.UniformDistribution(8), 0, []byte{}
	for scanner := read("4.txt"); scanner.Scan(); {
		s := heaviest(<-lib.DecodeHex(scanner.Bytes()), u8)
		if w := weigh(s); w > max {
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
	u.Apply(func(s1, s2 *lib.Set) *lib.Set {
		for _, b := range s1.V.Bytes() {
			x := lib.XorK(s2.V.Bytes(), b)
			if w := weigh(x); w > max {
				max, r = w, x
			}
		}
		return s1
	}, lib.Sb(from))
	return r
}

func weigh(raw []byte) int {
	freq, w := map[string]float64{}, 0.0
	freq["e"], freq["t"], freq["a"], freq["o"], freq["i"], freq["n"], freq["s"], freq["r"], freq["h"], freq["d"], freq["l"], freq["u"], freq["c"], freq["m"], freq["f"], freq["y"], freq["w"], freq["g"], freq["p"], freq["b"], freq["v"], freq["k"], freq["z"], freq["q"], freq["j"], freq["z"] = 12.02, 9.10, 8.12, 7.68, 7.31, 6.95, 6.28, 6.02, 5.92, 4.32, 3.98, 2.88, 2.71, 2.61, 2.30, 2.11, 2.09, 2.03, 1.82, 1.49, 1.11, 0.69, 0.17, 0.11, 0.10, 0.07
	freq = map[string]float64{"a": 8.12, "b": 1.49, "c": 2.71, "d": 4.32, "e": 12.02, "f": 2.3, "g": 2.03, "h": 5.92, "i": 7.31, "j": 0.1, "k": 0.69, "l": 3.98, "m": 2.61, "n": 6.95, "o": 7.68, "p": 1.82, "q": 0.11, "r": 6.02, "s": 6.28, "t": 9.1, "u": 2.88, "v": 1.11, "w": 2.09, "y": 2.11, "z": 0.07}
	for _, b := range raw {
		if v, exists := freq[string(b)]; exists {
			w += v
			continue
		}
		w -= float64(len(raw)) * 0.01
	}
	return int(w * 10.0000000005)
}

func viewSpace(name string, u *lib.Universe) {
	msb0, msb1 := u.Slice(lib.F_MSB0)
	lsb0, lsb1 := u.Slice(lib.F_LSB0)
	log.Printf("- - %s", name)
	log.Printf("size |   -  %d  -  |", u.Size())
	log.Printf("msb0 |  %d x %d  | msb1", msb0.Size(), msb1.Size())
	log.Printf("lsb0 |  %d x %d  | lsb1", lsb0.Size(), lsb1.Size())
}
