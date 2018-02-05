package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ranjithkumarshetty/sample-grpc-serverstream/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 4443, "The server port")
)

type streamerServer struct {
}

func (s *streamerServer) StreamMessages(msg *pb.Message, srv pb.Streamer_StreamMessagesServer) error {
	log.Printf("Received message from client:\"%s\"", msg)
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
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "certificate.pem"
		}
		if *keyFile == "" {
			*keyFile = "server_key.pem"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	srv := &streamerServer{}
	pb.RegisterStreamerServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
