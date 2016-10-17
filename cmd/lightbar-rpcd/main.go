package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"os"

	lightbarpc "github.com/remexre/go-lightbar/rpc"
)

func main() {
	var addr string
	if len(os.Args) == 1 {
		addr = ":15917"
	} else if len(os.Args) == 2 {
		addr = os.Args[1]
	} else {
		fmt.Fprintf(os.Stderr, "Usage: %s [addr]\n", os.Args[0])
		os.Exit(1)
	}

	must(rpc.RegisterName("lightbar", lightbarpc.Get()))
	rpc.HandleHTTP()
	http.ListenAndServe(addr, nil)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
