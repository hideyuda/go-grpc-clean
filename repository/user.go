package repository

import (
	"fmt"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utility"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewUserRepositoryImpl(ex SQLExecuter) usecase.UserRepository {
	return &UserRepositoryImpl{
		Name:     "UserRepository",
		executer: ex,
	}
}

/***** Create *****/
func (r *UserRepositoryImpl) Create(param *entity.User) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"SignUp",
		`INSERT INTO users (
			uuid,
			firebase_id,
			name, 
			email, 
			password,
			user_type,
			created_at,
			updated_at
			) VALUES (
				?,
				?,
				?, 
				?,
				?,
				?,
				?,
				?
		)`,
		utility.CreateUUID(),
		param.FirebaseId,
		param.Name,
		param.Email,
		param.Password,
		param.UserType,
		now,
		now,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Update *****/
func (r *UserRepositoryImpl) Update(user *entity.User) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE users SET
			name = ?,
			email = ?,
			password = ?,
			user_type = ?,
			updated_at = ?
		WHERE _id = ?`,
		user.Name,
		user.Email,
		user.Password,
		user.UserType,
		time.Now(),
		user.Id,
	)

	if err != nil {
		err = fmt.Errorf("failed to update user: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) UpdateColumnStr(lineUserId, column, value string) error {
	_, err := r.executer.Exec(
		r.Name+"UpdateColumn",
		"UPDATE users SET "+column+" = ? WHERE line_user_id = ?",
		value,
		lineUserId,
	)

	if err != nil {
		err = fmt.Errorf("failed to update user column: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) UpdateColumnInt(lineUserId, column string, value int) error {
	_, err := r.executer.Exec(
		r.Name+"UpdateColumn",
		"UPDATE users SET "+column+" = ? WHERE line_user_id = ?",
		value,
		lineUserId,
	)

	if err != nil {
		err = fmt.Errorf("failed to update user column: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

/***** Get *****/
func (r *UserRepositoryImpl) GetById(id uint) (user *entity.User, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		user,
		"SELECT * FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) SignIn(email, password string) (user *entity.User, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		user,
		"SELECT * FROM users WHERE email = ? AND password = ?",
		email,
		password,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByFirebaseId(firebaseId string) (*entity.User, error) {
	var (
		user entity.User
	)
	err := r.executer.Get(
		r.Name+"GetByFirebaseId",
		&user,
		"SELECT * FROM users WHERE firebase_id = ?",
		firebaseId,
	)

	if err != nil {
		err = fmt.Errorf("failed to get user by firebase id: %w", err)
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

// getAll
func (r *UserRepositoryImpl) GetAll() ([]*entity.User, error) {
	var (
		users []*entity.User
	)
	err := r.executer.Select(
		r.Name+"GetAll",
		&users,
		"SELECT * FROM users ORDER BY id DESC",
	)

	if err != nil {
		err = fmt.Errorf("failed to get all users: %w", err)
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}
