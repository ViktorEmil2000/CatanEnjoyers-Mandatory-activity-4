package p2p

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
)

var (
	streamToClient grpc.BidiStreamingServer[FromClient, FromServer]
)

type P2PServerService struct {
}

func (p2p *P2PServerService) mustEmbedUnimplementedP2PServer() {
	panic("unimplemented")
}

func (p2p *P2PServerService) ClientConnect(stream grpc.BidiStreamingServer[FromClient, FromServer]) error {
	errch := make(chan error)

	streamToClient = stream

	log.Printf("Server sucessfully connectd to client in node @")

	return <-errch
}

func (p2p *P2PServerService) RequestResponse(context.Context, *ResquestFromClient) (*ResponseFromServer, error) {
	request := ResquestFromClient{}.Id
	streamToClient.Send(&FromServer{Id: request})

	clientresponse, _ := streamToClient.Recv()

	return &ResponseFromServer{Response: clientresponse.Response}, nil

}
