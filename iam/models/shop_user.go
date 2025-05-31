/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package models

import (
	"github.com/google/uuid"
)

type ShopUser struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	ShopID   uuid.UUID `gorm:"not null" json:"shop_id"`
	Shop     Shop      `gorm:"foreignKey:ShopID;references:ID"`
	RoleID   uuid.UUID `gorm:"not null" json:"role_id"`
	Role     Role      `gorm:"foreignKey:RoleID;references:ID"`
	UserID   uuid.UUID `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID;references:ID"`
	CreateAt int64     `json:"create_at"`
	UpdateAt int64     `json:"update_at"`
	CreateBy string    `json:"create_by"`
	UpdateBy string    `json:"update_by"`
}
