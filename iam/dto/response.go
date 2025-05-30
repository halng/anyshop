/*
 * ****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ALL RIGHTS RESERVED
 * ****************************************************************************************
 */

package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/logging"
	"go.uber.org/zap"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Details interface{} `json:"details"`
}

// SuccessResponse sends a success JSON response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Code:   statusCode,
		Status: "SUCCESS",
		Data:   data,
	})
}

// ErrorResponse sends a structured error response
func ErrorResponse(c *gin.Context, statusCode int, message string, details interface{}) {
	logging.LOGGER.Error("Failed to handle request",
		zap.String("endpoint", c.Request.RequestURI),
		zap.String("method", c.Request.Method),
		zap.String("remote_address", c.Request.RemoteAddr),
		zap.Any("header", c.Request.Header),
		zap.Any("error", details),
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)

	c.AbortWithStatusJSON(statusCode, APIResponse{
		Code:    statusCode,
		Status:  "ERROR",
		Error:   message,
		Details: details,
	})

	//c.Abort()
}

func ForbiddenResponse(c *gin.Context, message string, traceData any) {
	ErrorResponse(c, http.StatusForbidden, message, traceData)
}

func UnauthorizedResponse(c *gin.Context, message string, traceData any) {
	ErrorResponse(c, http.StatusUnauthorized, message, traceData)
}

func BadRequestResponse(c *gin.Context, message string, traceData any) {
	ErrorResponse(c, http.StatusBadRequest, message, traceData)
}

func InternalServerErrorResponse(c *gin.Context, message string, traceData any) {
	ErrorResponse(c, http.StatusInternalServerError, message, traceData)
}

func NotFoundResponse(c *gin.Context, message string, traceData any) {
	ErrorResponse(c, http.StatusNotFound, message, traceData)
}
