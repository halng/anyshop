/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package models

import (
	"github.com/google/uuid"
	"github.com/halng/anyshop/constants"
	"github.com/halng/anyshop/db"
	"time"
)

type Shop struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"not null,unique" json:"name"`
	Slug     string    `gorm:"not null,unique" json:"slug"`
	IsActive bool      `gorm:"default:true" json:"is_active"`
	CreateAt int64     `json:"create_at"`
	UpdateAt int64     `json:"update_at"`
	CreateBy string    `json:"create_by"`
	UpdateBy string    `json:"update_by"`
}

func (shop *Shop) BeforeSave() error {
	shop.UpdateAt = time.Now().Unix()
	if len(shop.UpdateBy) == 0 {
		shop.UpdateBy = constants.DefaultCreator
	}
	return nil
}

func (shop *Shop) Save() error {
	shop.ID = uuid.New()
	shop.CreateAt = time.Now().Unix()
	if len(shop.CreateBy) == 0 {
		shop.CreateBy = constants.DefaultCreator
	}

	if err := db.DB.Postgres.Create(&shop).Error; err != nil {
		return err
	}

	return nil
}

func ExistByNameOrSlug(name, slug string) (bool, error) {
	var count int64
	if err := db.DB.Postgres.Model(&Shop{}).Where("name = ? OR slug = ?", name, slug).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
