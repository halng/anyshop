/*
 * *****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * *****************************************************************************************
 */

package com.app.anyshop.cms.constant;

public class Message {
  public static final String SUCCESS = "Success";
  public static final String VALIDATE_FAILED =
      "The request is invalid. Please review your input or contact support for assistance.";
  public static final String CREATED_WAITING_APPROVAL = "Created, waiting for approval";
  public static final String UPDATED_WAITING_APPROVAL = "Updated, waiting for approval";
  public static final String CATEGORY_CAN_NOT_UPDATE = "Category %s have status %s, can not update";
  // Redis

  // example: "user-id:category:update:xxx-xxxx"
  public static final String REDIS_KEY_UPDATE_CATEGORY_TEMPLATE = "%s:category:update:%s";

  public static class Constants {
    public static final int MAX_PER_REQUEST = 25;
    public static final long DEFAULT_EXPIRED_TIME = 60 * 60 * 24;
  }
}
