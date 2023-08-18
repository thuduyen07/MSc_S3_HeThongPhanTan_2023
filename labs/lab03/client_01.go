package main

import (
	"fmt"
	"net/rpc"
	"time"
)

type Args struct {
	Data string
}

type Args struct {
	Ping string
}

func IsServerActive(client *rpc.Client) bool {
	pingMessage := "Ping from client"
	var reply string

	err := client.Call("Server.PingPong", Args{pingMessage}, &reply)
	if err != nil {
		fmt.Println("Error calling server:", err)
		return false
	}

	return reply == "Pong from the server"
}

func main() {
	client, err := rpc.Dial("tcp", "server1_address:port")
	if err != nil {
		fmt.Println("Error connecting to Server 1:", err)
		return
	}

	args := Args{"Hello from Client 1"}
	var reply string
	err = client.Call("Server.ProcessData", args, &reply)
	if err != nil {
		fmt.Println("Error calling Server 1:", err)
		return
	}
	fmt.Println("Server 1 Response:", reply)

	client.Close()
}
