package request

import (
	"context"
	"log"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserRequestImpl struct {
}

func NewUserRequestImpl() UserRequest {
	return &UserRequestImpl{}
}

type UserRequest interface {
	DetectTextFromImage()
}

func (r *UserRequestImpl) DetectTextFromImage() {
	var (
		conn *grpc.ClientConn
		err  error
	)

	// Set up a connection to the server.
	conn, err = grpc.Dial(config.App.PythonDomain, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := c.CreateUser(ctx, &pb.User{Name: "name", Email: "email", Password: "password"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", req.GetMessage())
}
