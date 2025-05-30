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

type Role struct {
	Id       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"unique;not null" json:"name"`
	CreateAt int64     `json:"createAt"`
	UpdateAt int64     `json:"updateAt"`
	CreateBy string    `json:"createBy"`
	UpdateBy string    `json:"updateBy"`
}

func (role *Role) SaveRole() error {

	role.CreateAt = time.Now().Unix()
	if len(role.CreateBy) == 0 {
		role.CreateBy = constants.DefaultCreator
	}

	if err := db.DB.Postgres.Create(&role).Error; err != nil {
		return err
	}

	return nil

}

func (role *Role) BeforeSave() error {
	role.UpdateAt = time.Now().Unix()
	role.UpdateBy = constants.DefaultCreator
	return nil
}
