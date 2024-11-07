package proto

import (
	grpc "google.golang.org/grpc"
)

type P2PBootServer struct {
}

type Nodes struct {
	id     string
	stream grpc.BidiStreamingServer[BootClient, BootServer]
}

func (BSC *P2PBootServer) mustEmbedUnimplementedP2PServer() {
	panic("unimplemented")
}
func (BSC *P2PBootServer) BootstrapConnect(stream grpc.BidiStreamingServer[BootClient, BootServer]) error {
	errch := make(chan error)

	return <-errch
}

func sendall(id string) {

}
