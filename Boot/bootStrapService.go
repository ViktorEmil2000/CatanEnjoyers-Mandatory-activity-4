package Boot

import (
	"sync"

	grpc "google.golang.org/grpc"
)

type BootServerService struct {
}

var mu = sync.Mutex{}

type Nodes struct {
	id     string
	port   int64
	stream grpc.BidiStreamingServer[ClientToBoot, BootToClient]
}

var NodeList = []Nodes{}

func (BSS *BootServerService) mustEmbedUnimplementedBootServer() {
	panic("unimplemented")
}
func (BSS *BootServerService) BootstrapConnect(stream grpc.BidiStreamingServer[ClientToBoot, BootToClient]) error {
	errch := make(chan error)
	node, _ := stream.Recv()
	mu.Lock()

	for _, element := range NodeList {
		element.stream.Send(&BootToClient{
			Id:   node.Id,
			Port: node.Port,
		})
	}

	for _, element := range NodeList {
		stream.Send(&BootToClient{
			Id:   element.id,
			Port: element.port,
		})
	}

	NodeList = append(NodeList, Nodes{
		id:     node.Id,
		port:   node.Port,
		stream: stream,
	})
	mu.Unlock()
	return <-errch
}
