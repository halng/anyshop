/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package shop

import (
	handlers2 "github.com/halng/anyshop/controller"
	"github.com/halng/anyshop/test/integration"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	integration.SetupTestServer()

	code := m.Run()

	integration.TearDownContainers()
	os.Exit(code)
}

func TestGetAllShops(t *testing.T) {
	url := "/api/v1/iam/shops"
	router := integration.SetUpRouter()

	router.GET(url, handlers2.GetAllShops)

}

func TestCreateShop(t *testing.T) {
	url := "/api/v1/iam/shops"
	router := integration.SetUpRouter()

	router.POST(url, handlers2.GetAllShops)
}
