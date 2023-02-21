package responses

import (
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser(res *entity.User) *pb.User {
	return &pb.User{
		Id:         int64(res.Id),
		Uuid:       res.Uuid.String(),
		FirebaseId: res.FirebaseId,
		Name:       res.Name,
		Email:      res.Email,
		Password:   res.Password,
		UserType:   pb.User_UserType(res.UserType),
		CreatedAt:  timestamppb.New(res.CreatedAt),
		UpdatedAt:  timestamppb.New(res.UpdatedAt),
	}
}

func NewUserResponse(res *pb.User) *pb.UserResponse {
	return &pb.UserResponse{
		Error: false,
		User:  res,
	}
}

// func NewUserList (res []*pb.User) *pb.UserListResponse {
// 	return &pb.UserListResponse{
// 		Error: false,
// 		UserList: res,
// 	}
// }

// func NewUserList(ress []*entity.User) []*pb.User {
// 	var resList []*pb.User
// 	for _, res := range ress {
// 		resList = append(resList, NewUser(res))
// 	}
// 	return resList
// }

// convert timestamp to time.Time
// CreatedAt:  req.CreatedAt.AsTime(),

// convert time.Time to timestamp
// timestamppb.New(req.CreatedAt),
