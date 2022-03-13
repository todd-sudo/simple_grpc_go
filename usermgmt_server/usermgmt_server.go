package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/todd-sudo/simple_grpc_go/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var userId int32 = int32(rand.Intn(1000))
	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   userId,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})

	log.Printf("server listening as %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
