/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { createFeatureSelector, createSelector } from "@ngrx/store";
import { AuthState } from "../../models/auth.model";

export const selectAuthState = createFeatureSelector<AuthState>("auth");


export const selectToken = createSelector(
    selectAuthState,
    (state) => state.token
)

export const selectIsAuthenticated = createSelector(
    selectAuthState,
    (state) => state.isAuthenticated
)

export const selectUsername = createSelector(
    selectAuthState,
    (state) => state.username
)

export const selectUserId = createSelector(
    selectAuthState,
    (state) => state.userId
)