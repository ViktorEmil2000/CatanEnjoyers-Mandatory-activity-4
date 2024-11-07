package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	//50051
)

func main() {

	fmt.Println("Enter port:")
	reader := bufio.NewReader(os.Stdin)
	port, _ := reader.ReadString('\n')

	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewServicesClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}
	ch := clienthandle{
		stream:   stream,
		username: username,
	}

	go sendMessage(ch)
	go receiveMessage(ch)

	initialMessage := &proto.FromClient{
		Name: ch.username,
		Body: "Connected!",
	}
	ch.stream.Send(initialMessage)

	bl := make(chan bool)
	<-bl
}

type clienthandle struct {
	stream   proto.Services_ChatServiceClient
	username string
}

func sendMessage(ch clienthandle) {
	for {
		reader := bufio.NewReader(os.Stdin)

		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &proto.FromClient{
			Name: ch.username,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)
		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}
func receiveMessage(ch clienthandle) {
	for {
		msg, _ := ch.stream.Recv()

		fmt.Printf("%s: %s  \n", msg.Name, msg.Body)
	}

}
