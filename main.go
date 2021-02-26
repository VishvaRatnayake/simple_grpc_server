package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"simple_grpc/pb"

	"google.golang.org/grpc"
)

var err error
var ctx context.Context

type HelloServiceServer struct{}

var p1 = pb.Person{
	Id:   1,
	Name: "John Doe",
	Age:  35,
}

var p2 = pb.Person{
	Id:   2,
	Name: "Cerberus Nastasya",
	Age:  38,
}

var p3 = pb.Person{
	Id:   3,
	Name: "Backer Smith",
	Age:  38,
}

var list = []*pb.Person{&p1, &p2, &p3}

func (s *HelloServiceServer) GetInfo(ctx context.Context, request *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {

	fmt.Printf("requested id is : %v \n", request.Id)

	for _, s := range list {
		if request.Id == s.Id {
			return &pb.GetInfoResponse{
				StatusCode: 200,
				Person:     s,
			}, nil
		}
	}

	return &pb.GetInfoResponse{
		StatusCode: 500,
		Person:     nil,
	}, nil
}

func main() {
	address := "localhost:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v... \n", address)
	s := grpc.NewServer()

	userService := &HelloServiceServer{}
	pb.RegisterHelloServiceServer(s, userService)
	s.Serve(lis)
}
