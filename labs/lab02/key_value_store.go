package key_value_store

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func key_value_store() {

	// server calls for http service
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// client dials server first time
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// make a remote call
	// Synchronous call
	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
}

// rpc (remote procedure call) in go
// rpc.Register()
// rpc.Dial()
// rpc.Serve()
// rpc.NewClient()
// rpc.NewServer()
// rpc.NewClientWithCodec()
// rpc.NewServerCodec()
// rpc.RegisterName()
// rpc.Register()

// write key-value database as a client-server application using rpc
// client: put(key, value), get(key)
// server: store(key, value), retrieve(key)
// use map to store key-value pairs
// use rpc to communicate between client and server
// use go to implement concurrency
// run in two terminals
// go run server.go
// go run client.go
// use go build to build the program
// go build server.go
// go build client.go
// ./server
// ./client

// in server.go file write wait for client to connect
// in client.go file write connect to server and send request
// in server.go file write handle request and send response
// in client.go file write handle response
// in server.go file write store and retrieve key-value pairs
// in client.go file write put and get key-value pairs
// in server.go file write concurrency
// in client.go file write concurrency
// in server.go file write rpc
// in client.go file write rpc
// in server.go file write rpc with concurrency
// in client.go file write rpc with concurrency
// in server.go file write rpc with concurrency and lock
// in client.go file write rpc with concurrency and lock
// in server.go file write rpc with concurrency and lock and wait
// in client.go file write rpc with concurrency and lock and wait
