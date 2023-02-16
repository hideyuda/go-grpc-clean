package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/utils"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/infra/router"

	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
)

func init() {

	time.Local = utils.Tokyo

	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file")
	}

	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			panic("Failed to load .env file")
		}
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	utils.LoggingSettings(config.App.LogFilePath)
	log.Println(fmt.Sprint("Server is running on port: ", config.App.Port))

	// ctx := context.Background()

	switch cfg.App.Service {
	case "api":
		// 一旦 apiコンテナを立ち上げる時にマイグレーションする
		db := database.NewDb()
		err := db.MigrateUp(".migrations")
		if err != nil {
			fmt.Println(err)
		}
		// cache := driver.NewRedisCacheImpl(cfg.Redis)
		if cfg.App.Env == "local" {
			firebase := driver.NewFirebaseImpl()
			fmt.Println("getTestUserToken:", uuid.New().String())
			getTestUserToken(firebase, uuid.New().String())
		}
		r := router.NewRouter(cfg)

		// // エラーハンドラー（dev or prdのみSlack通知）
		if cfg.App.Env != "local" {
			r.Engine.HTTPErrorHandler = customHTTPErrorHandler
		}

		// ルーティング
		r.SetUp().Start()

		// case "batch":
		// 	batch.NewBatch(cfg).Start()
	}
}

// func register(ctx context.Context, s *grpc.Server) {
// 	// c := infra.NewFirestoreClient(ctx)
// 	// repo := infra.NewMessageRepositoryImpl(c)
// 	repo := infra.NewLocalMessageRepositoryImpl()
// 	createMessageService := usecase.NewCreateMessageService(repo)
// 	getMessageService := usecase.NewGetMessageStreamService(repo)
// 	pb.RegisterChatServiceServer(s, NewServer(*createMessageService, *getMessageService))
// }

func getTestUserToken(fb usecase.Firebase, uuid string) {
	customToken, _ := fb.GetCustomToken(uuid)
	idToken, err := fb.GetIDToken(customToken)
	if err != nil {
		panic(err)
	}
	fmt.Println("test token is :", idToken)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		cfg, _        = config.New()
		code, message = entity.ErrorInfo(err)
		statusCode    = strconv.Itoa(code)
		path          = c.Path()
		method        = c.Request().Method
		errText       = err.Error()
	)

	fmt.Println(err)

	te := "*開発環境 Error*\n" +
		">>>status: " + message + "（" + statusCode + "）" + "\n" +
		"method: " + method + "\n" +
		"uri: " + path + "\n" +
		"error: `" + errText + "` \n"

	// アクセストークンを使用してクライアントを生成する
	// https://api.slack.com/apps からトークン取得
	// 参考: https://risaki-masa.com/how-to-get-api-token-in-slack/
	tkn := cfg.Slack.AccessToken
	chanelID := cfg.Slack.ChanelID
	s := slack.New(tkn)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err = s.PostMessage(chanelID, slack.MsgOptionBlocks(
		&slack.SectionBlock{
			Type: slack.MBTSection,
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: te,
			},
		},
	))
	if err != nil {
		fmt.Println(err)
	}

	c.Logger().Error(err)
}
