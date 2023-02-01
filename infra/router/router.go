package router

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/driver"
	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type ServiceServer struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedChatServiceServer
	Db       *database.Db
	Firebase *driver.FirebaseImpl
}

type Router struct {
	cfg    config.Config
	Engine *echo.Echo
}

func NewRouter(cfg config.Config) *Router {
	return &Router{
		cfg:    cfg,
		Engine: echo.New(),
	}
}

func (r *Router) SetUp() *Router {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// register(ctx, s)

	// https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/server#%5B%E3%82%B3%E3%83%A9%E3%83%A0%5D%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%83%AA%E3%83%95%E3%83%AC%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%81%A8%E3%81%AF%EF%BC%9F
	reflection.Register(s)
	pb.RegisterUserServiceServer(s, &ServiceServer{})
	pb.RegisterChatServiceServer(s, &ServiceServer{})

	go func() {
		fmt.Printf("start gRPC server, port: %d", config.App.Port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

	return r

	// var (
	// db       = database.NewDB(r.cfg.Db, true)
	// firebase = driver.NewFirebaseImpl(r.cfg.Firebase)
	// basicAuth = utility.NewBasicAuth(r.cfg)
	// )

	// r.Engine.HidePort = true
	// r.Engine.HideBanner = true
	// r.Engine.Use(middleware.Recover())
	// // TODO: Webクライアントのドメインが決まったら設定する 👆の`r.Engine.Use(middleware.CORS())`は消す
	// // r.Engine.Use(middleware.CORSWithConfig((middleware.CORSConfig{
	// // AllowOrigins: r.cfg.App.CorsDomains,
	// // 	AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderOrigin, echo.HeaderAccessControlAllowOrigin},
	// // 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	// // })))
	// r.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	// }))
	// r.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if strings.Contains(c.Request().URL.Path, "healthz") {
	// 			return true
	// 		} else {
	// 			return false
	// 		}
	// 	},
	// }))
	// r.Engine.HidePort = true
	// r.Engine.HideBanner = true
	// r.Engine.Use(middleware.Recover())

	// var origins []string

	// if r.cfg.App.Env == "local" {
	// 	origins = []string{
	// 		"http://localhost:9090",
	// 		"http://localhost:3000",
	// 		"http://localhost:3001",
	// 		"http://localhost:3002",
	// 	}
	// } else if r.cfg.App.Env == "dev" {
	// 	origins = r.cfg.App.CorsDomains
	// } else if r.cfg.App.Env == "prd" {
	// 	origins = r.cfg.App.CorsDomains
	// }

	// fmt.Println("------------")
	// fmt.Println(r.cfg.App.Env)
	// fmt.Println(origins)
	// fmt.Println("------------")

	// r.Engine.Use(middleware.CORSWithConfig((middleware.CORSConfig{
	// 	AllowOrigins: origins,
	// 	AllowHeaders: []string{
	// 		echo.HeaderAuthorization,
	// 		echo.HeaderAccessControlAllowHeaders,
	// 		echo.HeaderContentType,
	// 		echo.HeaderOrigin,
	// 		echo.HeaderAccessControlAllowOrigin,
	// 		"FirebaseAuthorization",
	// 	},
	// 	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	// })))

	// r.Engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: func(c echo.Context) bool {
	// 		if strings.Contains(c.Request().URL.Path, "healthz") {
	// 			return true
	// 		} else {
	// 			return false
	// 		}
	// 	},
	// }))

	// api := r.Engine.Group("")
	// {
	// 	api.GET("/healthz", func(c echo.Context) error {
	// 		return c.NoContent(http.StatusOK)
	// 	})

	// 	api.GET("/*", func(c echo.Context) error {
	// 		return c.NoContent(http.StatusNotFound)
	// 	})

	// 	api.POST("/*", func(c echo.Context) error {
	// 		return c.NoContent(http.StatusNotFound)
	// 	})

	// 	api.PUT("/*", func(c echo.Context) error {
	// 		return c.NoContent(http.StatusNotFound)
	// 	})
	// }

	// /****************************************************************************************/
	// /// No Auth API
	// //

	// noAuthAPI := api.Group("api")
	// {
	// 	noAuthAPI.GET("/healthz", func(c echo.Context) error {
	// 		return c.NoContent(http.StatusOK)
	// 	})

	// 	// ユーザーの新規登録
	// 	noAuthAPI.POST("/signup", userRoutes.SignUp(db, firebase))

	// 	// ユーザーのログイン
	// 	noAuthAPI.PUT("/signin", userRoutes.SignIn(db, firebase))

	// }

	// /****************************************************************************************/
	// /// UserAPI
	// //
	// // userAPI := noAuthAPI.Group("/user")
	// {
	// 	// ユーザーのログイン

	// }
}

func (r *Router) Start() {
	r.Engine.Start(fmt.Sprintf(":%d", config.App.Port))
}
