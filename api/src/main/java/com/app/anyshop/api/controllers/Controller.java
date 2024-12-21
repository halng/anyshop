/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.controllers;

import com.app.anyshop.api.annotation.ValidAction;
import com.app.anyshop.api.viewmodel.response.ResVM;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import jakarta.validation.Valid;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestParam;

public interface Controller {
  String BASE_V1 = "/api/v1";

  @Operation(summary = "Get all objects")
  @ApiResponses(
      value = {
        @ApiResponse(responseCode = "200", description = "Successfully retrieved objects"),
        @ApiResponse(responseCode = "400", description = "Invalid request"),
        @ApiResponse(responseCode = "500", description = "Internal server error")
      })
  ResponseEntity<ResVM> getAll(@Min(1) int page);

  @Operation(summary = "Get object details by ID")
  @ApiResponses(
      value = {
        @ApiResponse(responseCode = "200", description = "Successfully retrieved object details"),
        @ApiResponse(responseCode = "400", description = "Invalid request"),
        @ApiResponse(responseCode = "404", description = "Object not found"),
        @ApiResponse(responseCode = "500", description = "Internal server error")
      })
  ResponseEntity<ResVM> getById(@NotBlank String id);

  @Operation(summary = "Create a new objects")
  @ApiResponses(
      value = {
        @ApiResponse(responseCode = "201", description = "Object created successfully"),
        @ApiResponse(responseCode = "400", description = "Invalid request"),
        @ApiResponse(responseCode = "500", description = "Internal server error")
      })
  ResponseEntity<ResVM> create(@Valid Object obj);

  @Operation(summary = "Update a object by ID")
  @ApiResponses(
      value = {
        @ApiResponse(responseCode = "200", description = "Object updated successfully"),
        @ApiResponse(responseCode = "400", description = "Invalid request"),
        @ApiResponse(responseCode = "404", description = "Object not found"),
        @ApiResponse(responseCode = "500", description = "Internal server error")
      })
  ResponseEntity<ResVM> update(@NotBlank String id, @Valid Object obj);

  @Operation(summary = "Update the status of a object by ID")
  @ApiResponses(
      value = {
        @ApiResponse(responseCode = "200", description = "Object status updated successfully"),
        @ApiResponse(responseCode = "400", description = "Invalid request"),
        @ApiResponse(responseCode = "404", description = "Object not found"),
        @ApiResponse(responseCode = "500", description = "Internal server error")
      })
  ResponseEntity<ResVM> updateStatus(@NotBlank String id, @ValidAction String action);
}
