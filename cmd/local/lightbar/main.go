package main

import (
	"fmt"
	"os"
	"strconv"

	lightbar "github.com/remexre/go-lightbar"
)

func main() {
	args := make([]byte, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		n, err := strconv.ParseUint(arg, 16, 8)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			usage()
		}
		args[i] = byte(n)
	}
	lb := lightbar.Get()

	if len(args) == 1 {
		must(lb.SetBrightness(args[0]))
	} else if len(args) == 4 {
		must(lb.SetLED(args[0], args[1], args[2], args[3]))
	} else if len(args) == 12 {
		must(lb.SetLEDs([4][3]byte{
			[3]byte{args[0], args[1], args[2]},
			[3]byte{args[3], args[4], args[5]},
			[3]byte{args[6], args[7], args[8]},
			[3]byte{args[9], args[10], args[11]},
		}))
	} else {
		usage()
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	for _, s := range []string{
		"brightness",
		"led r g b",
		"r1 g1 b1 r2 g2 b2 r3 g3 b3 r4 g4 b4",
	} {
		fmt.Fprintf(os.Stderr, "\t%s %s\n", os.Args[0], s)
	}
	os.Exit(1)
}
