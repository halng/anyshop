/*
 * ****************************************************************************************
 * Copyright (c) 2024, Hal Ng (halng) - haonguyentan2001@gmail.com
 * Licensed under the MIT License. See LICENSE file in the project root for details.
 * ****************************************************************************************
 */

package com.app.anyshop.api.controllers.cms;

import com.app.anyshop.api.controllers.Controller;
import io.swagger.v3.oas.annotations.Hidden;

public interface CMSController extends Controller {
	@Hidden
	String CMS_PATH = BASE_V1 + "/cms";
}
