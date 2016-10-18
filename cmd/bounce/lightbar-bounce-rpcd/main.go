package main

import (
	"fmt"
	"net/rpc"
	"net/url"
	"os"

	"golang.org/x/net/websocket"

	lightbarpc "github.com/remexre/go-lightbar/rpc"
)

func main() {
	var host string
	if len(os.Args) == 1 {
		host = "ws://localhost:8063/api/v1/reverse-rpc"
	} else if len(os.Args) == 2 {
		host = os.Args[1]
	} else {
		fmt.Fprintf(os.Stderr, "Usage: %s [host]\n", os.Args[0])
		os.Exit(1)
	}

	originURL := (&url.URL{
		Scheme: "http",
		Host:   host,
	}).String()
	wsURL := (&url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "/api/v1/reverse-rpc",
	}).String()

	ws, err := websocket.Dial(wsURL, "", originURL)
	must(err)

	rpcServer := rpc.NewServer()
	must(rpcServer.RegisterName("lightbar", lightbarpc.Get()))
	rpc.DefaultServer.ServeConn(ws)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
