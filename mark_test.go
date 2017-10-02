package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Mark_Equal_True(t *testing.T) {
	a := Mark{Col: 42, Line: 51}
	b := Mark{Col: 42, Line: 51}
	assert.True(t, a.Equal(b))
}

func Test_Mark_Equal_False_Cols(t *testing.T) {
	a := Mark{Col: 42, Line: 51}
	b := Mark{Col: 41, Line: 51}
	assert.False(t, a.Equal(b))
}

func Test_Mark_Equal_False_Lines(t *testing.T) {
	a := Mark{Col: 42, Line: 51}
	b := Mark{Col: 42, Line: 52}
	assert.False(t, a.Equal(b))
}

func Test_Mark_Equal_False_Both(t *testing.T) {
	a := Mark{Col: 42, Line: 51}
	b := Mark{Col: 41, Line: 52}
	assert.False(t, a.Equal(b))
}
