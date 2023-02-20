package responses

import (
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser(res *pb.User) *pb.UserResponse {
	return &pb.UserResponse{
		// Error: false,
		User: &pb.User{
			Id:         uint32(res.Id),
			Uuid:       res.Uuid,
			FirebaseId: res.FirebaseId,
			Name:       res.Name,
			Email:      res.Email,
			Password:   res.Password,
			UserType:   pb.User_UserType(res.UserType),
			CreatedAt: res.CreatedAt,
			UpdatedAt:  res.UpdatedAt,
			// CreatedAt:  timestamppb.New(res.CreatedAt),
			// UpdatedAt:  timestamppb.New(res.UpdatedAt),
		},
	}
}

// func NewUserList(ress []*entity.User) []*pb.User {
// 	var resList []*pb.User
// 	for _, res := range ress {
// 		resList = append(resList, NewUser(res))
// 	}
// 	return resList
// }
