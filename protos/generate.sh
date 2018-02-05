#!/bin/sh

set -ev

go get github.com/golang/protobuf/{proto,protoc-gen-go}

protoc --go_out=plugins=grpc:. streamer.proto
