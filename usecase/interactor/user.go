package interactor

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/hidenari-yuda/umerun-resume/domain/entity"
	"github.com/hidenari-yuda/umerun-resume/usecase"
	"gopkg.in/guregu/null.v4"
)

type UserInteractor interface {
	// Gest API
	SignUp(input SignUpInput) (output SignUpOutput, err error)
	SignIn(input SignInInput) (output SignInOutput, err error)
	GetByFirebaseToken(input GetByFirebaseTokenInput) (output GetByFirebaseTokenOutput, err error)
	DetectTextFromJobSeekerResume(input DetectTextFromJobSeekerResumeInput) (DetectTextFromJobSeekerResumeOutput, error)
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

type SignUpInput struct {
	Param *entity.SignUpParam
}

type SignUpOutput struct {
	Ok bool
}

func (i *UserInteractorImpl) SignUp(input SignUpInput) (output SignUpOutput, err error) {
	// ユーザー登録
	err = i.userRepository.SignUp(input.Param)
	if err != nil {
		return output, err
	}

	output.Ok = true

	return output, nil
}

type SignInInput struct {
	Param *entity.SignInParam
}

type SignInOutput struct {
	User *entity.User
}

func (i *UserInteractorImpl) SignIn(input SignInInput) (output SignInOutput, err error) {

	firebaseId, err := i.firebase.VerifyIDToken(input.Param.Token)
	if err != nil {
		return output, err
	}

	fmt.Println("exported firebaseToken is:", input.Param.Token)
	fmt.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	output.User, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		err = fmt.Errorf("failed to get user by firebaseId: %w", err)
		return output, err
	}

	return output, nil

}

type GetByFirebaseTokenInput struct {
	Token string
}

type GetByFirebaseTokenOutput struct {
	User *entity.User
}

func (i *UserInteractorImpl) GetByFirebaseToken(input GetByFirebaseTokenInput) (output GetByFirebaseTokenOutput, err error) {

	firebaseId, err := i.firebase.VerifyIDToken(input.Token)
	if err != nil {
		return output, err
	}

	fmt.Println("exported firebaseId is:", firebaseId)

	output.User, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return output, err
	}

	fmt.Println("exported user is:", output.User)

	return output, nil
}

type DetectTextFromJobSeekerResumeInput struct {
	FilePath string
}

type DetectTextFromJobSeekerResumeOutput struct {
	JobSeeker *entity.JobSeeker
}

func (i *UserInteractorImpl) DetectTextFromJobSeekerResume(input DetectTextFromJobSeekerResumeInput) (output DetectTextFromJobSeekerResumeOutput, err error) {
	// io.Writer
	w := io.Writer(os.Stdout)

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return output, err
	}

	file := ".public/images/townwork_template.webp"

	f, err := os.Open(file)
	if err != nil {
		return output, err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return output, err
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return output, err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
		return output, err
	}

	var jobSeeker *entity.JobSeeker

	for i, _ := range annotations {
		fmt.Fprintf(w, "%q\n", annotations[i].Description)

		if Contains(NameCheckerList, annotations[i].Description) && !(Contains(NameCheckerList, annotations[i+1].Description)) {
			jobSeeker.Name = annotations[i+1].Description
		}

		for _, birthdayChecker := range BirthdayCheckerList {
			if annotations[i].Description == birthdayChecker {
				jobSeeker.Birthday = annotations[i+1].Description
			}
		}

		if annotations[i].Description == "性別" {
			if annotations[i+1].Description == "男" {
				jobSeeker.Gender = null.NewInt(0, true)
			} else if annotations[i+1].Description == "女" {
				jobSeeker.Gender = null.NewInt(1, true)
			}
		}

	}

	fmt.Println("detected text is:", jobSeeker)

	output.JobSeeker = jobSeeker

	return output, nil
}

var NameCheckerList []string = []string{
	"氏名",
	"氏",
	"名",
	"名前",
	"お名前",
	"ご氏名",
}

var FuriganaCheckerList []string = []string{
	"フリガナ",
	"ふりがな",
}

var BirthdayCheckerList []string = []string{
	"生年月日",
	"誕生日",
	"生年月",
	"生年",
	"生月日",
	"生月",
	"生日",
}

func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			// 型変換可能か確認する
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			// 型変換する
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			// 等価判定をする
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}

func Includes(list, checkList []string) string {
	for _, v := range list {
		for _, c := range checkList {
			if v == c {
				return v
			}
		}
	}
	return ""
}
