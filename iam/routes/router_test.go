/*
 * ****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ALL RIGHTS RESERVED
 * ****************************************************************************************
 */

package routes

import (
	"github.com/halng/anyshop/constants"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSHeaders(t *testing.T) {
	router := Routes()

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
