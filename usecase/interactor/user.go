package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/infra/request"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserInteractor interface {
	// Gest API
	// Create
	Create(user *entity.User) (*entity.User, error)

	// Update
	Update(user *entity.User) (*entity.User, error)

	// Get
	GetById(id uint) (*entity.User, error)
	SignIn(param *entity.SignInParam) (*entity.User, error)
	GetByFirebaseToken(token string) (*entity.User, error)

	// admin API
	GetAll() ([]*entity.User, error)
}

type UserInteractorImpl struct {
	firebase       usecase.Firebase
	userRepository usecase.UserRepository
}

func NewUserInteractorImpl(
	fb usecase.Firebase,
	uR usecase.UserRepository,
) UserInteractor {
	return &UserInteractorImpl{
		firebase:       fb,
		userRepository: uR,
	}
}

// Create
// type Createuser struct {
// 	Param *entity.User
// }

// type CreateOutput struct {
// 	Ok bool
// }

func (i *UserInteractorImpl) Create(user *entity.User) (*entity.User, error) {
	var (
		err error
	)

	// ユーザー登録
	err = i.userRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update
// type Updateuser struct {
// 	Param *entity.User
// }

// type UpdateOutput struct {
// 	Ok bool
// }

func (i *UserInteractorImpl) Update(user *entity.User) (*entity.User, error) {
	var (
		err error
	)

	// ユーザー登録
	err = i.userRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Get
// type GetByIduser struct {
// 	Id uint
// }

// type GetByIdOutput struct {
// 	User *entity.User
// }

func (i *UserInteractorImpl) GetById(id uint) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	// ユーザー登録
	user, err = i.userRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// type SignInuser struct {
// 	Param *entity.SignInParam
// }

// type SignInOutput struct {
// 	User *entity.User
// }

func (i *UserInteractorImpl) SignIn(param *entity.SignInParam) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(param.Token)
	if err != nil {
		return user, err
	}

	fmt.Println("exported firebaseToken is:", param.Token)
	fmt.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	return user, nil

}

// type GetByFirebaseTokenuser struct {
// 	Token string
// }

// type GetByFirebaseTokenOutput struct {
// 	User *entity.User
// }

func (i *UserInteractorImpl) GetByFirebaseToken(token string) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(token)
	if err != nil {
		return user, err
	}

	fmt.Println("exported firebaseId is:", firebaseId)

	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	fmt.Println("exported user is:", user)

	return user, nil
}

// type GetAllOutput struct {
// 	Users []*entity.User
// }

func (i *UserInteractorImpl) GetAll() ([]*entity.User, error) {
	var (
		users []*entity.User
		err   error
	)

	users, err = i.userRepository.GetAll()
	if err != nil {
		return users, err
	}

	req := request.NewUserRequestImpl()
	req.DetectTextFromImage()

	return users, nil
}

// func gprcReq() {
// 	var (
// 		conn *grpc.ClientConn
// 		err  error
// 	)

// 	// Set up a connection to the server.
// 	conn, err = grpc.Dial(config.App.PythonDomain, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	defer conn.Close()
// 	c := pb.NewUserServiceClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.CreateUser(ctx, &pb.User{Name: "name", Email: "email", Password: "password"})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.GetMessage())
// }
