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

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8100")
	if err != nil {
		fmt.Println("Error connecting to Secondary Server:", err)
		return
	}
	defer client.Close()

	args := Args{"Hello from Client_02"}

	for {
		var reply string
		err := client.Call("Server.ProcessData", args, &reply)
		if err != nil {
			fmt.Println("Error calling secondary server:", err)
		}

		fmt.Println("Secondary Server Response:", reply)
		time.Sleep(time.Second)
	}
}
