package Boot

import (
	"log"
	"net"
	"os"

	"github.com/ViktorEmil2000/CatanEnjoyers-Mandatory-activity-4/Boot"
	"google.golang.org/grpc"
)

func main() {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "50051"
	}

	listen, _ := net.Listen("tcp", ":"+Port)
	log.Println("Listening @ : " + Port)

	grpcserver := grpc.NewServer()

	cs := Boot.BootServerService{}
	Boot.RegisterBootServer(grpcserver, &cs)

	grpcserver.Serve(listen)
}
