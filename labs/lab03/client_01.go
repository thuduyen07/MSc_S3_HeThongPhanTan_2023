package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	Data string
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
