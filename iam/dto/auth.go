/*
 * ****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ALL RIGHTS RESERVED
 * ****************************************************************************************
 */

package dto

type RegisterRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	ApiToken string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}
