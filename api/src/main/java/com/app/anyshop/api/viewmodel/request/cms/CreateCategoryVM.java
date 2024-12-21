/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.viewmodel.request.cms;

import jakarta.validation.constraints.NotBlank;

public record CreateCategoryVM(@NotBlank String name, @NotBlank String description) {
}
