/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.annotation;

import com.app.anyshop.api.constant.ActionStatus;
import jakarta.validation.ConstraintValidator;
import jakarta.validation.ConstraintValidatorContext;

/**
 * Validator for the {@link ValidAction} annotation.
 *
 * <p>This validator checks if the value of the annotated field is a valid {@link ActionStatus} value. The
 * value is considered valid if it can be parsed into a {@link ActionStatus} enum. If the value is {@code
 * null}, it is considered invalid.
 */
public class ValidActionValidator implements ConstraintValidator<ValidAction, Object> {

  @Override
  public boolean isValid(Object o, ConstraintValidatorContext constraintValidatorContext) {
    if (null == o) {
      return false;
    }
    try {
      ActionStatus.valueOf(o.toString());
      return true;
    } catch (Exception e) {
      throw new IllegalArgumentException("Invalid action.");
    }
  }
}
