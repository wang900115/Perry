package init

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wang900115/Perry/internal/adapter/middleware"
	"github.com/wang900115/Perry/internal/adapter/router"
)

type serverOption struct {
	RunMode  string
	HTTPPort string

	UseTLS      bool
	TLSCertFile string
	TLSKeyFile  string

	HeaderReadTimeout time.Duration // header 最大讀取時間
	BodyReadTimeout   time.Duration // body 最大讀取時間
	WriteTimeout      time.Duration // 整個response 回應的最大時間
	IdleTimeout       time.Duration // keep - alive 最大時間
	CancelTimeout     time.Duration // 清理資源用途(背景JOB)

	MaxHeaderBytes int // 最大標頭體積
	MaxBodyBytes   int // 最大傳輸體積
}

func defaultOption() serverOption {
	return serverOption{
		RunMode:  gin.DebugMode,
		HTTPPort: "8080",

		UseTLS:      false,
		TLSCertFile: "",
		TLSKeyFile:  "",

		HeaderReadTimeout: 5 * time.Second,

		BodyReadTimeout: 5 * time.Second,

		WriteTimeout:  5 * time.Second,
		IdleTimeout:   5 * time.Second,
		CancelTimeout: 5 * time.Second,

		MaxHeaderBytes: 64000,
		MaxBodyBytes:   10000000,
	}
}

func NewServerOption(conf *viper.Viper) serverOption {
	defaultOptions := defaultOption()
	/*伺服器設定*/
	if conf.IsSet("server.run_mode") {
		defaultOptions.RunMode = conf.GetString("server.run_mode")
	}
	if conf.IsSet("server.http_port") {
		defaultOptions.HTTPPort = conf.GetString("server.http_port")
	}
	/*響應時間設定*/
	if conf.IsSet("server.header_read_timeout") {
		defaultOptions.HeaderReadTimeout = conf.GetDuration("server.header_read_timeout")
	}
	if conf.IsSet("server.body_read_timeout") {
		defaultOptions.BodyReadTimeout = conf.GetDuration("server.body_read_timeout")
	}
	if conf.IsSet("server.write_timeout") {
		defaultOptions.WriteTimeout = conf.GetDuration("server.write_timeout")
	}
	if conf.IsSet("server.idle_timeout") {
		defaultOptions.IdleTimeout = conf.GetDuration("server.idle_timeout")
	}
	if conf.IsSet("server.cancel_timeout") {
		defaultOptions.CancelTimeout = conf.GetDuration("server.cancel_timeout")
	}
	/*傳輸體積設定*/
	if conf.IsSet("server.max_header_bytes") {
		defaultOptions.MaxHeaderBytes = conf.GetInt("server.max_header_bytes")
	}
	if conf.IsSet("server.max_body_bytes") {
		defaultOptions.MaxBodyBytes = conf.GetInt("server.max_body_bytes")
	}

	return defaultOptions
}

type App struct {
	routes      []router.IRoute
	middlewares []middleware.IMiddleware
}

func NewApp(routes []router.IRoute, middlewares []middleware.IMiddleware) *App {
	return &App{routes: routes, middlewares: middlewares}
}

func Run(a *App, option serverOption) {
	// 初始化設定以及路由註冊
	gin.SetMode(option.RunMode)
	router := gin.Default()
	for _, middleware := range a.middlewares {
		router.Use(middleware.Middleware)
	}
	for _, route := range a.routes {
		route.SetUp(router.Group("/api"))
	}
	if option.RunMode == "debug" {
		pprof.Register(router)
	}
	srv := &http.Server{
		Addr:    ":" + option.HTTPPort,
		Handler: router,

		ReadHeaderTimeout: option.HeaderReadTimeout,
		WriteTimeout:      option.WriteTimeout,
		IdleTimeout:       option.IdleTimeout,

		MaxHeaderBytes: option.MaxHeaderBytes,
	}

	// 伺服器啟動
	go func() {
		var err error
		if option.UseTLS {
			cert, err := tls.LoadX509KeyPair(option.TLSCertFile, option.TLSKeyFile)
			if err != nil {
				panic(err)
			}
			srv.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}}

			err = srv.ListenAndServeTLS("", "")
		} else {
			err = srv.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// 伺服器關機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⏻ Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server Shutdown Error: %v", err)
	}
}
