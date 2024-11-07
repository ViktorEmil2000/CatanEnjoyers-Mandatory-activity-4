package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	"google.golang.org/grpc"
	//50051
)

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

	}

}
