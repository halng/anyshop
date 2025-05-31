/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package utils

import "testing"

func TestHelper(t *testing.T) {
	t.Run("IsNullOrEmpty", func(t *testing.T) {
		if !IsNullOrEmpty("") {
			t.Errorf("Expected true for empty string")
		}
		if IsNullOrEmpty("not empty") {
			t.Errorf("Expected false for non-empty string")
		}
	})

	t.Run("Equal", func(t *testing.T) {
		if !Equal("test", "test") {
			t.Errorf("Expected strings to be equal")
		}
		if Equal("test", "") {
			t.Errorf("Expected strings to not be equal")
		}
	})
}
