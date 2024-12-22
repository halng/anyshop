/*
 * *****************************************************************************************
 * Copyright 2024 By Hal Nguyen
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * *****************************************************************************************
 */

package com.app.anyshop.cms.controllers;

import com.app.anyshop.cms.services.ProductService;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(ProductController.PATH)
public class ProductController extends CMSController {

  protected static final String PATH = BASE_V1 + "/products";

  public ProductController(ProductService service) {
    super(service);
  }
}
