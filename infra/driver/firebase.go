package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type FirebaseImpl struct {
	auth      *auth.Client
	firestore *firestore.Client
	webAPIKey string
}

func NewFirebaseImpl() usecase.Firebase {
	ctx := context.Background()
	opts := option.WithCredentialsFile(config.Firebase.JsonFilePath)

	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		panic(err.Error())
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		panic(err.Error())
	}

	firestore, err := app.Firestore(ctx)
	if err != nil {
		panic(err.Error())
	}

	return &FirebaseImpl{
		auth:      auth,
		firestore: firestore,
		webAPIKey: config.Firebase.WebApiKey,
	}
}

func (d *FirebaseImpl) VerifyIDToken(idToken string) (string, error) {
	token, err := d.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "ID token has expired"):
			return "", entity.ErrFirebaseExpiredToken
		case strings.Contains(err.Error(), "failed to verify token signature"): // token文字列が不正
			return "", entity.ErrFirebaseFailedToVerify
		case strings.Contains(err.Error(), "ID token has invalid"): // tokenが不正（異なるfirebaseプロジェクトなど）
			return "", entity.ErrFirebaseInvalidToken
		case strings.Contains(err.Error(), "ID token issued at future timestamp"):
			return "", entity.ErrFirebaseFutureIssued
		default:
			fmt.Println(err)
			return "", entity.ErrServerError
		}
	}
	return token.UID, nil
}

func (d *FirebaseImpl) GetCustomToken(uid string) (string, error) {
	token, err := d.auth.CustomToken(context.Background(), uid)
	if err != nil {
		return "", entity.ErrServerError
	}
	return token, nil
}

func (d *FirebaseImpl) GetIDToken(token string) (string, error) {

	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=" + d.webAPIKey
	input := &struct {
		Token             string `json:"token"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}{
		Token:             token,
		ReturnSecureToken: true,
	}
	params, err := json.Marshal(input)
	if err != nil {
		return "", errors.Wrap(entity.ErrServerError, err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return "", errors.Wrap(entity.ErrServerError, err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(entity.ErrServerError, err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(entity.ErrServerError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.Wrap(entity.ErrServerError, "status code is wrong")
	}

	output := &struct {
		IDToken string `json:"idToken"`
	}{}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", errors.Wrap(entity.ErrServerError, err.Error())
	}

	return output.IDToken, nil
}

func (d *FirebaseImpl) GetPhoneNumber(uid string) (string, error) {
	record, err := d.auth.GetUser(context.Background(), uid)
	if err != nil {
		return "", fmt.Errorf("%s:%w", err.Error(), entity.ErrServerError)
	}
	return record.PhoneNumber, nil
}

func (d *FirebaseImpl) Set(doc string, data map[string]interface{}) error {
	_, err := d.firestore.
		Collection("pitacoin").
		Doc(doc).
		Set(context.Background(), data, firestore.MergeAll)
	if err != nil {
		return errors.Wrap(entity.ErrServerError, err.Error())
	}
	return nil
}

// https://qiita.com/nisitanisubaru/items/3ff4e0b08b20700f408c
func (d *FirebaseImpl) CreateUser(email, password string) (string, error) {
	user := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		Disabled(false)

	u, err := d.auth.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Println("email: ", email)
		return "", fmt.Errorf("%s:%w", strings.Split(err.Error(), ""), entity.ErrFirebaseEmailExists)
	}

	return u.UID, err
}

// https://qiita.com/nisitanisubaru/items/3ff4e0b08b20700f408c
func (d *FirebaseImpl) UpdateEmail(email, uid string) error {
	user := (&auth.UserToUpdate{}).
		Email(email).
		EmailVerified(true)

	_, err := d.auth.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return fmt.Errorf("%s:%w", strings.Split(err.Error(), ""), entity.ErrFirebaseEmailExists)
	}

	return err
}

func (d *FirebaseImpl) UpdatePassword(password, uid string) error {
	user := (&auth.UserToUpdate{}).Password(password)

	_, err := d.auth.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return fmt.Errorf("%s:%w", strings.Split(err.Error(), ""), entity.ErrFirebaseEmailExists)
	}

	return err
}

// create chat
func (d *FirebaseImpl) CreateChat(ctx context.Context, chat *pb.Chat) error {
	_, err := d.firestore.Collection("chat").Doc(fmt.Sprint(chat.Id)).Set(ctx, chat)
	if err != nil {
		return fmt.Errorf("failed to chatRepositoryImpl.CreateChat, firestore.Collection: %s", err)
	}
	return nil
}

// update chat
func (d *FirebaseImpl) UpdateChat(ctx context.Context, chat *pb.Chat) error {
	_, err := d.firestore.Collection("chat").Doc(fmt.Sprint(chat.Id)).Set(ctx, chat)
	if err != nil {
		return fmt.Errorf("failed to chatRepositoryImpl.UpdateChat, firestore.Collection: %s", err)
	}
	return nil
}

// get chat
func (d *FirebaseImpl) GetChatById(ctx context.Context, id int64) (*pb.Chat, error) {
	chat := pb.Chat{}
	doc, err := d.firestore.Collection("chat").Doc(fmt.Sprint(id)).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to chatRepositoryImpl.GetChat, firestore.Collection: %s", err)
	}
	if err := doc.DataTo(&chat); err != nil {
		return nil, fmt.Errorf("failed to chatRepositoryImpl.GetChat, doc.DataTo: %s", err)
	}
	return &chat, nil
}

// Listen はchatコレクションのリアルタイムアップデートを確認する処理です
//  https://firebase.google.com/docs/firestore/query-data/listen#view_changes_between_snapshots
func (d *FirebaseImpl) GetChatStream(ctx context.Context, stream chan<- pb.Chat) error {
	chat := pb.Chat{}

	snapIter := d.firestore.Collection("chat").Snapshots(ctx)
	defer snapIter.Stop()

	for {
		snap, err := snapIter.Next()
		if err != nil {
			return fmt.Errorf("failed to chatRepositoryImpl.Listen, snapIter.Next: %s", err)
		}
		log.Printf("change size: %d\n", len(snap.Changes))
		for _, diff := range snap.Changes {
			switch diff.Kind {
			case firestore.DocumentAdded:
				if err := diff.Doc.DataTo(&chat); err != nil {
					return fmt.Errorf("failed to chatRepositoryImpl.Listen: %s", err)
				}
			}
			chat.Uuid = diff.Doc.Ref.ID

			select {
			case <-ctx.Done():
				return nil
			default:
				stream <- chat
			}
		}
	}
}
