package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/p2p"
	"google.golang.org/grpc"
	//50051
)

type Node struct {
	id     int64
	client p2p.ClientClient
}

var NodeList = []Node{}

var nodeLock = sync.Mutex{}

func main() {

	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	userId := rand.Intn(10000000)

	go bootstrap(port, int64(userId))

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
		nodeCLient := p2p.NewClientClient(conn)

		newNode := Node{
			id:     node.Id,
			client: nodeCLient,
		}

		NodeList = append(NodeList, newNode)
		log.Printf("New node was discovered ID:%v", node.Id)
		nodeLock.Unlock()
	}

}
