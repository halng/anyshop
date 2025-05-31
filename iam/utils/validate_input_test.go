/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"gte=18,lte=65"`
}

// TestValidateInputTest tests the ValidateInput function
func TestValidateInputTest(t *testing.T) {
	emailMock := "test@example.com"
	userMock := "testuser"

	t.Run("Valid input", func(t *testing.T) {
		input := TestStruct{
			Email:    emailMock,
			Username: userMock,
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.True(t, valid)
		assert.Nil(t, errors)
	})

	t.Run("Missing required field", func(t *testing.T) {
		input := TestStruct{
			Email:    emailMock,
			Username: "",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.True(t, strings.Contains(errors.Error(), "The username field is required"))
	})

	t.Run("Invalid email format", func(t *testing.T) {
		input := TestStruct{
			Email:    "invalid-email",
			Username: userMock,
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.True(t, strings.Contains(errors.Error(), "The email field must be a valid email address"))
	})

	t.Run("Field value less than minimum", func(t *testing.T) {
		input := TestStruct{
			Email:    emailMock,
			Username: userMock,
			Age:      17,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.True(t, strings.Contains(errors.Error(), "The age field must be greater than or equal to 18"))
	})

	t.Run("Field value greater than maximum", func(t *testing.T) {
		input := TestStruct{
			Email:    emailMock,
			Username: userMock,
			Age:      66,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.True(t, strings.Contains(errors.Error(), "The age field must be less than or equal to 65"))
	})

	t.Run("Missing multiple required fields", func(t *testing.T) {
		input := TestStruct{
			Email:    "",
			Username: "",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.True(t, strings.Contains(errors.Error(), "The email field is required"))
		assert.True(t, strings.Contains(errors.Error(), "The username field is required"))
	})
}
