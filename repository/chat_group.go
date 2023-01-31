package repository

import (
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utility"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatGroupRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewChatGroupRepositoryImpl(ex SQLExecuter) usecase.ChatGroupRepository {
	return &ChatGroupRepositoryImpl{
		Name:     "ChatGroupRepository",
		executer: ex,
	}
}

/***** Create *****/
func (r *ChatGroupRepositoryImpl) Create(param *entity.ChatGroup) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chat_groups (
			uuid,
			name,
			created_at,
			updated_at
			) VALUES (
				?,
				?, 
				?
		)`,
		utility.CreateUUID(),
		param.Name,
		now,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Update *****/
func (r *ChatGroupRepositoryImpl) Update(param *entity.ChatGroup) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE chat_groups SET
			name = ?,
			updated_at = ?`,
		param.Name,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Delete *****/
func (r *ChatGroupRepositoryImpl) Delete(id uint) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"DELETE FROM chat_groups WHERE id = ?",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatGroupRepositoryImpl) GetById(id uint) (ChatGroup *entity.ChatGroup, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		ChatGroup,
		"SELECT * FROM chat_groups WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return ChatGroup, nil
}
