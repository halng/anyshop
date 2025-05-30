/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/halng/anyshop/models"
	"os"
	"time"
)

var (
	EnvApiSecretKey = "API_SECRET"
)

type ACL struct {
	ShopId      string      `json:"shop_id"`
	Role        models.Role `json:"role"`
	Permissions []string    `json:"permissions"`
}

func getAclsFromPolicies(acls []models.AccessPolicy) []ACL {
	var result []ACL

	for _, policy := range acls {
		acl := ACL{
			ShopId:      policy.ShopUser.ShopID.String(),
			Role:        policy.ShopUser.Role,
			Permissions: []string{policy.Action},
		}

		// Check if the ACL for this shop already exists
		exists := false
		for i, existingACL := range result {
			if existingACL.ShopId == acl.ShopId && existingACL.Role == acl.Role {
				result[i].Permissions = append(result[i].Permissions, acl.Permissions...)
				exists = true
				break
			}
		}

		if !exists {
			result = append(result, acl)
		}
	}

	return result
}

func GenerateJWT(id string, username string, acls []models.AccessPolicy) (string, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)
	acl := getAclsFromPolicies(acls)

	claims := jwt.MapClaims{
		"sub":  id,
		"name": username,
		"acl":  acl,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  "iam",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))
}
