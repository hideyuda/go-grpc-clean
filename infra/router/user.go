package router

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"

	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	// Convert context.Context to echo.Context in gRPC server

	fmt.Println("Create")

	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	// err := bindAndValidate(c, req)
	// if err != nil {
	// 	return nil, err
	// }

	input := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// input := entity.User{
	// 	Id:        uint(req.Id),
	// 	Email:     req.Email,
	// 	CreatedAt: req.CreatedAt.AsTime(),
	// }

	tx, _ := db.Begin()
	h := di.InitializeUserHandler(tx, firebase)
	presenter, err := h.Create(input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println(h)
	fmt.Println(presenter)

	return &pb.UserResponse{}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Get")

	var (
		db       = database.NewDb()
		firebase = driver.NewFirebaseImpl()
	)

	tx, _ := db.Begin()
	h := di.InitializeUserHandler(tx, firebase)
	presenter, err := h.GetById(uint(req.Id))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println(h)
	fmt.Println(presenter)

	return &pb.UserResponse{}, nil
}

// func (s *Server) GetMessageStream(req *emptypb.Empty, server pb.ChatService_GetMessageStreamServer) error {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	stream := make(chan entity.Chat)

// 	go func() {
// 		defer close(stream)
// 		eg, _ := errgroup.WithContext(ctx)
// 		eg.Go(func() error {
// 			// if err := g.messageRepository.Listen(ctx, stream); err != nil {
// 			// 	return err
// 			// }
// 			return nil
// 		})
// 		// if err := eg.Wait(); err != nil {
// 		// 	return fmt.Errorf("failed to GetMessageStreamService.Handle: %s", err)
// 		// }
// 		// return nil
// 		// if err := s.GetMessageStreamService.Handle(ctx, stream); err != nil {
// 		// 	log.Println(err)
// 		// }
// 	}()

// 	for {
// 		v := <-stream
// 		createdAt := timestamppb.New(v.CreatedAt)
// 		if err := server.Send(&pb.GetMessageStreamResponse{
// 			Message: &pb.Message{
// 				From:      v.From,
// 				Content:   v.Content,
// 				CreatedAt: createdAt,
// 			},
// 		}); err != nil {
// 			return err
// 		}
// 	}
// }
