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
	"github.com/halng/anyshop/logging"
	"go.uber.org/zap"
	"os"
)

func Initialize() {
	isCleanUp := os.Getenv("IS_CLEAN_DB")
	if isCleanUp == "1" {
		logging.LOGGER.Info("Cleaning up database...")
		cleanUp()
	} else {
		logging.LOGGER.Info("Skipping database cleanup")
	}

	DB := db.DB
	DB.Postgres.AutoMigrate(&User{})
	DB.Postgres.AutoMigrate(&Shop{})
	DB.Postgres.AutoMigrate(&Role{})
	DB.Postgres.AutoMigrate(&ShopUser{})
	DB.Postgres.AutoMigrate(&AccessPolicy{})
	//initMasterUser()
	initRoles()

}

//func initMasterUser() {
//	masterUsername := os.Getenv("MASTER_USERNAME")
//	masterPassword := os.Getenv("MASTER_PASSWORD")
//	masterEmail := os.Getenv("MASTER_EMAIL")
//	masterFirstName := os.Getenv("MASTER_FIRST_NAME")
//	masterLastName := os.Getenv("MASTER_LAST_NAME")
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(masterPassword), bcrypt.DefaultCost)
//	if err != nil {
//		logging.LOGGER.Error("Cannot hash password", zap.Any("error", err))
//		panic("Cannot hash password")
//	}
//	masterAccount := Account{
//		Username:  masterUsername,
//		Password:  string(hashedPassword),
//		Email:     masterEmail,
//		FirstName: masterFirstName,
//		LastName:  masterLastName}
//
//	_, err = masterAccount.SaveAccount()
//	if err != nil {
//		logging.LOGGER.Error("Cannot create master account", zap.Any("error", err))
//		panic("Cannot create master account")
//	}
//
//	//role := Role{
//	//	UserId:      savedAcc.ID,
//	//	Roles:       []string{RoleAppOwner},
//	//	Permissions: Permissions[RoleAppOwner],
//	//}
//	//
//	//if err = role.SaveRole(); err != nil {
//	//	logging.LOGGER.Error("Cannot create master role", zap.Any("error", err))
//	//	panic("Cannot create master role")
//	//}
//
//	logging.LOGGER.Info("Master account created successfully: " + masterUsername + " - " + masterPassword)
//}

func cleanUp() {
	db.DB.Postgres.Exec(`DO $$ DECLARE r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END $$;`)

}

func initRoles() {
	// Initialize roles and permissions
	logging.LOGGER.Info("Initializing roles...")

	for _, roleName := range constants.ROLES {
		// Check if the role already exists

		role := Role{
			Id:   uuid.New(),
			Name: roleName,
		}
		if err := role.SaveRole(); err != nil {
			logging.LOGGER.Error("Failed to save role", zap.Any("role", role), zap.Error(err))
		}
	}

	logging.LOGGER.Info("Roles initialized successfully")
}
