package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/authentication-service/cmd/inithandler"

	"github.com/authentication-service/auth"
)

func main() {
	config, err := inithandler.LoadConfig(".")
	if err != nil {
		log.Panic("failed to load config", err.Error())
	}

	logger := inithandler.InitLogger()
	db := inithandler.InitDB(logger)
	if db == nil {
		log.Fatal("failed to connect db")
	}	
	repo := inithandler.InitRepository(db)
	bl := auth.NewBL(logger, repo)
	endpoints := auth.NewEndpoints(logger, bl)
	handlers := auth.NewHTTPHandler(logger, endpoints)
	var errs = make(chan error)
	go func() {
		logger.Log("[debug]", "starting the server on port", config.WebPort)
		errs <- http.ListenAndServe(fmt.Sprintf(":%d", config.WebPort), handlers)
	}()

	logger.Log("terminated", <-errs)
}
