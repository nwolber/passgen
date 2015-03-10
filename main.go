// Copyright (c) 2015 Niklas Wolber
// This file is licensed under the MIT license.
// See the LICENSE file for more information.

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	defaultLength   = 16
	defaultCount    = 1
	defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPYRSTUVWXYZ1234567890"
)

type password string

func main() {
	var (
		length, num int
		alphabet    string
		help        bool
	)

	flag.IntVar(&length, "l", defaultLength, "length of the generated password")
	flag.IntVar(&num, "n", defaultCount, "number of passwords generated")
	flag.StringVar(&alphabet, "a", defaultAlphabet, "characters used for generating a password (may contain any valid unicode character)")
	flag.BoolVar(&help, "h", false, "display this help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	for pw := range generatePasswords(alphabet, length, num) {
		fmt.Println(pw)
	}
}

func generatePasswords(alphabet string, l, n int) []password {
	pw := make([]password, n)
	for i := 0; i < n; i++ {
		pw[i] = generatePassword(alphabet, l)
	}
	return pw
}

func generatePassword(alphabet string, l int) password {
	p := ""
	runes := getRunes(alphabet)
	numRunes := big.NewInt(int64(len(runes)))

	for i := 0; i < l; i++ {
		r, _ := rand.Int(rand.Reader, numRunes)
		p += string(runes[r.Int64()])
	}
	return password(p)
}

func getRunes(s string) []rune {
	return bytes.Runes([]byte(s))
}
