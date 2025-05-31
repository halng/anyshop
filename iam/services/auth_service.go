/*
 * ****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/db"
	"github.com/halng/anyshop/dto"
	"github.com/halng/anyshop/kafka"
	"github.com/halng/anyshop/logging"
	"github.com/halng/anyshop/models"
	"github.com/halng/anyshop/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

// Register func to register new user account in the system
func Register(c *gin.Context) {
	userRegister, errors := parseAndValidateInput(c)
	if errors != nil {
		dto.BadRequestResponse(c, constants.MessageErrorBindJson, errors.Error())
		return
	}

	if models.ExistsByEmailOrUsername(userRegister.Email, userRegister.Username) {
		dto.BadRequestResponse(c, constants.AccountExists, nil)
		return
	}

	hashedPassword, err := generatePassword(userRegister.Password)
	if err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return
	}

	user := createAccount(userRegister, hashedPassword, "REGISTER")
	if !saveAccountAndRespond(c, user) {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, nil)
		return
	}

	postCreateAccount(user)
	dto.SuccessResponse(c, http.StatusCreated, constants.AccountCreated)
}

func Login(c *gin.Context) {
	var userInput dto.LoginRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		dto.BadRequestResponse(c, constants.MessageErrorBindJson, err)
		return
	}

	if _, errors := utils.ValidateInput(userInput); errors != nil {
		dto.BadRequestResponse(c, constants.MissingParams, errors.Error())
		return
	}

	if utils.IsNullOrEmpty(userInput.Username) && utils.IsNullOrEmpty(userInput.Email) {
		dto.BadRequestResponse(c, constants.BadRequest, "Username or Email are required")
		return
	}

	user, err := models.GetUserByUsernameOrEmail(userInput.Username, userInput.Email)
	if err != nil {
		dto.NotFoundResponse(c, constants.AccountNotFound, nil)
		return
	}

	if user.Status == constants.ACCOUNT_STATUS_INACTIVE {
		dto.UnauthorizedResponse(c, constants.AccountInactive, nil)
		return
	}

	if !user.ComparePassword(userInput.Password) {
		dto.UnauthorizedResponse(c, constants.PasswordDoesNotMatch, userInput)
		return
	}

	acls, err := models.GetAllAccessPoliciesByUserId(user.ID)

	if err != nil {
		logging.LOGGER.Error("Cannot get role for user %s ", zap.Any("user", user.ID.String()))
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
	}

	jwtToken, err := utils.GenerateJWT(user.ID.String(), user.Username, acls)
	if err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return
	}

	dto.SuccessResponse(c, http.StatusOK, dto.LoginResponse{ApiToken: jwtToken, Username: user.Username, Email: user.Email, ID: user.ID.String()})

}

func Activate(c *gin.Context) {
	username := c.Query("username")
	token := c.Query("token")

	if utils.IsNullOrEmpty(username) || utils.IsNullOrEmpty(token) {
		dto.BadRequestResponse(c, constants.BadRequest, "Username and token are required")
		return
	}

	key := fmt.Sprintf(constants.REDIS_PENDING_ACTIVE_STAFF_KEY, username)
	activeToken, err := db.GetDataFromCache(key)

	if activeToken == nil || activeToken == "" || err != nil {
		dto.NotFoundResponse(c, constants.TokenNotFount, err)
		return
	}

	if !utils.Equal(activeToken.(string), token) {
		dto.ForbiddenResponse(c, constants.InvalidToken, nil)
		return
	}

	user, err := models.GetUserByUsernameOrEmail(username, "")
	if err != nil {
		dto.NotFoundResponse(c, constants.AccountNotFound, nil)
		return
	}

	user.Status = constants.ACCOUNT_STATUS_ACTIVE
	if err = user.UpdateUser(); err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return
	}

	err = db.DeleteDataFromCache(key)
	if err != nil {
		logging.LOGGER.Error("Cannot delete token from cache", zap.String("key", key), zap.Error(err))
	}

	dto.SuccessResponse(c, http.StatusOK, constants.AccountActivated)
}

// PRIVATE FUNCTIONS

func postCreateAccount(user models.User) {
	// send message to kafka for new user account
	serializedMessage, token := getMessageForActiveNewUser(user)
	kafka.PushMessageNewUser(serializedMessage)
	logging.LOGGER.Info("New user account created", zap.String("user", user.Username))

	// save active token to redis
	key := fmt.Sprintf(constants.REDIS_PENDING_ACTIVE_STAFF_KEY, user.Username)
	err := db.SaveDataToCache(key, token)
	if err != nil {
		logging.LOGGER.Error("Cannot save token in cache")
	}

}

func createAccount(userInput dto.RegisterRequest, hashedPassword, createdBy string) models.User {
	return models.User{
		Email:    userInput.Email,
		Username: userInput.Username,
		Password: hashedPassword,
		CreateBy: createdBy,
		Status:   constants.ACCOUNT_STATUS_INACTIVE,
	}
}

func generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func saveAccountAndRespond(c *gin.Context, account models.User) bool {
	_, err := account.SaveUser()
	if err != nil {
		dto.InternalServerErrorResponse(c, constants.InternalServerError, err)
		return false
	}
	return true
}

func getMessageForActiveNewUser(user models.User) (string, string) {
	var activeNewUser dto.ActiveNewUser
	activeNewUser.Username = user.Username
	activeNewUser.Email = user.Email
	activeNewUser.Token = utils.ComputeHMAC256(user.Username, user.Email)
	expiredTime := time.Now().UnixMilli() + 1000*60*60*24
	activeNewUser.ExpiredTime = fmt.Sprintf("%d", expiredTime) // 1 day

	// build activation link for new user
	apiHost := os.Getenv("API_GATEWAY_HOST")
	activeNewUser.ActivationLink = fmt.Sprintf("%s/api/v1/iam/activate?username=%s&token=%s", apiHost, activeNewUser.Username, activeNewUser.Token)

	var activeNewUserMsg dto.ActiveNewUserMsg
	activeNewUserMsg.Action = constants.ActiveNewUserAction
	activeNewUserMsg.Data = activeNewUser

	serialized, err := json.Marshal(activeNewUserMsg)
	if err != nil {
		log.Printf("Cannot serialize data")
		return "", ""
	}
	return string(serialized), activeNewUser.Token
}

func parseAndValidateInput(c *gin.Context) (dto.RegisterRequest, error) {
	var userInput dto.RegisterRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		return userInput, err
	}

	if ok, errors := utils.ValidateInput(userInput); !ok {
		return userInput, errors
	}

	return userInput, error(nil)
}
