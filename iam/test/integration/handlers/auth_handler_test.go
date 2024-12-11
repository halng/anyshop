package handlers

import (
	"encoding/json"
	"github.com/halng/anyshop/db"
	"github.com/halng/anyshop/models"
	"net/http"
	"os"
	"testing"

	"github.com/halng/anyshop/constants"
	handlers2 "github.com/halng/anyshop/handlers"
	"github.com/halng/anyshop/middleware"
	"github.com/halng/anyshop/test/integration"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	urlPathLogin := "/api/v1/login"
	router := integration.SetUpRouter()

	router.POST(urlPathLogin, handlers2.Login)

	/**
	Test case 1: Login with master credentials.
	Test case 2: Check for invalid JSON data binding.
	Test case 3: User not found scenario.
	Test case 4: Existing account with incorrect password.
	*/

	t.Run("Login: when master credential is used", func(t *testing.T) {
		// Act
		validJsonRequest := `{"password": "changeme", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		// Assert
		if code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, code)
		}
		assert.Equal(t, code, http.StatusOK)
		assert.Contains(t, res, "token")
	})
	t.Run("Login: when data invalid to bind json", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{""password": "changeme" "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Equal(t, res, `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`)
	})
	t.Run("Login: when user is not found", func(t *testing.T) {
		// Act
		validJsonRequest := `{"password": "not-found", "username": "not-found"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		assert.Equal(t, code, http.StatusNotFound)
		assert.Equal(t, res, `{"code":404,"error":"Account not found","status":"ERROR"}`)
	})
	t.Run("Login: when account exist and password is not match", func(t *testing.T) {
		// account register successfully in #TestRegister: Register when user is successfully registered
		// check
		jsonLoginRequest := `{"password": "not-match", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)

		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Invalid credentials","status":"ERROR"}`)

	})
}

