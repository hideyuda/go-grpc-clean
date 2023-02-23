package repository

import (
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
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
func (r *ChatRepositoryImpl) Create(param *pb.Chat) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chats (
			uuid,
			group_id,
			name,
			from,
			to,
			content,
			created_at,
			updated_at,
			is_deleted,
			) VALUES (
				?,
				?, 
				?,
				?,
				?,
				?,
				?,
				?,
				?
		)`,
		utils.CreateUuid(),
		param.GroupId,
		param.Name,
		param.From,
		param.To,
		param.Content,
		now,
		now,
		false,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Update *****/
func (r *ChatRepositoryImpl) Update(param *pb.Chat) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE chats SET
			name = ?,
			to = ?,
			content = ?,
			updated_at = ?`,
		param.Name,
		param.To,
		param.Content,
		time.Now(),
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
		"UPDATE chats SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatRepositoryImpl) GetById(id uint) (Chat *pb.Chat, err error) {
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

// get list by group id
func (r *ChatRepositoryImpl) GetListByGroupId(groupId uint) (Chats []*pb.Chat, err error) {
	err = r.executer.Select(
		r.Name+"GetListByGroupId",
		&Chats,
		`
		SELECT * 
		FROM chats 
		WHERE group_id = ? AND is_deleted = false
		ORDER BY id DESC
		`,
		groupId,
	)

	if err != nil {
		return nil, err
	}

	return Chats, nil
}
