package repository

import (
	"fmt"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewChatRepositoryImpl(ex SQLExecuter) usecase.ChatRepository {
	return &ChatRepositoryImpl{
		Name:     "ChatRepository",
		executer: ex,
	}
}

/***** Create *****/
func (r *ChatRepositoryImpl) Create(param *entity.Chat) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chats (
			uuid,
			content,
			created_at,
			updated_at
			) VALUES (
				?,
				?,
				?, 
				?
		)`,
		utils.CreateUUID(),
		param.Content,
		now,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Delete *****/
func (r *ChatRepositoryImpl) Delete(id uint) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"DELETE FROM chats WHERE id = ?",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatRepositoryImpl) GetById(id uint) (Chat *entity.Chat, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		Chat,
		"SELECT * FROM chats WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return Chat, nil
}

// getListByGroupId
func (r *ChatRepositoryImpl) GetListByGroupId(groupId uint) ([]*entity.Chat, error) {
	var (
		chats []*entity.Chat
	)
	err := r.executer.Select(
		r.Name+"GetListByGroupId",
		&chats,
		"SELECT * FROM chats WHERE group_id = ? ORDER BY id DESC",
		groupId,
	)

	if err != nil {
		err = fmt.Errorf("failed to get all chats: %w", err)
		fmt.Println(err)
		return nil, err
	}

	return chats, nil
}
