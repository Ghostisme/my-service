package main

import (
	"log"
	"my-service/global"
	"my-service/internal/model"
	"my-service/internal/routers"
	"my-service/pkg/logger"
	"my-service/pkg/setting"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

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
	err = setting.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	} else {
		global.ServerSettings.ReadTimeout *= time.Second
		global.ServerSettings.WriteTimeout *= time.Second
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

	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSettings err: %v", err)
	}
	// err = setupDBEngine()
	// if err != nil {
	// 	log.Fatalf("init.setupDBEngine err: %v", err)
	// }
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