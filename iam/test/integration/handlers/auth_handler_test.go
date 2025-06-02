/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/db"
	"github.com/halng/anyshop/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"testing"

	handlers2 "github.com/halng/anyshop/handlers"
	"github.com/halng/anyshop/test/integration"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	path := "/api/v1/auth/register"
	router := integration.SetUpRouter()

	router.POST(path, handlers2.Register)

	/**
	Test case 1: Register with invalid JSON data.
	Test case 2: Register with invalid data that does not meet requirements.
	Test case 3: Register with valid data.
	Test case 4: Register with an account that already exists.
	*/

	t.Run("Register: when data is invalid to process", func(t *testing.T) {

		invalidCases := `{"email": "change@gmail.com, "password": "12345678", "confirm_password": "12345678", "username": "hellothere" }`

		code, res := integration.ServeRequest(router, "POST", path, invalidCases)
		assert.Equal(t, code, http.StatusBadRequest)
		assert.Equal(t, res, `{"code":400,"status":"ERROR","data":null,"error":"Please check your input. Something went wrong","details":"invalid character 'p' after object key:value pair"}`)

	})

	t.Run("Register: when data not meet requirement", func(t *testing.T) {

		invalidCases := `{"email": "changecom", "password": "123", "confirm_password": "12345678", "username": "hellothere" }`

		code, res := integration.ServeRequest(router, "POST", path, invalidCases)
		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, "invalid input data: map[0:The email field must be a valid email address 1:The password field is invalid 2:The confirm_password field must be equal to Password")
	})

	t.Run("Register: when user is successfully registered", func(t *testing.T) {
		mockPass := integration.GetRandomString(10)
		validData := fmt.Sprintf(`{ "email": "changeme@gmail.com", "password": "%s", "confirm_password": "%s", "username": "hellothere" }`, mockPass, mockPass)
		code, res := integration.ServeRequest(router, "POST", path, validData)

		assert.Equal(t, code, http.StatusCreated)
		assert.Contains(t, res, "Account created successfully")

		// verify database
		user, _ := models.GetUserByUsernameOrEmail("hellothere", "")
		assert.NotNil(t, user)
		assert.Equal(t, user.Email, "changeme@gmail.com")

		// verify cache
		activeToken, err := db.GetDataFromCache("pending_active_user_hellothere")
		if err != nil {
			t.Errorf("Error getting data from cache: %v", err)
		}

		assert.NotNil(t, activeToken)
	})

	t.Run("Register: when account already exists", func(t *testing.T) {
		mockPass := integration.GetRandomString(10)
		validData := fmt.Sprintf(`{ "email": "changeme@gmail.com", "password": "%s", "confirm_password": "%s", "username": "hellothere" }`, mockPass, mockPass)
		code, res := integration.ServeRequest(router, "POST", path, validData)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, "Account with username: hellothere or email: changeme@gmail.com already exists")
	})

}

