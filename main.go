package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	format := `
number of args: %d
args[0] (exe): %s
args[1] (abs-file): %s
args[2] (start-line): %s
args[3] (start-col): %s
args[4] (end-line): %s
args[5] (end-col): %s

args: %v
`
	args := toArgs(os.Args)
	st := fmt.Sprintf(
		format,
		len(os.Args),
		os.Args[0],
		os.Args[1],
		os.Args[2],
		os.Args[3],
		os.Args[4],
		os.Args[5],
		args,
	)
	file := "/Users/lucascaballero/fillp.txt"
	ioutil.WriteFile(file, []byte(st), 0777)
}

// toArgs does a naive parse of the given command line tokens.  Naive here
// means it doesn't check before attempting to access the array and therefor
// could easily be out of bounds.
func toArgs(args []string) Args {
	return Args{
		Exe:     args[0],
		AbsFile: args[1],
		Range: &Range{
			Start: &Mark{
				Line: toInt(args[2]),
				Col:  toInt(args[3]),
			},
			End: &Mark{
				Line: toInt(args[4]),
				Col:  toInt(args[5]),
			},
		},
	}
}

// toInt converts the given string to an int, and if that fails it will
// return 0.
func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
