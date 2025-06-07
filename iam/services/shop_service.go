/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/dto"
	"github.com/halng/anyshop/logging"
	"github.com/halng/anyshop/models"
	"github.com/halng/anyshop/utils"
	"go.uber.org/zap"
	"net/http"
)

// GetAllShops retrieves all shops from the database.
// Extract user id from token and fetch shops associated with that user.
// shop name/id - user role
func GetAllShops(c *gin.Context) {

	userID := GetUserIdFromRequest(c)
	if utils.IsNullOrEmpty(userID) {
		dto.UnauthorizedResponse(c, constants.Unauthorized, nil)
		return
	}

	// Fetch shops associated with the user ID
	shops, err := models.GetAllShopUserByID(userID)
	if err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return
	}

	var results []dto.AllShopResponse

	for _, shop := range shops {
		results = append(results, dto.AllShopResponse{
			Id:       shop.Shop.ID.String(),
			Name:     shop.Shop.Name,
			Slug:     shop.Shop.Slug,
			IsActive: shop.Shop.IsActive,
			Role:     shop.Role.Name,
		})
	}

	dto.SuccessResponse(c, http.StatusOK, results)
}

func CreateShop(c *gin.Context) {
	userID := GetUserIdFromRequest(c)
	if utils.IsNullOrEmpty(userID) {
		dto.UnauthorizedResponse(c, constants.Unauthorized, nil)
		return
	}

	var shopInput dto.CreateShopRequest
	if err := c.ShouldBindJSON(&shopInput); err != nil {
		dto.BadRequestResponse(c, constants.MessageErrorBindJson, err)
		return
	}

	if _, errors := utils.ValidateInput(shopInput); errors != nil {
		dto.BadRequestResponse(c, constants.MissingParams, errors.Error())
		return
	}

	exists, err := models.ExistByNameOrSlug(shopInput.Name, shopInput.Slug)
	if err != nil {
		logging.LOGGER.Error("Failed to check if shop exists", zap.Error(err))
		dto.InternalServerErrorResponse(c, constants.InternalServerError, nil)
		return
	}

	if exists {
		dto.BadRequestResponse(c, fmt.Sprintf(constants.ShopAlreadyExists, shopInput.Name, shopInput.Slug), nil)
		return
	}

	adminRole, err := models.GetRoleByName("ADMIN")
	if err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return
	}

	newShop := models.Shop{
		Name:     shopInput.Name,
		Slug:     shopInput.Slug,
		IsActive: true,
		CreateBy: userID,
		UpdateBy: userID,
	}

	if err := newShop.Save(); err != nil {
		logging.LOGGER.Error("Failed to create new shop", zap.Any("shop", newShop), zap.Error(err))
		dto.InternalServerErrorResponse(c, constants.InternalServerError, nil)
		return
	}

	newShopUser := models.ShopUser{
		ShopID: newShop.ID,
		RoleID: adminRole.Id,
		UserID: uuid.MustParse(userID),
	}

	if err := newShopUser.Save(); err != nil {
		logging.LOGGER.Error("Failed to create shop user", zap.Any("shopUser", newShopUser), zap.Error(err))
		dto.InternalServerErrorResponse(c, constants.InternalServerError, nil)
		return
	}

	dto.SuccessResponse(c, http.StatusCreated, nil)
}
