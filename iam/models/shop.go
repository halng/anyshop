/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package models

import "github.com/google/uuid"

type Shop struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Slug     string    `gorm:"not null,unique" json:"slug"`
	IsActive bool      `gorm:"default:true" json:"is_active"`
	CreateAt int64     `json:"create_at"`
	UpdateAt int64     `json:"update_at"`
	CreateBy string    `json:"create_by"`
	UpdateBy string    `json:"update_by"`
}
