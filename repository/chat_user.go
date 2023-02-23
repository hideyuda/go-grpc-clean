package repository

import (
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatUserRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewChatUserRepositoryImpl(ex SQLExecuter) usecase.ChatUserRepository {
	return &ChatUserRepositoryImpl{
		Name:     "ChatUserRepository",
		executer: ex,
	}
}

/***** Create *****/
func (r *ChatUserRepositoryImpl) Create(param *pb.ChatUser) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chat_users (
			uuid,
			group_id,
			user_id,
			created_at,
			updated_at,
			) VALUES (
				?,
				?, 
				?,
				?,
				?
		)`,
		utils.CreateUuid(),
		param.GroupId,
		param.UserId,
		now,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Delete *****/
func (r *ChatUserRepositoryImpl) Delete(id uint) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"DELETE FROM chat_users WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatUserRepositoryImpl) GetById(id uint) (ChatUser *pb.ChatUser, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		ChatUser,
		"SELECT * FROM chat_users WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return ChatUser, nil
}

// get list by user id
func (r *ChatUserRepositoryImpl) GetListByGroupId(groupId uint) (ChatUsers []*pb.ChatUser, err error) {
	err = r.executer.Select(
		r.Name+"GetListByGroupId",
		&ChatUsers,
		`
		SELECT * 
		FROM chat_users 
		WHERE group_id = ?
		ORDER BY id DESC
		`,
		groupId,
	)

	if err != nil {
		return nil, err
	}

	return ChatUsers, nil
}
