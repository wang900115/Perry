package main

import (
	"github.com/wang900115/Perry/config"
	"github.com/wang900115/Perry/internal/adapter/controller"
	"github.com/wang900115/Perry/internal/adapter/middleware"
	"github.com/wang900115/Perry/internal/adapter/middleware/cors"
	"github.com/wang900115/Perry/internal/adapter/middleware/jwt"
	"github.com/wang900115/Perry/internal/adapter/middleware/ratelimiter"
	secureheader "github.com/wang900115/Perry/internal/adapter/middleware/secure-header"
	responser "github.com/wang900115/Perry/internal/adapter/response"
	"github.com/wang900115/Perry/internal/adapter/router"
	"github.com/wang900115/Perry/internal/application/usecase"
	gormimplement "github.com/wang900115/Perry/internal/implement/gorm"
	redisimplement "github.com/wang900115/Perry/internal/implement/redis"
	initializeCache "github.com/wang900115/Perry/internal/init/cache"
	initializeDB "github.com/wang900115/Perry/internal/init/database"
	initializeServer "github.com/wang900115/Perry/internal/init/server"
)

func main() {
	conf := config.NewConfig()

	redisPool := initializeCache.NewRedisPool(initializeCache.NewRedisOption(conf))
	mysql := initializeDB.NewMySQL(initializeDB.NewMysqlOption(conf))
	// !TODO 封裝Response
	response := responser.Response

	userRepo := gormimplement.NewUserImplement(mysql)
	todoRepo := gormimplement.NewToDoImplement(mysql)

	sessionRepo := redisimplement.NewSessionImplement(redisPool)
	// !TODO 設定issuer , secret , expiration
	tokenRepo := redisimplement.NewTokenImplement(redisPool)

	todoUsecase := usecase.NewToDoUsecase(&todoRepo)
	userUsecase := usecase.NewUserUsecase(&userRepo, &tokenRepo, &sessionRepo)

	todoController := controller.NewToDoController(todoUsecase, response)
	userController := controller.NewUserController(userUsecase, response)
	// !TODO 設定允許來源
	corsMiddleware := cors.NewCORS(response, allowOrigins)
	jwtMiddleware := jwt.NewJWT(response, &tokenRepo)
	secureMiddleware := secureheader.NewSecureHeader()
	// !TODO 設定每秒幾次流量
	redisRateLimiter := ratelimiter.NewRateLimiter(response, *redisPool, limitPerSecond)

	todoRoute := router.NewToDoRouter(todoController)
	userRoute := router.NewUserRouter(userController, jwtMiddleware)

	server := initializeServer.NewApp(
		[]router.IRoute{
			todoRoute,
			userRoute,
		},
		[]middleware.IMiddleware{
			corsMiddleware,
			secureMiddleware,
			redisRateLimiter,
		},
	)

	initializeServer.Run(server, initializeServer.NewServerOption(conf))
}
