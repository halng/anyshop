/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/services"
)

// Register a new user
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body dto.RegisterRequest true "Register Request"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /register [post]
func Register(c *gin.Context) {
	services.Register(c)
}

// Login verify user credentials and return uuid pair with token saved in redis
// @Summary Login user
// @Description Verify user credentials and return uuid pair with token saved in redis
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /login [post]
func Login(c *gin.Context) {
	services.Login(c)
}

// Activate a user account by giving token and expired time
// @Summary Activate user account
// @Description Activate a user account by giving token and expired time
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /activate [post]
func Activate(context *gin.Context) {
	services.Activate(context)
}
