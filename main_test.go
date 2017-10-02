package main

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToArgs(t *testing.T) {
	args := []string{"exe", "abs-file", "1", "2", "3", "4"}
	a := toArgs(args)
	Equal(t, a.Exe, args[0])
	Equal(t, a.AbsFile, args[1])
	Equal(t, a.Range.Start.Line, toInt(args[2]))
	Equal(t, a.Range.Start.Col, toInt(args[3]))
	Equal(t, a.Range.End.Line, toInt(args[4]))
	Equal(t, a.Range.End.Col, toInt(args[5]))
}
