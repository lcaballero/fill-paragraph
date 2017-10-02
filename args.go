package main

import (
	"encoding/json"
	"errors"
)

// Args is intended to hold the name of the Exe, it's full absolute path, and
// the Range as provided at the command line, typically via number of flags.
type Args struct {
	Exe     string
	AbsFile string
	Range   *Range
}

// ToJson outputs this Args instance as a JSON string.
func (a *Args) ToJson() string {
	bin, err := json.MarshalIndent(&a, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}

// toSelection finds withing the text bytes the slice which matches the Args
// start and end Range.  If the Range cannot be found in the bytes then an
// empty slice and an error is returned.
func (a *Args) toSelection(text []byte) ([]byte, error) {
	p, start, end := NewMark(), NewMark(), NewMark()
	n := len(text)

	for ; p.Offset < n; p.Offset++ {
		if a.isStart(p) {
			start = p
			start.HasValue = true
		}

		if a.isEnd(p) {
			end = p
			end.HasValue = true
		}

		if start.HasValue && end.HasValue {
			return text[start.Offset:end.Offset], nil
		}

		if rune(text[p.Offset]) == '\n' {
			p.Line++
			p.Col = 1
		} else {
			p.Col++
		}
	}
	return []byte{}, errors.New("couldn't find selection")
}

// stripPrefix should remove the prifix bytes from the lines
func (a *Args) stripPrefix(prefix []byte, lines [][]byte) ([][]byte, error) {
	lineCount := len(lines)
	prefixSize := len(prefix)
	res := [][]byte{}

	for i := 0; i < lineCount; i++ {
		line := lines[i]
		lineSize := len(line)

		for j := 0; j < prefixSize && j < lineSize; j++ {
			a := line[j]
			b := prefix[j]

			if a != b {
				return res, errors.New("prefix inconsistent")
			}
		}
		res = append(res, line[prefixSize:])
	}

	return res, nil
}

// toLines produces lines from the given bytes of text.
func (a *Args) toLines(text []byte) [][]byte {
	n := len(text)
	lines := [][]byte{}
	m, i := 0, 0

	for ; i < n; i++ {
		c := text[i]

		if c == '\n' {
			lines = append(lines, text[m:i])
			i++ // Remove '\n' from the line
			m = i
		}
	}

	if m < n - 1 {
		lines = append(lines, text[m:n])
	}

	return lines
}

// isStart tests if the given Mark is equal to the start provided by Args
func (a Args) isStart(p Mark) bool {
	return a.Range.Start.Equal(p)
}

// isEnd tests if the given Mark is equal to the send provided by Args
func (a Args) isEnd(p Mark) bool {
	return a.Range.End.Equal(p)
}