// TestCreateStaffHandler combine both test for middleware and CreateStaff Func
func TestCreateStaffHandler(t *testing.T) {
	urlPathCreateStaff := "/api/v1/create-staff"
	urlPathLogin := "/api/v1/login"

	headers := map[string]string{
		constants.ApiTokenRequestHeader:  "test-secret",
		constants.ApiUserIdRequestHeader: "test-user",
	}

	router := integration.SetUpRouter()
	router.POST(urlPathCreateStaff, middleware.ValidateRequest, handlers2.CreateStaff)
	router.POST(urlPathLogin, handlers2.Login)

	var masterAdminAuthObject map[string]interface{}

	t.Run("Create Staff & Validate Request: when request not fulfill - Missing Header", func(t *testing.T) {
		// Act
		code, res := integration.ServeRequest(router, "POST", urlPathCreateStaff, "")

		// Assert
		if code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}
		expectedResponse := `{"code":401,"error":"Missing credentials. X-API-SECRET-TOKEN and X-API-USER-ID are required","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 401, code)
	})

	t.Run("Create Staff & Validate Request: when request not valid", func(t *testing.T) {
		// act
		code, res, _ := integration.ServeRequestWithHeader(router, "POST", urlPathCreateStaff, "", headers)

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Your login session has expired. Please login again","status":"ERROR"}`)
	})

	t.Run("Create Staff: Master Admin Login", func(t *testing.T) {
		loginJsonRequest := `{"password": "changeme", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, loginJsonRequest)

		var data map[string]interface{}
		_ = json.Unmarshal([]byte(res), &data)

		masterAdminAuthObject = data["data"].(map[string]interface{})

		headers[constants.ApiTokenRequestHeader] = masterAdminAuthObject["api-token"].(string)
		headers[constants.ApiUserIdRequestHeader] = masterAdminAuthObject["id"].(string)

		assert.Equal(t, 200, code)
	})

	t.Run("Create Staff: when requester have permission but input not valid", func(t *testing.T) {

		// can not bind json data
		code, res, _ := integration.ServeRequestWithHeader(
			router,
			"POST",
			urlPathCreateStaff,
			`{"email":"test@gmail.com",
					"username": "test,
					"password" "test",
					"lastname": "test",
					"firstname": "test"}`,
			headers)

		assert.Equal(t, http.StatusBadRequest, code)
		assert.Equal(t, res, `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`)

		// invalid input
		code, res, _ = integration.ServeRequestWithHeader(
			router,
			"POST",
			urlPathCreateStaff,
			`{"email":"not-valid-email",
					"username": "test",
					"password": "test",
					"lastname": "test",
					"firstname": "test"}`,
			headers)

		assert.Equal(t, http.StatusBadRequest, code)
		assert.Contains(t, res, "must be a valid email address")

	})

	t.Run("Create Staff: when user is successfully registered", func(t *testing.T) {
		// Act
		validUserInput := `{"email":"createdstaff@gmail.com", "username": "createdstaff", "lastname": "createdstaff", "firstname": "createdstaff"}`
		code, res, _ := integration.ServeRequestWithHeader(router, "POST", urlPathCreateStaff, validUserInput, headers)

		// Assert
		if code != http.StatusCreated {
			t.Errorf("Expected status code %d but got %d", http.StatusCreated, code)
		}
		expectedResponse := `{"code":201,"data":null,"status":"SUCCESS"}`
		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Create Staff: when user is exists", func(t *testing.T) {
		// Act
		validUserInput := `{"email":"createdstaff@gmail.com", "username": "createdstaff", "lastname": "createdstaff", "firstname": "createdstaff"}`
		code, res, _ := integration.ServeRequestWithHeader(router, "POST", urlPathCreateStaff, validUserInput, headers)

		// Assert
		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}
		expectedResponse := `{"code":400,"error":"Account with username: createdstaff@gmail.com or email: createdstaff already exists","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Create Staff: when requester don't have permission", func(t *testing.T) {
		// newly created user login
		loginJson := `{"password": "12345678", "username": "createdstaff"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, loginJson)

		assert.Equal(t, http.StatusOK, code)
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(res), &data)

		objectData := data["data"].(map[string]interface{})

		headers[constants.ApiTokenRequestHeader] = objectData["api-token"].(string)
		headers[constants.ApiUserIdRequestHeader] = objectData["id"].(string)

		// create new user with staff role
		// Act
		validUserInput := `{"email":"no-nope@gmail.com", "username": "nope", "password": "nope", "lastname": "nope", "firstname": "nope"}`
		code, res, _ = integration.ServeRequestWithHeader(router, "POST", urlPathCreateStaff, validUserInput, headers)

		assert.Equal(t, http.StatusForbidden, code)
		assert.Equal(t, res, `{"code":403,"error":"User does not have permission to access this resource","status":"ERROR"}`)
	})
}

func TestValidate(t *testing.T) {
	urlPathValidate := "/api/v1/validate"
	urlPathLogin := "/api/v1/login"

	router := integration.SetUpRouter()

	router.GET(urlPathValidate, handlers2.Validate)
	router.POST(urlPathLogin, handlers2.Login)

	t.Run("Validate: when api token is missing", func(t *testing.T) {
		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathValidate, "")

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Unauthorized","status":"ERROR"}`)
	})
	t.Run("Validate: when api token is invalid", func(t *testing.T) {
		// Act
		code, res, _ := integration.ServeRequestWithHeader(router, "GET", urlPathValidate, "", nil)

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Unauthorized","status":"ERROR"}`)

	})
	t.Run("Validate: when user token was expired", func(t *testing.T) {
		headers := map[string]string{
			constants.ApiTokenRequestHeader:  "XXX",
			constants.ApiUserIdRequestHeader: "USER_ID",
		}

		// act
		code, res, _ := integration.ServeRequestWithHeader(router, "GET", urlPathValidate, "", headers)

		// Assert
		assert.Equal(t, code, http.StatusBadRequest)
		assert.Equal(t, res, `{"code":400,"error":"Your login session has expired. Please login again","status":"ERROR"}`)
	})

	t.Run("Validate: success validate user", func(t *testing.T) {
		// Login
		loginJsonRequest := `{"password": "changeme", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, loginJsonRequest)

		assert.Equal(t, 200, code)
		var data map[string]interface{}
		err := json.Unmarshal([]byte(res), &data)

		authData := data["data"].(map[string]interface{})
		//err = json.Unmarshal([]byte(authData), &resData)
		assert.NoError(t, err)

		headers := map[string]string{
			constants.ApiTokenRequestHeader:  authData["api-token"].(string),
			constants.ApiUserIdRequestHeader: authData["id"].(string),
		}

		// act
		code, res, resHeaders := integration.ServeRequestWithHeader(router, "GET", urlPathValidate, "", headers)

		// Assert
		assert.Equal(t, code, http.StatusOK)
		assert.Equal(t, res, `{"code":200,"data":null,"status":"SUCCESS"}`)

		assert.Equal(t, "super_admin", resHeaders.Get(constants.ApiUserRole))
		assert.Equal(t, "changeme", resHeaders.Get(constants.ApiUserRequestHeader))
		assert.Equal(t, authData["id"].(string), resHeaders.Get(constants.ApiUserIdRequestHeader))

	})
}

func TestActivateStaff(t *testing.T) {
	urlPathActivate := "/api/v1/activate"

	router := integration.SetUpRouter()
	router.GET(urlPathActivate, handlers2.Activate)

	t.Run("Activate: when require parameters are missing", func(t *testing.T) {
		invalidPath := [3]string{"", "?username=testuser", "?token=testtoken"}
		for _, path := range invalidPath {

			// Act
			code, res := integration.ServeRequest(router, "GET", urlPathActivate+path, "")

			// Assert
			assert.Equal(t, code, http.StatusBadRequest)
			assert.Equal(t, res, `{"code":400,"error":"Missing required parameters. Please check your input","status":"ERROR"}`)
		}
	})

	t.Run("Activate: when token is not found in cache", func(t *testing.T) {
		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathActivate+"?token=testtoken&username=testuser", "")

		// Assert
		assert.Equal(t, code, http.StatusNotFound)
		assert.Equal(t, res, `{"code":404,"error":"Your login session has expired. Please login again","status":"ERROR"}`)
	})

	t.Run("Activate: when token is invalid", func(t *testing.T) {
		// Mock cache data
		err := db.SaveDataToCache("pending_active_user_testuser", `validtoken`)
		if err != nil {
			return
		}

		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathActivate+"?token=invalidtoken&username=testuser", "")

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Unauthorized","status":"ERROR"}`)
	})

	t.Run("Activate: when account is not found", func(t *testing.T) {
		// Mock cache data
		err := db.SaveDataToCache("pending_active_user_testuser", `validtoken`)
		if err != nil {
			return
		}

		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathActivate+"?token=validtoken&username=testuser", "")

		// Assert
		assert.Equal(t, code, http.StatusNotFound)
		assert.Equal(t, res, `{"code":404,"error":"Account not found","status":"ERROR"}`)
	})

	t.Run("Activate: when account is successfully activated", func(t *testing.T) {
		// Mock cache data
		err := db.SaveDataToCache("pending_active_user_testuser", `validtoken`)
		if err != nil {
			return
		}

		// Mock account data
		account := models.Account{
			Username: "testuser",
			Status:   constants.ACCOUNT_STATUS_INACTIVE,
		}
		_, err = account.SaveAccount()
		if err != nil {
			return
		}

		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathActivate+"?token=validtoken&username=testuser", "")

		// Assert
		assert.Equal(t, code, http.StatusOK)
		assert.Equal(t, res, `{"code":200,"data":null,"status":"SUCCESS"}`)

		// Verify account status
		updatedAccount, _ := models.GetAccountByUsername("testuser")
		assert.Equal(t, constants.ACCOUNT_STATUS_ACTIVE, updatedAccount.Status)
	})
}

func TestMain(m *testing.M) {
	integration.SetupTestServer()

	code := m.Run()

	integration.TearDownContainers()
	os.Exit(code)
}
