package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"restapi/internal/config"
	"restapi/internal/user"
	"restapi/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	//cfgMongo := cfg.MongoDB
	//mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
	//	cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	//if err != nil {
	//	fmt.Errorf("panic")
	//}
	//
	//storage := db.NewStorage(mongoDBClient, cfgMongo.Collection, logger)
	//
	//users, err := storage.FindAll(context.Background())
	//fmt.Println(users)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("create unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Info(fmt.Sprintf("server is listening unix socket: %s", socketPath))
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen(
			"tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Info(fmt.Sprintf("server is listening %s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
