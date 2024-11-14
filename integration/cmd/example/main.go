package main

import (
	"fmt"
	"github.com/dimoktorr/unit_tests/integration/internal/pkg/api"
	v1 "github.com/dimoktorr/unit_tests/integration/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()

	serverApi := &api.Server{}
	server.RegisterService(&v1.ExampleService_ServiceDesc, serverApi)

	l, err := NewTCPListener("localhost", "5390")
	if err != nil {
		log.Fatal(err)
	}

	_ = server.Serve(l)
}

func NewTCPListener(host, port string) (net.Listener, error) {
	addr := host + ":" + port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %q: %w", addr, err)
	}

	return l, nil
}
