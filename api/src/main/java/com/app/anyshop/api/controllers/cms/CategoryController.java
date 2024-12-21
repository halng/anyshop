/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.controllers.cms;

import com.app.anyshop.api.viewmodel.response.ResVM;
import jakarta.validation.constraints.NotBlank;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
public class CategoryController implements CMSController {

  @RequestMapping(
      path = CMS_PATH + "/categories",
      method = RequestMethod.GET)
  @Override
  public ResponseEntity<ResVM> getAll(int page) {
    return ResponseEntity.ok(ResVM.builder().code(HttpStatus.OK).msg("getAll").build());
  }

  @RequestMapping(
      path = CMS_PATH + "/categories/{id}",
      method = RequestMethod.GET)
  @Override
  public ResponseEntity<ResVM> getById(@PathVariable String id) {
    return ResponseEntity.ok(ResVM.builder().code(HttpStatus.OK).msg("getById").build());
  }

  @RequestMapping(
      path = CMS_PATH + "/categories",
      method = RequestMethod.POST)
  @Override
  public ResponseEntity<ResVM> create(Object obj) {
    return ResponseEntity.ok(ResVM.builder().code(HttpStatus.CREATED).msg("create").build());
  }

  @RequestMapping(
      path = CMS_PATH + "/categories/{id}",
      method = RequestMethod.PUT)
  @Override
  public ResponseEntity<ResVM> update(@PathVariable String id, @RequestBody Object obj) {
    return ResponseEntity.ok(ResVM.builder().code(HttpStatus.ACCEPTED).msg("update").build());
  }

  @RequestMapping(path = CMS_PATH + "/categories", method = RequestMethod.PATCH)
  @Override
  public ResponseEntity<ResVM> updateStatus(@RequestParam String id, @RequestParam String action) {
    return ResponseEntity.ok(ResVM.builder().code(HttpStatus.ACCEPTED).msg("updateStatus").build());
  }
}
