/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.annotation;

import jakarta.validation.Constraint;
import jakarta.validation.Payload;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Constraint(validatedBy = ValidActionValidator.class)
@Target({ElementType.PARAMETER, ElementType.FIELD})
@Retention(RetentionPolicy.RUNTIME)
public @interface ValidAction {
	String message() default "Invalid action.";

	Class<?>[] groups() default {};

	Class<? extends Payload>[] payload() default {};
}
