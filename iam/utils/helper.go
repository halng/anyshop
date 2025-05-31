/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
 */

package utils

func IsNullOrEmpty(s string) bool {
	return len(s) == 0 || s == ""
}

func Equal(a, b string) bool {
	return a == b && !IsNullOrEmpty(a) && !IsNullOrEmpty(b)
}
