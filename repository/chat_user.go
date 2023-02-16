package repository

import (
	"fmt"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
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
func (r *ChatUserRepositoryImpl) Create(param *entity.ChatUser) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chat_users (
			uuid,
			user_id,
			created_at,
			updated_at
			) VALUES (
				?,
				?,
				?, 
				?
		)`,
		utils.CreateUUID(),
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
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatUserRepositoryImpl) GetById(id uint) (ChatUser *entity.ChatUser, err error) {
	err = r.executer.Get(
		r.Name+"GetById",
		ChatUser,
		"SELECT * FROM chat_users WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return ChatUser, nil
}

// getListByGroupId
func (r *ChatUserRepositoryImpl) GetListByGroupId(groupId uint) ([]*entity.ChatUser, error) {
	var (
		chat_users []*entity.ChatUser
	)
	err := r.executer.Select(
		r.Name+"GetListByGroupId",
		&chat_users,
		"SELECT * FROM chat_users WHERE group_id = ? ORDER BY id DESC",
		groupId,
	)

	if err != nil {
		err = fmt.Errorf("failed to get all chat_users: %w", err)
		fmt.Println(err)
		return nil, err
	}

	return chat_users, nil
}
