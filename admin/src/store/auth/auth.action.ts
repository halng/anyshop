/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import {createAction, props} from '@ngrx/store'

export const login = createAction(
    "[Auth] Login",
    props<{username: string, token: string, userId: string}>()
);

export const logout = createAction(
    "[Auth] Logout"
)

export const clearToken = createAction(
    "[Auth] Clear Token"
)