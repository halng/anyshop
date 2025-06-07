/*
 * ****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ALL RIGHTS RESERVED
 * ****************************************************************************************
 */

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/controller"
	"github.com/halng/anyshop/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {
	router := gin.Default()

	// set up cors origin
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", constants.ApiTokenRequestHeader, constants.ApiUserIdRequestHeader},
		AllowCredentials: true,
	}))

	// auth routes
	authGroup := router.Group("/api/v1/iam")
	authGroup.POST("/login", controller.Login)
	authGroup.POST("/register", controller.Register)
	authGroup.POST("/activate", controller.Activate)

	// shop routes
	shopGroup := router.Group("/api/v1/iam/shops")
	shopGroup.GET("", controller.GetAllShops)
	shopGroup.POST("", controller.CreateShop)

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1/iam"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
