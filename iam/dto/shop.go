/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package dto

type AllShopResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	IsActive bool   `json:"is_active"`
	Role     string `json:"role"`
}

type CreateShopRequest struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug" validate:"required"`
}

// add other fields as needed
