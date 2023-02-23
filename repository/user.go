package repository

import (
	"fmt"
	"time"

	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
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

// create
func (r *UserRepositoryImpl) Create(param *pb.User) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
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
		utils.CreateUuid(),
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

// update
func (r *UserRepositoryImpl) Update(user *pb.User) error {
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

// delete
func (r *UserRepositoryImpl) Delete(id uint) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE users SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// get
func (r *UserRepositoryImpl) GetById(id uint) (*pb.User, error) {
	var (
		// userEntity entity.User
		user pb.User
	)

	err := r.executer.Get(
		r.Name+"GetById",
		// &userEntity,
		&user,
		"SELECT * FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// auth
// SignIn
func (r *UserRepositoryImpl) SignIn(email, password string) (user *pb.User, err error) {
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

// getByFirebaseId
func (r *UserRepositoryImpl) GetByFirebaseId(firebaseId string) (*pb.User, error) {
	var (
		user pb.User
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

// admin
// getAll
func (r *UserRepositoryImpl) GetAll() ([]*pb.User, error) {
	var (
		users []*pb.User
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
