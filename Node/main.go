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
	"time"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/p2p"
	"google.golang.org/grpc"
)

var (
	streamToServer    grpc.BidiStreamingClient[p2p.FromClient, p2p.FromServer]
	usingCriticalCode bool
)

type Node struct {
	id     int64
	client p2p.P2PClient
}

var NodeList = []Node{}

var nodeLock = sync.Mutex{}

func main() {
	usingCriticalCode = false
	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	userId := rand.Intn(10000000)

	go bootstrap(port, int64(userId))

	go initialize(port, userId)

	bl := make(chan bool)

	<-bl
}

func initialize(port string, userid int) {
	//Create Node Server
	PortForServer := os.Getenv("PORT")
	if PortForServer == "" {
		PortForServer = port
	}

	listen, _ := net.Listen("tcp", ":"+PortForServer)
	log.Println("Listening @ : " + PortForServer)

	grpcserver := grpc.NewServer()

	cs := p2p.P2PServerService{}
	p2p.RegisterP2PServer(grpcserver, &cs)
	//go initializeClient(port, userid)

	grpcserver.Serve(listen)
}

func initializeClient(port string, userid int) {
	//Create Node Client
	conn, _ := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	defer conn.Close()
	Client := p2p.NewP2PClient(conn)
	stream, _ := Client.ClientConnect(context.Background())
	streamToServer = stream

	go nodeServer()
	go tryonCriticalCode(userid)
}
func tryonCriticalCode(id int) {
	for {
		check := false

		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		usingCriticalCode = true

		for _, element := range NodeList {
			response, _ := element.client.RequestResponse(context.Background(), &p2p.ResquestFromClient{Id: int64(id)})

			check = response.Response
			if check == true {
				break
			}
		}
		if check == false {
			log.Printf("This node got to run the critical code nodeID:%v", id)
		}
		usingCriticalCode = false

	}
}

func nodeServer() {
	for {
		streamToServer.Recv()

		streamToServer.Send(&p2p.FromClient{Response: usingCriticalCode})
	}
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
