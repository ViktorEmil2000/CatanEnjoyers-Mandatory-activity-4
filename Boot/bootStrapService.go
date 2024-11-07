package Boot

import (
	grpc "google.golang.org/grpc"
)

type BootServerService struct {
}

type Nodes struct {
	id     string
	stream grpc.BidiStreamingServer[ClientToBoot, BootToClient]
}

var NodeList = []Nodes{}

func (BSS *BootServerService) mustEmbedUnimplementedBootServer() {
	panic("unimplemented")
}
func (BSS *BootServerService) BootstrapConnect(stream grpc.BidiStreamingServer[ClientToBoot, BootToClient]) error {
	
	node,_ := stream.Recv()

	for

	NodeList = append(NodeList,Nodes{
		id: node.Id,
		stream: stream,
	}) 



}
