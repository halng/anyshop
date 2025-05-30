/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package constants

// Message constants
const (
	MessageSuccess       = "Success"
	MessageErrorBindJson = "Please check your input. Something went wrong"

	BadRequest                 = "Bad Request"
	Unauthorized               = "Unauthorized"
	Forbidden                  = "Forbidden"
	NotFound                   = "Not Found"
	Conflict                   = "Conflict"
	InternalServerError        = "There was an error processing your request. Please try again later"
	TokenNotFount              = "Your login session has expired. Please login again"
	MissingCredentials         = "Missing credentials. X-API-SECRET-TOKEN and X-API-USER-ID are required"
	MissingParams              = "Missing required parameters. Please check your input"
	ForbiddenMissingPermission = "You do not have permission to perform this action"
	// account constant
	AccountCreated       = "Account created successfully"
	AccountNotFound      = "Account not found"
	AccountUpdated       = "Account updated successfully"
	AccountDeleted       = "Account deleted successfully"
	DefaultCreator       = "SYSTEM"
	AccountExists        = "Account with username: %s or email: %s already exists"
	PasswordDoesNotMatch = "Invalid credentials"
	InvalidPermission    = "User does not have permission to access this resource"
	DefaultPassword      = "12345678"
	AccountInactive      = "Account is inactive. Please check your email to activate your account"
	InvalidToken         = "Invalid token. Please check your activation link or token"
	AccountActivated     = "Account activated successfully"
	// define key
	ApiTokenRequestHeader  = "X-API-SECRET-TOKEN"
	ApiUserIdRequestHeader = "X-API-USER-ID"
	ApiUserRequestHeader   = "X-API-USER"
	ApiUserRole            = "X-API-USER-ROLE"
	ApiOriginMethod        = "X-API-ORIGIN-METHOD"

	ActiveNewUserAction = "ACTIVATE_NEW_STAFF"

	ACCOUNT_STATUS_ACTIVE   = "ACTIVE"
	ACCOUNT_STATUS_INACTIVE = "INACTIVE"
	ACCOUNT_STATUS_DELETED  = "DELETED"

	REDIS_PENDING_ACTIVE_STAFF_KEY = "pending_active_user_%s"
)

// ROLES defines the available roles in the system
var ROLES = []string{"ADMIN", "MANAGER", "STAFF"}
