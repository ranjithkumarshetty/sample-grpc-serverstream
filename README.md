[![Build Status](https://travis-ci.com/ranjithkumarshetty/sample-grpc-serverstream.svg?branch=master)](https://travis-ci.com/ranjithkumarshetty/sample-grpc-serverstream)

# Streamer
Demo of gRPC-based client/server pair implemented in Go (In this case Server streaming in RPC)

## Setup:
#### 1. Setup Go and gRPC
Setup is described in more details in https://grpc.io/docs/quickstart/go.html
* Install Go - https://golang.org/dl/ 
* Install protoc - `brew install grpc` or from https://github.com/google/protobuf/releases
* `go get -u google.golang.org/grpc`
* `go get -u github.com/golang/protobuf/protoc-gen-go`

#### 2. Checkout the code
`git clone git@github.com:ranjithkumarshetty/sample-grpc-serverstream`

#### 3. Generate .go file from .proto
`go generate -v ./...`

#### 4. Build client and server code
`cd client; go build; cd ../server; go build;`

#### 5. Generate certificate for the server
`cd server; cd openssl req -newkey rsa:2048 -nodes -keyout server_key.pem -x509 -days 365 -out certificate.pem;`

`cp certificate.pem ../client/`

#### 6. Run server in one terminal
`./server -tls=true`

#### 7. Run client in the other terminal
`./client -tls=true`
