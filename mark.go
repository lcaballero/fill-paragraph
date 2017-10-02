package main

import "encoding/json"

// Mark is a position in some text file which can be identified by an offset
// in to the bytes of the text, but that isn't as useful for humans as the
// column and line describing a point in the file, for which a normal text
// editor will typically display.  Because editors provide col and line it's
// easy enough to pass these values to a command line tools for processing.
type Mark struct {
	HasValue bool
	Offset   int
	Col      int
	Line     int
}

// Equal determines if the provided mark is equivalent by comparing the Col
// and Line (only).
func (a Mark) Equal(b Mark) bool {
	return a.Col == b.Col && a.Line == b.Line
}

// NewMark creates a default mark at column 1 and line 1.
func NewMark() Mark {
	return Mark{
		Col:  1,
		Line: 1,
	}
}

// ToJson produces a JSON string for this Mark.
func (m Mark) ToJson() string {
	bin, err := json.MarshalIndent(&m, "", "  ")
	if err != nil {
		return ""
	}
	return string(bin)
}