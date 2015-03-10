package main

import (
	"bytes"
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