func TestLoginHandler(t *testing.T) {
	urlPathLogin := "/api/v1/login"
	router := integration.SetUpRouter()

	router.POST(urlPathLogin, handlers2.Login)

	t.Run("Login: when data invalid to bind json", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{""password": "changeme" "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, "Please check your input. Something went wrong")
	})
	t.Run("Login: when missing password", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{"email": "hao@gmail.com"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, constants.MissingParams)
	})
	t.Run("Login: when missing email and username", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{"password": "hellul"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, "Username or Email are required")
	})
	t.Run("Login: when user is not found", func(t *testing.T) {
		// Act
		validJsonRequest := `{"password": "not-found", "username": "not-found"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		assert.Equal(t, code, http.StatusNotFound)
		assert.Contains(t, res, constants.AccountNotFound)
	})
	t.Run("Login: when account exist and not activate yet", func(t *testing.T) {
		email := integration.GetRandomEmail()
		tempS := integration.GetRandomString(10)
		db.DB.Postgres.Save(&models.User{
			ID:       uuid.New(),
			Email:    email,
			Username: tempS,
			Password: tempS,
			Status:   constants.ACCOUNT_STATUS_INACTIVE,
		})

		jsonLoginRequest := fmt.Sprintf(`{"password": "%s", "username": "%s"}`, tempS, tempS)
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Contains(t, res, constants.AccountInactive)
	})

	t.Run("Login: when account exist and password is not match", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.DefaultCost)
		db.DB.Postgres.Save(&models.User{
			ID:       uuid.New(),
			Email:    "not_match@gmail.com",
			Username: "not_match",
			Password: string(hashedPassword),
			Status:   constants.ACCOUNT_STATUS_ACTIVE,
		})
		jsonLoginRequest := `{"password": "changem", "username": "not_match"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)

		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Contains(t, res, constants.PasswordDoesNotMatch)
	})

	t.Run("Login: when account exist and password is match", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.DefaultCost)
		db.DB.Postgres.Save(&models.User{
			ID:       uuid.New(),
			Email:    "match@gmail.com",
			Username: "match",
			Password: string(hashedPassword),
			Status:   constants.ACCOUNT_STATUS_ACTIVE,
		})

		jsonLoginRequest := `{"password": "changeme", "username": "match"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)
		assert.Equal(t, code, http.StatusOK)
		assert.Contains(t, res, "token")
		assert.Contains(t, res, "username")
	})
}

func TestActivateHandler(t *testing.T) {
	url := "/api/v1/activate"
	router := integration.SetUpRouter()

	router.POST(url, handlers2.Activate)

	mockUsername := integration.GetRandomString(6)
	mockPass := integration.GetRandomString(10)
	mockEmail := integration.GetRandomEmail()

	err := db.SaveDataToCache(fmt.Sprintf("pending_active_user_%s", mockUsername), "to-ke-n")
	if err != nil {
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(mockPass), bcrypt.DefaultCost)
	db.DB.Postgres.Save(&models.User{
		ID:       uuid.New(),
		Email:    mockEmail,
		Username: mockUsername,
		Password: string(hashedPassword),
		Status:   constants.ACCOUNT_STATUS_INACTIVE,
	})

	t.Run("Activate: when data is invalid to process", func(t *testing.T) {
		param := "?&token=token"
		code, res := integration.ServeRequestWithoutBody(router, "POST", url+param)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Contains(t, res, "Username and token are required")
	})

	t.Run("Activate: when key doesn't exist", func(t *testing.T) {
		param := "?username=hello&token=not-found"
		code, res := integration.ServeRequestWithoutBody(router, "POST", url+param)

		assert.Equal(t, code, http.StatusNotFound)
		assert.Contains(t, res, constants.TokenNotFount)
	})

	t.Run("Activate: when token is not match", func(t *testing.T) {
		param := fmt.Sprintf("?username=%s&token=not-match", mockUsername)
		code, res := integration.ServeRequestWithoutBody(router, "POST", url+param)

		assert.Equal(t, code, http.StatusForbidden)
		assert.Contains(t, res, constants.InvalidToken)
	})

	t.Run("Activate: when key exists and token is match", func(t *testing.T) {
		param := fmt.Sprintf("?username=%s&token=to-ke-n", mockUsername)
		code, res := integration.ServeRequestWithoutBody(router, "POST", url+param)

		assert.Equal(t, code, http.StatusOK)
		assert.Contains(t, res, constants.AccountActivated)

		// verify database
		user, _ := models.GetUserByUsernameOrEmail(mockUsername, "")
		assert.NotNil(t, user)

		assert.Equal(t, user.Status, constants.ACCOUNT_STATUS_ACTIVE)

		// verify cache
		activeToken, err := db.GetDataFromCache(fmt.Sprintf("pending_active_user_%s", mockUsername))
		assert.Equal(t, activeToken, "")
		assert.NotNil(t, err)
	})

}

func TestMain(m *testing.M) {
	integration.SetupTestServer()

	code := m.Run()

	integration.TearDownContainers()
	os.Exit(code)
}
