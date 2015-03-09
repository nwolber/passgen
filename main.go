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
	defaultCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPYRSTUVWXYZ1234567890"
)

func main() {
	var (
		length, num int
		character string
		help bool
	)

	flag.IntVar(&length, "l", defaultLength, "length of the generated password")
	flag.IntVar(&num, "n", defaultCount, "number of passwords generated")
	flag.StringVar(&character, "a", defaultCharacters, "characters used for generating a password")
	flag.BoolVar(&help, "h", false, "display this help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	runes := bytes.Runes([]byte(character))

	for i := 0; i < num; i++ {
		fmt.Println(generatePassword(runes, length))
	}
}

func generatePassword(runes []rune, l int) string {
	p := ""
	numRunes := big.NewInt(int64(len(runes)))

	for i := 0; i < l; i++ {
		r, _ := rand.Int(rand.Reader, numRunes)
		p += string(runes[r.Int64()])
	}
	return p
}
