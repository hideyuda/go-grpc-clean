package requests

import (
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
)

func NewUser(req *pb.User) (*entity.User, error) {

	uuid, err := utils.ParseUuid(req.Uuid)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:         uint(req.Id),
		Uuid:       uuid,
		FirebaseId: req.FirebaseId,
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		UserType:   uint(req.UserType),
		CreatedAt:  req.CreatedAt.AsTime(),
		UpdatedAt:  req.UpdatedAt.AsTime(),
	}, nil
}

// func NewUserList(reqs []*entity.User) []*pb.User {
// 	var reqList []*pb.User
// 	for _, req := range reqs {
// 		reqList = append(reqList, NewUser(req))
// 	}
// 	return reqList
// }
