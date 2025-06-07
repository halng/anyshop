/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/logging"
	"github.com/halng/anyshop/utils"
	"go.uber.org/zap"
	"strings"
)

func GetUserIdFromRequest(c *gin.Context) string {
	bearerToken := c.GetHeader(constants.Authorization)
	if utils.IsNullOrEmpty(bearerToken) {
		return ""
	}

	token, ok := strings.CutPrefix(bearerToken, "Bearer ")
	if !ok || utils.IsNullOrEmpty(token) {
		return ""
	}

	userID, _, err := utils.ExtractJWT(token)
	if err != nil {
		logging.LOGGER.Error("Failed to extract user ID from token", zap.Any("error", err))
		return ""
	}

	return userID
}
