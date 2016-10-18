package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"sync"

	"golang.org/x/net/websocket"

	"github.com/gin-gonic/gin"
)

// DefaultAddr is the default address served on.
const DefaultAddr = ":8063"

var logCh = make(chan *rpc.Call, 16)

var rpcConnections map[*rpc.Client]struct{}
var rpcConnectionsLock sync.Mutex

func main() {
	addr := flag.String("addr", DefaultAddr, "the address to serve on")
	var debug bool
	flag.BoolVar(&debug, "debug", false, "enable debugging")
	flag.Parse()

	if debug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	go doLogs(logCh)

	r := gin.Default()
	r.Group("/api/v1").
		GET("/reverse-rpc", func(c *gin.Context) {
			websocket.Handler(func(c *websocket.Conn) {
				rpcConnectionsLock.Lock()
				rpcConnections[rpc.NewClient(c)] = struct{}{}
				rpcConnectionsLock.Unlock()
			}).ServeHTTP(c.Writer, c.Request)
		}).
		POST("/led/all", RGBSetter(4))

	if addr == nil {
		r.Run()
	} else {
		r.Run(*addr)
	}
}

func doLogs(ch chan *rpc.Call) {
	for call := range ch {
		log.Println(call)
	}
}

// RGBSetter returns a handler that sets a RGB values for an LED.
func RGBSetter(n byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		if n > 4 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		r, err := PostFormToByte(c, "r")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		g, err := PostFormToByte(c, "g")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		b, err := PostFormToByte(c, "b")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		rpcConnectionsLock.Lock()
		for connection := range rpcConnections {
			connection.Go("lightbar.SetLED", [4]byte{n, r, g, b}, nil, logCh)
		}
		rpcConnectionsLock.Unlock()
	}
}

// PostFormToByte extracts a named byte from an HTTP POST form.
func PostFormToByte(c *gin.Context, key string) (byte, error) {
	val, b := c.GetPostForm(key)
	if !b {
		return 0, errors.New("no such key: " + key)
	}
	n, err := strconv.ParseUint(val, 10, 8)
	return byte(n), err
}
