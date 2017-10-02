package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func Test_Strip_Prefix(t *testing.T) {
	a := &Args{}

	text := `
	Line 1, with a 1 of tokens, k of tokens
	Line 2, a of tokens, b of tokens, some tokens
	Line 3, may small tokens a character wide.
`

	// remove the lame single newline
	rawLines := a.toLines([]byte(text[1:]))
	lines, err := a.stripPrefix([]byte("	Line "), rawLines)

	assert.Nil(t, err)
	assert.Equal(t, len(lines), len(rawLines))

	Dump(t, lines[0])
	Dump(t, lines[1])
	Dump(t, lines[2])

	assert.Equal(t, "1, with a 1 of tokens, k of tokens", string(lines[0]))
	assert.Equal(t, "2, a of tokens, b of tokens, some tokens", string(lines[1]))
	assert.Equal(t, "3, may small tokens a character wide.", string(lines[2]))
}

func Test_To_Lines(t *testing.T) {
	a := &Args{}

	text := `	Line 1
	Line 2
	Line 3
`
	lines := a.toLines([]byte(text))
	t.Log(lines)
	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "\tLine 1", string(lines[0]))
	assert.Equal(t, "\tLine 2", string(lines[1]))
	assert.Equal(t, "\tLine 3", string(lines[2]))
}

func Test_To_Selection(t *testing.T) {
	a := &Args{
		Range: &Range{
			Start: &Mark{
				Line: 3,
				Col:  1,
			},
			End: &Mark{
				Line: 11,
				Col:  69,
			},
		},
	}

	actual, err := a.toSelection([]byte(Index))
	assert.Nil(t, err)

	expected := IndexLoremIpsum[1:] // skipping first newline

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, string(actual))
}

func Test_Selection_In_Line(t *testing.T) {
	line := `Here is a line of text to select from.`
	expected := "is a line of text"

	start := strings.Index(line, expected)
	base := 1

	a := &Args{
		Range: &Range{
			Start: &Mark{
				Line: 1,
				Col:  start + base,  // 6
			},
			End: &Mark{
				Line: 1,
				Col:  start + base + len(expected), // 23
			},
		},
	}

	actual, err := a.toSelection([]byte(line))
	assert.Nil(t, err)

	assert.Equal(t, 5 + base, a.Range.Start.Col)
	assert.Equal(t, start + base + len(expected), a.Range.End.Col)
	assert.Equal(t, expected, line[start:start + len(expected)])

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, string(actual))
}

func Test_Selection_With_Zero_Length(t *testing.T) {

}