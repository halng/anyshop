/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package dto

type ResDTO struct {
	StatusCode int      `json:"status_code"`
	Status     string   `json:"status"`
	Data       any      `json:"data"`
	Error      ErrorDTO `json:"error"`
}

type ErrorDTO struct {
	Message any `json:"msg"`
}

type ActiveNewUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Token          string `json:"token"`
	ExpiredTime    string `json:"expired_time"`
	ActivationLink string `json:"activation_link"`
}

type ActiveNewUserMsg struct {
	Action string        `json:"action"`
	Data   ActiveNewUser `json:"data"`
}
