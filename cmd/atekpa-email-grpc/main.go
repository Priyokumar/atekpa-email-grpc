package main

import (
	"log"
	"net"

	"github.com/Priyokumar/atekpa-email-grpc/api/pb"
	"github.com/Priyokumar/atekpa-email-grpc/internal/email"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalln("server failed to start")
	}

	grpc := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpc, email.New())

	err = grpc.Serve(lis)
	if err != nil {
		log.Fatalln("grpc server failed to start")
	}
}
