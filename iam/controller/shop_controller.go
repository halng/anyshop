/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/services"
)

// GetAllShops retrieves all shops from the database.
// @Summary GetAllShops retrieves all shops belonging to the current user
// @Tags shop
// @Accept json
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /all [post]
func GetAllShops(c *gin.Context) {
	services.GetAllShops(c)
}

// CreateShop creates a new shop for the current user.
// @Summary CreateShop creates a new shop for the current user
// @Tags shop
// @Accept json
// @Produce json
// @Param createShopRequest body dto.CreateShopRequest true "Create Shop Request"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
func CreateShop(c *gin.Context) {
	services.CreateShop(c)
}
