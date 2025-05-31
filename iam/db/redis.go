/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package db

import (
	"context"
	"time"
)

var (
	DefaultCacheExpireTime = 1
)

func GetDataFromCache(key string) (interface{}, error) {
	ctx := context.Background()

	redisClient := DB.Redis

	value, err := redisClient.Get(ctx, key).Result()

	return value, err
}

func SaveDataToCache(key string, data interface{}) error {
	ctx := context.Background()

	redisClient := DB.Redis
	err := redisClient.Set(ctx, key, data, time.Duration(DefaultCacheExpireTime)*time.Hour).Err()
	return err
}

func DeleteDataFromCache(key string) error {
	ctx := context.Background()

	redisClient := DB.Redis
	err := redisClient.Del(ctx, key).Err()
	return err
}
