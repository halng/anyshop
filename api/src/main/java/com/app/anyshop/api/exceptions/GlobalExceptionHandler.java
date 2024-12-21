/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.exceptions;

import com.app.anyshop.api.viewmodel.response.ResVM;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@Slf4j
@RestControllerAdvice
public class GlobalExceptionHandler extends ResponseEntityExceptionHandler {

  @ExceptionHandler(IllegalArgumentException.class)
  public ResponseEntity<ResVM> illegalArgumentExceptionHandler(IllegalArgumentException ex) {
    log.error("IllegalArgumentException: Validation error: {}", ex.getMessage());
    var errorRes =
        ResVM.builder()
            .code(HttpStatus.BAD_REQUEST)
            .msg(
                "The request is invalid. Please review your input or contact support for assistance.")
            .build();
    return ResponseEntity.ok(errorRes);
  }
}
