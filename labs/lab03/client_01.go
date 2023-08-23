package lab03

import (
	"fmt"
	"net/rpc"
	"time"
)

type Args struct {
	Data string
	Ping string
}

func IsServerActive(client *rpc.Client) bool {
	pingMessage := "Ping from client_01"
	var reply string

	err := client.Call("Server.PingPong", Args{pingMessage}, &reply)
	if err != nil {
		fmt.Println("Error calling server:", err)
		return false
	}

	return reply == "Pong from the server"
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error connecting to Primary Server :", err)
		return
	}

	args := Args{"Hello from Client_01"}
	var reply string
	err = client.Call("Server.ProcessData", args, &reply)
	if err != nil {
		fmt.Println("Error calling Primary Server:", err)
		return
	}
	fmt.Println("Primary Server Response:", reply)

	client.Close()
}
