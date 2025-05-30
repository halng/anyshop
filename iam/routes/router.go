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
	"github.com/halng/anyshop/docs"
	"github.com/halng/anyshop/handlers"
	"github.com/halng/anyshop/middleware"
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

	router.Use(middleware.ErrorHandler())

	// auth routes
	authGroup := router.Group("/api/v1/auth")
	authGroup.POST("/login", handlers.Login)
	authGroup.POST("/register", handlers.Register)
	authGroup.POST("/activate", handlers.Activate)

	//userGroup := router.Group("/api/v1/user")
	//// user routes
	//userGroup.POST("/login", handlers.Login)
	//userGroup.POST("/create-staff", middleware.ValidateRequest, handlers.CreateStaff)
	//userGroup.POST("/register", handlers.Register)
	//userGroup.GET("/validate", handlers.Validate)
	//userGroup.POST("/activate", handlers.Activate)
	//
	//// shop routes
	//shopGroup := router.Group("/api/v1/shop")
	//shopGroup.POST("", middleware.ValidateRequest, handlers.CreateShop)
	//shopGroup.PUT("", middleware.ValidateRequest, handlers.UpdateShop)

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1/auth"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
