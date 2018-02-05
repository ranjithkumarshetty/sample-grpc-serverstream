package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/ranjithkumarshetty/sample-grpc-serverstream/protos"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:4443", "The server address in the format of host:port")
)

func receiveMessages(client pb.StreamerClient) {
	log.Print("Looking for messages")
	stream, err := client.StreamMessages(context.Background(), &pb.Message{"Hello from client"})
	if err != nil {
		log.Fatalf("Error while receving message :%v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receving message :%v", err)
		}
		log.Println(msg)
	}
	log.Print("Got EOF, closing client")
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewStreamerClient(conn)
	receiveMessages(client)
}
