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

func main() {

	var length, num int
	var alphabet string
	var help bool

	flag.IntVar(&length, "l", 16, "length of the generated password")
	flag.IntVar(&num, "n", 1, "number of passwords generated")
	flag.StringVar(&alphabet, "a", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPYRSTUVWXYZ1234567890", "characters used for generating a password")
	flag.BoolVar(&help, "h", false, "display this help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	runes := bytes.Runes([]byte(alphabet))

	for i := 0; i < num; i++ {
		fmt.Println(generatePassword(runes, length))
	}
}

func generatePassword(runes []rune, l int) string {
	p := ""
	maxR := big.NewInt(int64(len(runes)))

	for i := 0; i < l; i++ {
		r, _ := rand.Int(rand.Reader, maxR)
		p += string(runes[r.Int64()])
	}
	return p
}
