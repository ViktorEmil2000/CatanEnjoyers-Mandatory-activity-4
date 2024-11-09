package p2p

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
)

var (
	streamToClient grpc.BidiStreamingServer[FromClient, FromServer]
)

type p2pServerService struct {
}

func (p2p *p2pServerService) mustEmbedUnimplementedP2PServer() {
	panic("unimplemented")
}

func (p2p *p2pServerService) ClientConnect(stream grpc.BidiStreamingServer[FromClient, FromServer]) error {
	errch := make(chan error)

	streamToClient = stream

	log.Printf("Server sucessfully connectd @")

	return <-errch
}

func (p2p *p2pServerService) RequestResponse(context.Context, *ResquestFromClient) (*ResponseFromServer, error) {
	request := ResquestFromClient{}.Id
	streamToClient.Send(&FromServer{Id: request})

	clientresponse, _ := streamToClient.Recv()

	return &ResponseFromServer{Response: clientresponse.Response}, nil

}
