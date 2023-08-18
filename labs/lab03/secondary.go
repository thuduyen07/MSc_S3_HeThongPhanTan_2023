package lab03

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type Args struct {
	Data string
	Ping string
}

type Server struct {
	Name      string
	IsPrimary bool
	Mutex     sync.Mutex
}

func (s *Server) ProcessData(args Args, reply *string) error {
	if s.IsPrimary {
		*reply = "Primary Server (" + s.Name + ") processed: " + args.Data
	} else {
		*reply = "Secondary Server (" + s.Name + ") processed: " + args.Data
	}
	return nil
}

func (s *Server) SetPrimaryStatus(status bool) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.IsPrimary = status
}

func main() {
	server := &Server{Name: "Secondary", IsPrimary: false}
	rpc.Register(server)

	listener, err := net.Listen("tcp", "127.0.0.1:8100")
	if err != nil {
		fmt.Println("Error starting Secondary Server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Secondary Server listening on port 8100...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
