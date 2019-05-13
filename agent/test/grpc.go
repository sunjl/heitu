package main

import (
	pb "github.com/jishufan/heitu/common/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userServiceClient := pb.NewUserServiceClient(conn)

	var userId int
	if len(os.Args) > 1 {
		userId, err = strconv.Atoi(os.Args[1])
	}
	user, err := userServiceClient.Get(context.Background(), &pb.User{Id: int32(userId)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("--GetUser--\n%v", user)
}
