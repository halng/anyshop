package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
)

// ========================= Functions =========================

// CreateStaff create a new staff account and send a message to kafka
func CreateStaff(c *gin.Context) {
	var userInput dto.RegisterRequest

	// check if requester is super admin
	requesterRole, _ := c.Get(constants.ApiUserRoles)
	requesterId := c.GetHeader(constants.ApiUserIdRequestHeader)
	if requesterRole != models.RoleSuperAdmin {
		ResponseErrorHandler(c, http.StatusForbidden, constants.InvalidPermission, nil)
		return
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, userInput)
		return
	}

	if ok, errors := utils.ValidateInput(userInput); !ok {
		ResponseErrorHandler(c, http.StatusBadRequest, errors, userInput)
		return
	}

	if models.ExistsByEmailOrUsername(userInput.Email, userInput.Username) {
		ResponseErrorHandler(c, http.StatusBadRequest, fmt.Sprintf(constants.AccountExists, userInput.Email, userInput.Username), userInput)
		return
	}

	// only for newly created user. default password is "123456"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(constants.DefaultPassword), bcrypt.DefaultCost)
	if err != nil {
		logging.LOGGER.Error("CreateStaff", zap.Error(err))
		ResponseErrorHandler(c, http.StatusInternalServerError, constants.InternalServerError, nil)
		return
	}
	account := models.Account{}
	account.Email = userInput.Email
	account.Username = userInput.Username
	account.Password = string(hashedPassword)
	account.LastName = userInput.LastName
	account.FirstName = userInput.FirstName
	account.CreateBy = requesterId
	account.Status = constants.ACCOUNT_STATUS_INACTIVE

	// get role id for staff
	roleId, err := models.GetRoleIdByName(models.RoleStaff)
	if err != nil {
		logging.LOGGER.Error("Error when getting role id.", zap.Any("error", err))
		ResponseErrorHandler(c, http.StatusInternalServerError, constants.InternalServerError, nil)
		return
	}

	account.RoleId = roleId

	_, err = account.SaveAccount()

	if err != nil {
		logging.LOGGER.Error("Error when saving account.", zap.Any("error", err))
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, account)
		return
	}

	// send msg to kafka
	serializedMessage := account.GetSerializedMessageForActiveNewUser()
	_ = db.SaveActiveTokenToCache(account.Username, serializedMessage)

	kafka.PushMessageNewUser(serializedMessage)

	ResponseSuccessHandler(c, http.StatusCreated, nil)
}

// Login verify user credentials and return uuid pair with token saved in redis
func Login(c *gin.Context) {
	var userInput dto.LoginRequest

	if err := c.ShouldBindJSON(&userInput); err != nil {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MessageErrorBindJson, userInput)
		return
	}

	var account models.Account
	var err error

	if account, err = models.GetAccountByUsername(userInput.Username); err != nil {
		ResponseErrorHandler(c, http.StatusNotFound, constants.AccountNotFound, userInput)
		return
	}

	if !account.ComparePassword(userInput.Password) {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.PasswordNotMatch, userInput)
		return
	}

	token := account.GenerateAccessToken()
	ResponseSuccessHandler(c, http.StatusOK, dto.LoginResponse{ApiToken: token, Username: account.Username, Email: account.Email, ID: account.ID.String()})

}

// Validate user credentials and return username and role
func Validate(c *gin.Context) {
	// get api token from header
	apiToken := c.GetHeader(constants.ApiTokenRequestHeader)
	userId := c.GetHeader(constants.ApiUserIdRequestHeader)

	if apiToken == "" || userId == "" {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.Unauthorized, apiToken)
		return
	}

	// get bearer token from redis
	hashMD5 := utils.ComputeMD5([]string{userId})
	accessToken, err := db.GetDataFromKey(fmt.Sprintf("%s_%s", hashMD5, apiToken))
	if err != nil || accessToken == nil || accessToken == "" {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.TokenNotFount, accessToken)
		return
	}

	isValidToken, userId, username, role := utils.ExtractDataFromToken(accessToken.(string))
	if !isValidToken {
		ResponseErrorHandler(c, http.StatusInternalServerError, constants.InternalServerError, apiToken)
		return
	}
	ResponseSuccessHandler(c, http.StatusOK, nil)
	c.Header(constants.ApiUserIdRequestHeader, userId)
	c.Header(constants.ApiUserRoles, role)
	c.Header(constants.ApiUserRequestHeader, username)

}

func Activate(c *gin.Context) {
	token := c.Query("token")
	username := c.Query("username")
	
	if token == "" || username == "" {
		ResponseErrorHandler(c, http.StatusBadRequest, constants.MissingParams, nil)
		return
	}

	savedData, err := db.GetDataFromKey("active_" + username)
	if err != nil || savedData == nil {
		ResponseErrorHandler(c, http.StatusNotFound, constants.TokenNotFount, nil)
		return
	}
	
	var savedDataDto dto.ActiveNewUserMsg
	if err = json.Unmarshal(savedData.([]byte), &savedDataDto); err != nil {
		ResponseErrorHandler(c, http.StatusInternalServerError, constants.InternalServerError, nil)
		return
	}

	
	if token != savedDataDto.Data.Token {
		ResponseErrorHandler(c, http.StatusUnauthorized, constants.TokenNotFount, nil)
		return
	}

	account, err := models.GetAccountByUsername(username)
	if err != nil {
		ResponseErrorHandler(c, http.StatusNotFound, constants.AccountNotFound, nil)
		return
	}

	account.Status = constants.ACCOUNT_STATUS_ACTIVE
	_, err = account.SaveAccount()
	if err != nil {
		ResponseErrorHandler(c, http.StatusInternalServerError, constants.InternalServerError, nil)
		return
	}

	ResponseSuccessHandler(c, http.StatusOK, nil)
}