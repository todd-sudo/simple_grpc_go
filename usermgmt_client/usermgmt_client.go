package main

import (
	"context"
	"log"
	"time"

	pb "github.com/todd-sudo/simple_grpc_go/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Bob"] = 30

	for name, age := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
