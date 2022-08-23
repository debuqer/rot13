package main

import (
	"bytes"
	"testing"
)

func TestAllCharacterEncrypting(t *testing.T) {
	actual := rot13([]byte("abcdefghijklmnopqrstuvxyz"))
	expected := []byte("nopqrstuvwxyzabcdefghiklm")

	if !bytes.Equal(actual, expected) {
		t.Errorf("actual is %v but expected %v", actual, expected)
	}
}

func TestAllCharacterDecrypting(t *testing.T) {
	actual := rot13([]byte("nopqrstuvwxyzabcdefghiklm"))
	expected := []byte("abcdefghijklmnopqrstuvxyz")

	if !bytes.Equal(actual, expected) {
		t.Errorf("actual is %s but expected %s", actual, expected)
	}
}

func TestSpaceCanHandle(t *testing.T) {
	actual := rot13([]byte("hello john"))
	expected := []byte("uryyb wbua")

	if !bytes.Equal(actual, expected) {
		t.Errorf("actual is %s but expected %s", actual, expected)
	}
}
