/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 11:42:41
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 11:50:52
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"go-admin/global"
	"go-admin/internal/model"
	"go-admin/internal/routers"
	"go-admin/pkg/setting"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	port      string
	runMode   string
	config    string
	isVersion bool
)

func init() {
	err := initFlag()
	if err != nil {
		log.Fatalf("init.initFlag err: %v", err)
	}
	err = initSetting()
	if err != nil {
		log.Fatalf("init.initSetting err: %v", err)
	}
	err = initLogger()
	if err != nil {
		log.Fatalf("init.initLogger err:%v", err)
	}
	err = initDBEngine()
	if err != nil {
		log.Fatalf("init.initDBEngine err: %v", err)
	}
}

func main() {
	fmt.Println("hello world")
	// gin.SetMode(global.ServerSetting.RunMode)
	// routers.NewRouter()
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Server Runing in %s \n", "http://localhost:"+global.ServerSetting.HttpPort)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func initSetting() error {
	fmt.Println(config,"config---")

	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func initDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}
	global.DBEngine.Debug()

	return nil
}

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "dev", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()

	return nil
}

func initLogger() error {
	path := "storage/logs/app.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil || os.IsNotExist(err) {
		_, err = os.Create(path)
		return err
	}
	var log = logrus.New()
	log.SetReportCaller(true) //输出日志中添加文件名和方法信息
	log.Formatter = &logrus.JSONFormatter{}
	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.Info("init  logger  success2")
	global.Logger = log

	return err
}
