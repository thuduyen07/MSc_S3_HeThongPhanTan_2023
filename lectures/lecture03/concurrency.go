package main

import (
	"fmt"

	// install rpc
	// go get -u github.com/golang/net/rpc
	"rpc"
	"errors"
)

func main() {

	// // concurrency in go
	// for i := 0; i < 100; i++ {
	// 	ping(i)
	// }

	// // use lock in go
	// for i := 0; i < 100; i++ {
	// 	go ping(i)
	// }

	// // use lock and wait in go
	// for i := 0; i < 100; i++ {
	// 	go ping(i)
	// }
	// // sleep for 1 second
	// time.Sleep(1 * time.Second)

	// // concurrency and lock in go
	// countOdd := 0
	// countEven := 0
	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		countOdd += 1
	// 		go ping(countOdd)

	// 	} else {
	// 		countEven += 1
	// 		go ping(countEven)
	// 	}
	// }

	// //create a channel to communicate between goroutines
	// c := make(chan int)
	// // // write 1 to the channel c
	// // // if we don't have a goroutine here, the program will be blocked (go)
	// // go func() { c <- 1 }()

	// //write 1 to the channel c using pingWithChannel
	// go pingWithChannel(1, c)

	// // transfer value in channel
	// // if i is odd, transfer i to channel c, if i is even, print i directly
	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		go pingWithChannel(i, c)
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// // read from the channel c
	// fmt.Println(<-c)

	// // print hello world
	// fmt.Println("Hello World")

	// // open a thread to print "tick" every second
	// go func() {
	// 	//create a infinite loop
	// 	for {
	// 		fmt.Println("tick")
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// // wait forever
	// select {}

	// turn off the timer
	// timer.Stop()

	// // create timer to wait for 1 second and print "tick" after every second
	// for i := 0; i < 10; i++ {
	// 	timer := time.NewTimer(1 * time.Second)
	// 	<-timer.C
	// 	fmt.Println("tick")
	// }

	// module in go
	// fmt.Println(math.Sqrt(4))

	// // create a random number generator
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(100))

	// rpc (remote produce call) in go
	rpc.Register()
	rpc.Dial()
	rpc.Serve()
	rpc.NewClient()
	rpc.NewServer()
	rpc.NewClientWithCodec()
	rpc.NewServerCodec()
	rpc.RegisterName()
	rpc.Register()



	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	type Arith int

	func (t *Arith) Multiply(args *Args, reply *int) error {
		*reply = args.A * args.B
		return nil
	}

	func (t *Arith) Divide(args *Args, quo *Quotient) error {
		if args.B == 0 {
			return errors.New("divide by zero")
		}
		quo.Quo = args.A / args.B
		quo.Rem = args.A % args.B
		return nil
	}

}

func ping(i int) {
	fmt.Println(i)
}

func pingWithChannel(i int, c chan int) {
	c <- i
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
