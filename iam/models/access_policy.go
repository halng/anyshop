/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package models

import (
	"github.com/google/uuid"
	"github.com/halng/anyshop/db"
)

type AccessPolicy struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	ShopUserId uuid.UUID `gorm:"not null" json:"shop_user_id"`
	ShopUser   ShopUser  `gorm:"foreignKey:ShopUserId,references:ID" json:"shop_user"`
	Resource   string    `gorm:"not null" json:"resource"`
	Action     string    `gorm:"not null" json:"action"`
	Effect     string    `gorm:"not null" json:"effect"`
	CreateAt   int64     `json:"create_at"`
	UpdateAt   int64     `json:"update_at"`
	CreateBy   string    `json:"create_by"`
	UpdateBy   string    `json:"update_by"`
}

func GetAllAccessPoliciesByUserId(userId uuid.UUID) ([]AccessPolicy, error) {
	var policies []AccessPolicy
	if err := db.DB.Postgres.Where("shop_user_id = ?", userId).Find(&policies).Error; err != nil {
		return nil, err
	}
	return policies, nil
}
