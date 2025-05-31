/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package utils

import (
	"github.com/google/uuid"
	"github.com/halng/anyshop/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	// Setup
	apiSecret := "testsecret"
	os.Setenv(EnvApiSecretKey, apiSecret)

	t.Run("Create and extract JWT", func(t *testing.T) {
		username := "changeme"
		id := "XXX-YYY-ZZZ"

		// Test JWT
		token, err := GenerateJWT(id, username, nil)
		if err != nil {
			t.Errorf("Error generating JWT: %v", err)
		}

		if token == "" {
			t.Errorf("Token is empty")
		}

		assert.True(t, len(token) > 0)
	})

	t.Run("Create and extract JWT with ACLs", func(t *testing.T) {
		username := "changeme"
		id := "XXX-YYY-ZZZ"
		// Set up environment variable for API secret key
		roleAdmin := models.Role{
			Id:   uuid.New(),
			Name: "ADMIN",
		}

		roleManager := models.Role{
			Id:   uuid.New(),
			Name: "MANAGER",
		}
		// Test JWT with ACLs
		acls := []models.AccessPolicy{
			{
				Action: "read",
				ShopUser: models.ShopUser{
					ShopID: uuid.New(),
					Role:   roleAdmin,
				},
			},
			{
				Action: "write",
				ShopUser: models.ShopUser{
					ShopID: uuid.New(),
					Role:   roleAdmin,
				},
			},
			{
				Action: "delete",
				ShopUser: models.ShopUser{
					ShopID: uuid.New(),
					Role:   roleManager,
				},
			},
		}

		token, err := GenerateJWT(id, username, acls)
		if err != nil {
			t.Errorf("Error generating JWT with ACLs: %v", err)
		}

		if token == "" {
			t.Errorf("Token is empty")
		}

		assert.True(t, len(token) > 0)
	})
}
