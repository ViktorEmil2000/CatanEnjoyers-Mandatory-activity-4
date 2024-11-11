package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
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
	port = strings.TrimSpace(port)
	userId := rand.Intn(10000000)

	go bootstrap(port, int64(userId))

	go initializeServer(port, userId)

	bl := make(chan bool)

	<-bl
}

func initializeServer(port string, userid int) {

	log.Printf(port)
	//Create Node Server
	PortForServer := os.Getenv("PORT")
	if PortForServer == "" {
		PortForServer = port
	}

	listen, err := net.Listen("tcp", ":"+PortForServer)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", PortForServer, err)
	}
	log.Println("Listening @ : " + PortForServer)

	grpcserver := grpc.NewServer()

	cs := p2p.P2PServerService{}
	p2p.RegisterP2PServer(grpcserver, &cs)

	go initializeClient(port, userid)

	grpcserver.Serve(listen)

}

func initializeClient(port string, userid int) {
	//Create Node Client
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial ServerFrom within node %s:", err)
	}

	Client := p2p.NewP2PClient(conn)
	stream, err := Client.ClientConnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to listen to same stream from node server %s:", err)
	}
	streamToServer = stream
	go nodeServer()

	go tryonCriticalCode(userid)
}
func tryonCriticalCode(id int) {
	for {
		check := false

		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		usingCriticalCode = true
		nodeLock.Lock()

		for _, element := range NodeList {
			response, err := element.client.RequestResponse(context.Background(), &p2p.ResquestFromClient{Id: int64(id)})
			if err != nil {
				log.Fatalf("Failed to request a response from other node %s:", err)
			}

			check = response.Response
			if check == true {
				log.Printf("Other nodes on the network i using the critcal code, so this node is skipping")
				break
			}
			log.Printf("This node got to run the critical code nodeID:%v", id)
			time.Sleep(time.Duration(2000) * time.Millisecond)
		}

		usingCriticalCode = false
		nodeLock.Unlock()
	}
}

func nodeServer() {
	for {
		streamToServer.Recv()

		streamToServer.Send(&p2p.FromClient{Response: usingCriticalCode})
	}
}

func bootstrap(port string, userId int64) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial BootStrap %s:", err)
	}

	BootClient := Boot.NewBootClient(conn)

	stream, _ := BootClient.BootstrapConnect(context.Background())

	stream.Send(&Boot.ClientToBoot{
		Id:   userId,
		Port: port,
	})

	for {
		node, _ := stream.Recv()

		nodeLock.Lock()

		conn, err := grpc.Dial("localhost:"+node.Port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed To dial Other node %s:", err)
		}
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
