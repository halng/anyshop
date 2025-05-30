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
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `json:"password"`
	Status   string    `json:"status"`
	CreateAt int64     `json:"create_at"`
	UpdateAt int64     `json:"update_at"`
	CreateBy string    `json:"create_by"`
	UpdateBy string    `json:"update_by"`
}

func ExistsByEmailOrUsername(email string, username string) bool {
	var count int64
	db.DB.Postgres.Model(&User{}).Where("email = ? OR username = ?", email, username).Count(&count)
	return count > 0
}

func (user *User) SaveUser() (*User, error) {

	user.ID = uuid.New()
	user.CreateAt = time.Now().Unix()
	if user.CreateBy != "" {
		user.CreateBy = constants.DefaultCreator
	}

	if err := db.DB.Postgres.Create(&user).Error; err != nil {
		return &User{}, err
	}

	return user, nil

}

func (user *User) UpdateUser() error {
	user.UpdateAt = time.Now().Unix()
	user.UpdateBy = user.Username
	if err := db.DB.Postgres.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) BeforeSave() error {
	user.UpdateAt = time.Now().Unix()
	user.UpdateBy = user.Username
	return nil
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func GetUserByUsernameOrEmail(username, email string) (*User, error) {
	var user User
	if err := db.DB.Postgres.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
