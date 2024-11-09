package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	P2P "github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/p2p"
	"google.golang.org/grpc"
)

var (
	streamToServer grpc.BidiStreamingClient[P2P.FromClient, P2P.FromServer]
)

type Node struct {
	id     int64
	client P2P.P2PClient
}

var NodeList = []Node{}

var nodeLock = sync.Mutex{}

func main() {

	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	userId := rand.Intn(10000000)

	go bootstrap(port, int64(userId))

	//Create Node Server
	PortForServer := os.Getenv("PORT")
	if PortForServer == "" {
		PortForServer = port
	}

	listen, _ := net.Listen("tcp", ":"+PortForServer)
	log.Println("Listening @ : " + PortForServer)

	grpcserver := grpc.NewServer()

	cs := P2P.p2pServerService{}
	P2P.RegisterP2PServer(grpcserver, &cs)

	grpcserver.Serve(listen)

	//Create Node Client
	conn, _ := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	defer conn.Close()
	Client := P2P.NewP2PClient(conn)
	stream, _ := Client.ClientConnect(context.Background())
	streamToServer = stream

	bl := make(chan bool)

	<-bl
}

func bootstrap(port string, userId int64) {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	BootClient := Boot.NewBootClient(conn)

	stream, _ := BootClient.BootstrapConnect(context.Background())

	stream.Send(&Boot.ClientToBoot{
		Id:   userId,
		Port: port,
	})

	for {
		node, _ := stream.Recv()

		nodeLock.Lock()

		conn, _ = grpc.Dial("localhost:"+node.Port, grpc.WithInsecure())
		defer conn.Close()
		nodeCLient := p2p.NewP2PClient(conn)

		newNode := Node{
			id:     node.Id,
			client: nodeCLient,
		}

		NodeList = append(NodeList, newNode)
		log.Printf("New node was discovered ID:%v", node.Id)
		nodeLock.Unlock()
	}

}
