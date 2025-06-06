/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/


export interface AuthState {
    token: string | null;
    isAuthenticated: boolean;
    username: string | null;
    userId: string | null;
}