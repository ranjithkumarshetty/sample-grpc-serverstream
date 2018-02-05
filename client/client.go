// Server code to run server in tls or non-tls mode based on argument.
// Implements receiveMessages() to read messages from Server stream over gRPC.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/ranjithkumarshetty/sample-grpc-serverstream/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "ranjithkumarshetty.com", "The server name use to verify the hostname returned by TLS handshake")
	serverAddr         = flag.String("server_addr", "127.0.0.1:4443", "The server address in the format of host:port")
)

func receiveMessages(client pb.StreamerClient) {
	clientID := makeTimestamp()
	log.Printf("ClientID:\"%v\". Looking for messages", clientID)
	stream, err := client.StreamMessages(context.Background(), &pb.Message{fmt.Sprintf("Hello from clientID:%v", clientID)})
	if err != nil {
		log.Fatalf("ClientID:\"%v\". Error while receving message :%v", clientID, err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ClientID:\"%v\". Error while receving message :%v", clientID, err)
		}
		log.Println(msg)
	}
	log.Printf("ClientID:\"%v\". Got EOF, closing client", clientID)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "certificate.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewStreamerClient(conn)
	receiveMessages(client)
}
