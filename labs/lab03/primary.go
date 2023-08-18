package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	Data string
}

type Server struct{}

func (s *Server) ProcessData(args Args, reply *string) error {
	*reply = "Server 1 processed: " + args.Data
	return nil
}

func main() {
	server := new(Server)
	rpc.Register(server)

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error starting Server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server 1 listening on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
