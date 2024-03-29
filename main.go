package main

import (
	"context"
	"log"
	"my-service/global"
	"my-service/internal/model"
	"my-service/internal/routers"
	"my-service/pkg/logger"
	"my-service/pkg/setting"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ctx = context.Background()

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DBSettings)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSettings)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &global.RedisSettings)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	} else {
		global.ServerSettings.ReadTimeout *= time.Second
		global.ServerSettings.WriteTimeout *= time.Second
	}
	err = setting.ReadSection("JWT", &global.JWTSettings)
	if err != nil {
		return err
	} else {
		global.JWTSettings.Expire *= time.Second
		global.JWTSettings.Issuer = "myspace"
		global.JWTSettings.AppSecret = "myspace-us"
	}
	return nil
}

func setupRedisEngine() error {
	var err error
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.RedisSettings.RedisAddr,
		Password: global.RedisSettings.RedisPassword,
		DB:       global.RedisSettings.RedisDB,
	})
	_, err = global.RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DBSettings)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + global.AppSettings.LogFileName + global.AppSettings.LogFileExt,
		MaxSize:   60,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	global.DaoLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + "dao" + global.AppSettings.LogFileExt,
		MaxSize:   60,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	global.ServiceLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + "service" + global.AppSettings.LogFileExt,
		MaxSize:   60,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	global.ModelLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + "model" + global.AppSettings.LogFileExt,
		MaxSize:   60,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	global.ApiLogger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSettings.LogSavePath + "/" + "api" + global.AppSettings.LogFileExt,
		MaxSize:   60,
		MaxAge:    30,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSettings err: %v", err)
	}
	// err = setupRedisEngine()
	// if err != nil {
	// 	log.Fatalf("init.setupRedisEngine err: %v", err)
	// }
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	//解析输入
	// setupFlag()
}

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(global.ServerSettings.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSettings.ReadTimeout,
		WriteTimeout:   global.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("Listening: %v", global.ServerSettings.HttpPort)
	// go func() {
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.Logger.Fatalf("s.ListenAndServe err: %v", err)
	}
	// }()
}
