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
	ShopId string `json:"shop_id"`
	Role   string `json:"role"`
}

func getAcls(acls []models.ShopUser) []ACL {
	var result []ACL

	for _, policy := range acls {
		acl := ACL{
			ShopId: policy.ShopID.String(),
			Role:   policy.Role.Name,
		}

		// Check if the ACL for this shop already exists
		exists := false
		for _, existingACL := range result {
			if existingACL.ShopId == acl.ShopId && existingACL.Role == acl.Role {
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

func GenerateJWT(id string, username string, acls []models.ShopUser) (string, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)
	acl := getAcls(acls)

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

func getClaimsFromJWT(tokenString string) (jwt.MapClaims, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(apiSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

func ExtractJWT(tokenString string) (string, string, error) {

	claims, err := getClaimsFromJWT(tokenString)
	if err != nil {
		return "", "", err
	}

	id := claims["sub"].(string)
	username := claims["name"].(string)

	return id, username, nil
}
