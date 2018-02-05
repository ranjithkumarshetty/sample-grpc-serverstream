package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ranjithkumarshetty/sample-grpc-serverstream/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 4443, "The server port")
)

type streamerServer struct {
}

func (s *streamerServer) StreamMessages(msg *pb.Message, srv pb.Streamer_StreamMessagesServer) error {
	log.Printf("Received message from client:%v\n", msg)
	i := 1
	for 1 == 1 {
		if err := srv.Send(&pb.Message{fmt.Sprintf("Server hello reply times:%v", i)}); err != nil {
			return err
		}
		i++
		time.Sleep(2000 * time.Millisecond)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	srv := &streamerServer{}
	pb.RegisterStreamerServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
