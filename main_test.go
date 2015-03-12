package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	// TODO: make this a table-driven test with more test cases
	alphabet := "a"
	got := generatePassword(alphabet, 10)
	got.testAlphabet(t, alphabet)
	got.testLength(t, 10)
}

func TestGeneratePasswords(t *testing.T) {
	// TODO: make this a table-driven test with more test cases
	alphabet := "a"
	got := generatePasswords(alphabet, 10, 10)

	for _, pw := range got {
		pw.testAlphabet(t, alphabet)
		pw.testLength(t, 10)
	}
}

func TestPrintPasswords(t *testing.T) {
	const want = "abc\n123\n"

	p := []password{"abc", "123"}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printPasswords(p)
	os.Stdout = old
	w.Close()

	b, _ := ioutil.ReadAll(r)

	if got := string(b); got != want {
		t.Errorf("Expected printPasswords to print %s, got: %s", want, got)
	}
}

func (p password) testAlphabet(t *testing.T, alphabet string) {

	for _, r := range bytes.Runes([]byte(p)) {
		if !strings.ContainsRune(alphabet, r) {
			t.Errorf("%c is not contained in the alphabet %s", r, alphabet)
		}
	}
}

func (p password) testLength(t *testing.T, length int) {
	if l := len(p); l != length {
		t.Errorf("Expected generated password to be %d characters long, was %d", length, l)
	}
}
