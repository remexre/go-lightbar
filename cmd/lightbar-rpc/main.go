package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	var addr, net string
	var base int
	flag.StringVar(&addr, "addr", "localhost:15917", "the address to connect to")
	flag.IntVar(&base, "b", 16, "the base to use for input")
	flag.StringVar(&net, "net", "tcp", "the network protocol to use")
	flag.Parse()

	args := make([]byte, flag.NArg())
	for i, arg := range flag.Args() {
		n, err := strconv.ParseUint(arg, base, 8)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			usage()
		}
		args[i] = byte(n)
	}

	client, err := rpc.DialHTTP("tcp", addr)
	must(err)

	if len(args) == 1 {
		must(client.Call("lightbar.SetBrightness", args[0], nil))
	} else if len(args) == 4 {
		must(client.Call("lightbar.SetLED", [4]byte{
			args[0], args[1], args[2], args[3],
		}, nil))
	} else if len(args) == 12 {
		must(client.Call("lightbar.SetLEDs", [4][3]byte{
			[3]byte{args[0], args[1], args[2]},
			[3]byte{args[3], args[4], args[5]},
			[3]byte{args[6], args[7], args[8]},
			[3]byte{args[9], args[10], args[11]},
		}, nil))
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
		fmt.Fprintf(os.Stderr, "\t%s [FLAGS...] %s\n", os.Args[0], s)
	}
	flag.Usage()
	os.Exit(1)
}
