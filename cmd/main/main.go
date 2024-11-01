package main

import (
	"github.com/CrudOperationUsingGin/pkg/logger"
	"github.com/CrudOperationUsingGin/pkg/router"

	"github.com/CrudOperationUsingGin/pkg/config"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")

	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
