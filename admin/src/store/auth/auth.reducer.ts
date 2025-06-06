/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { createReducer, on } from "@ngrx/store";
import { AuthState } from "../../models/auth.model";
import * as AuthActions from "./auth.action"

export const initialState: AuthState = {
    token: null,
    isAuthenticated: false,
    username: null,
    userId: null
}

export const authReducer = createReducer(
    initialState,
    on(AuthActions.login, (state, { username, token, userId }) => (
        {
            ...state,
            token,
            isAuthenticated: true,
            username,
            userId
        }
    )),
    on(AuthActions.logout, (state) => (
        {
            ...state,
            token: null,
            isAuthenticated: false,
            username: null,
            userId: null
        }
    )),
    on(AuthActions.clearToken, (state) => (
        {
            ...state,
            token: null,
            isAuthenticated: false
        }
    ))
);