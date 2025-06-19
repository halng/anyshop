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
	"time"
)

type ShopUser struct {
	ShopID   uuid.UUID `gorm:"primaryKey" json:"shop_id"`
	Shop     Shop      `gorm:"foreignKey:ShopID;references:ID"`
	RoleID   uuid.UUID `gorm:"not null" json:"role_id"`
	Role     Role      `gorm:"foreignKey:RoleID;references:ID"`
	UserID   uuid.UUID `gorm:"primaryKey" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID;references:ID"`
	CreateAt int64     `json:"create_at"`
	UpdateAt int64     `json:"update_at"`
	CreateBy string    `json:"create_by"`
	UpdateBy string    `json:"update_by"`
}

func GetAllRole(userId string) ([]ShopUser, error) {
	var shopUsers []ShopUser
	if err := db.DB.Postgres.Model(&ShopUser{}).Where("user_id = ?", userId).
		Preload("Shop").Preload("Role").Find(&shopUsers).Error; err != nil {
		return nil, err
	}
	return shopUsers, nil
}

func GetAllShopUserByID(userId string) ([]ShopUser, error) {
	var shopUsers []ShopUser
	if err := db.DB.Postgres.Model(&ShopUser{}).Where("user_id = ?", userId).Preload("Shop").Preload("Role").Find(&shopUsers).Error; err != nil {
		return nil, err
	}
	return shopUsers, nil
}

func (shopUser *ShopUser) Save() error {
	shopUser.CreateAt = time.Now().Unix()
	if len(shopUser.CreateBy) == 0 {
		shopUser.CreateBy = shopUser.User.ID.String()
	}

	if err := db.DB.Postgres.Create(&shopUser).Error; err != nil {
		return err
	}

	return nil
}

func (shopUser *ShopUser) BeforeSave() error {
	shopUser.UpdateAt = time.Now().Unix()
	shopUser.UpdateBy = shopUser.User.ID.String()
	return nil
}
