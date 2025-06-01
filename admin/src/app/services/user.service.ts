/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { UserCreate } from '../types';
import {environment} from '../../environments/environment';
@Injectable({
  providedIn: 'root',
})
export class UserService {
  HOST = environment.HOST;
  BASE_URL = `${this.HOST}/api/v1/iam`;
  authKey = environment.LOCAL_STORAGE.AUTH_KEY;
  authExpireKey = environment.LOCAL_STORAGE.AUTH_EXPIRE_KEY;

  constructor(private http: HttpClient) {}

  login(username: string, password: string) {
    return this.http.post(`${this.BASE_URL}/login`, { username, password });
  }

  register(username: string, email: string, password: string, confirmPassword: string) {
    return this.http.post(`${this.BASE_URL}/register`, {
      username,
      email,
      password,
      'confirm_password': confirmPassword,
    });
  }

  setApiToken(jsonObject: any) {
    const expiredTime = new Date().getTime() + 3600000; // 1 hour


    global.localStorage.setItem(this.authKey, JSON.stringify(jsonObject));
    global.localStorage.setItem(this.authExpireKey, expiredTime.toString());
  }

  isLogin() {
    // check if need to refresh token
    const expiredTime = global.localStorage.getItem(this.authExpireKey);
    if (expiredTime) {
      const expiredTimeInt = parseInt(expiredTime, 10);
      if (expiredTimeInt < new Date().getTime()) {
        global.localStorage.removeItem(this.authKey);
        global.localStorage.removeItem(this.authExpireKey);
        return false;
      } else {
        return !!global.localStorage.getItem(this.authKey);
      }
    }
    return false;
  }

  logout() {
    global.localStorage.removeItem(this.authKey);
    global.localStorage.removeItem(this.authExpireKey);
  }

  createStaff(data: UserCreate) {
    const raw = localStorage.getItem(this.authKey);
    if (!raw) {
      throw new Error('No auth token found');
    }
    const authObject = JSON.parse(raw);
    const token = authObject['api-token'];
    const id = authObject['id'];

    return this.http.post(`${this.BASE_URL}/create-staff`, data, {
      headers: {
        'X-API-SECRET-TOKEN': token,
        'X-API-USER-ID': id,
      },
    });
  }
}
