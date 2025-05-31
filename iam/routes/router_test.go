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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock handlers
func MockLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mock login successful"})
}

func MockRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mock register successful"})
}

func MockActivate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mock activate successful"})
}

func MockRoutes() *gin.Engine {
	router := gin.Default()

	// Set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", constants.ApiTokenRequestHeader, constants.ApiUserIdRequestHeader},
		AllowCredentials: true,
	}))

	// Mock auth routes
	authGroup := router.Group("/api/v1/auth")
	authGroup.POST("/login", MockLogin)
	authGroup.POST("/register", MockRegister)
	authGroup.POST("/activate", MockActivate)

	// Mock swagger route
	router.GET("/swagger/*any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Mock swagger route"})
	})

	return router
}

func TestRoutes(t *testing.T) {
	router := MockRoutes()

	tests := []struct {
		name       string
		method     string
		target     string
		body       string
		statusCode int
	}{
		{
			name:       "Login route works",
			method:     http.MethodPost,
			target:     "/api/v1/auth/login",
			body:       `{}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Register route works",
			method:     http.MethodPost,
			target:     "/api/v1/auth/register",
			body:       `{}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Activate route works",
			method:     http.MethodPost,
			target:     "/api/v1/auth/activate",
			body:       `{}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Swagger route works",
			method:     http.MethodGet,
			target:     "/swagger/index.html",
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.target, strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, tt.target, nil)
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}

func TestCORSHeaders(t *testing.T) {
	router := MockRoutes()

	req := httptest.NewRequest(http.MethodOptions, "/api/v1/auth/login", nil)
	req.Header.Set("Origin", "http://localhost")
	req.Header.Set("Access-Control-Request-Method", "POST")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, strings.ToUpper(w.Header().Get("Access-Control-Allow-Headers")), constants.ApiTokenRequestHeader)
	assert.Contains(t, strings.ToUpper(w.Header().Get("Access-Control-Allow-Headers")), constants.ApiUserIdRequestHeader)
}
