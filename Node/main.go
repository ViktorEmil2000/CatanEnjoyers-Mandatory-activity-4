package main

import (
	"Boot"
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	p2p "github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/P2P"
	"google.golang.org/grpc"
	//50051
)

func main() {

	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')
	userId := rand.Intn(10000000)

	go bootstrap(port, userId)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

	bl := make(chan bool)
	<-bl
}

type clienthandle struct {
	stream   p2p.ClientClient
	username string
}

func bootstrap(port string, userId int64) {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	BootClient := Boot.NewBootClient(conn)

	stream, _ := BootClient.BootStrapConnect(context.Background())

}
