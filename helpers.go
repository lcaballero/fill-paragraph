package main

import (
	"strings"
	"testing"
)

func ReplaceRawWithEscaped(bin []byte) string {
	s := string(bin)
	s = strings.Replace(s, "\t", "\\t", -1)
	s = strings.Replace(s, "\n", "\\n", -1)
	return s
}

func Dump(t *testing.T, bin []byte) {
	s := ReplaceRawWithEscaped(bin)
	t.Log(s)
}

