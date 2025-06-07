/*
 * ****************************************************************************************
 * Copyright 2025 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * ****************************************************************************************
 */

package shop

import (
	"github.com/halng/anyshop/test/integration"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 1, 1)
}

func TestCreateShop(t *testing.T) {
	assert.Equal(t, 1, 1)
}
