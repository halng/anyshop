/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package main

import (
	"fmt"
	"github.com/halng/anyshop/routes"
	"os"

	"github.com/halng/anyshop/db"
	"github.com/halng/anyshop/kafka"
	"github.com/halng/anyshop/logging"
	"github.com/halng/anyshop/models"
	"github.com/joho/godotenv"
)

func main() {
	logging.InitLogging()

	// connect database
	db.ConnectDB()
	models.Initialize()

	var err error

	_ = godotenv.Load(".env")

	// init kafka server
	bootstrapServer := os.Getenv("KAFKA_HOST")
	err = kafka.InitializeKafkaProducer(bootstrapServer)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	engine := routes.Routes()

	err = engine.Run(":" + port)
	logging.LOGGER.Info(fmt.Sprintf("Starting web service on port %s", port))
	if err != nil {
		return
	}
}
