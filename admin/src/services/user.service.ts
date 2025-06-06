/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {environment} from '../environments/environment';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import * as AuthSelectors from '../store/auth/auth.selector'
import * as AuthActions from '../store/auth/auth.action'

@Injectable({
  providedIn: 'root',
})
export class UserService {
  HOST = environment.HOST;
  BASE_URL = `${this.HOST}/api/v1/iam`;
  isAuthenticated$: Observable<boolean>;
  token$: Observable<string | null>;
  username$: Observable<string | null>;

  constructor(private store: Store, private http: HttpClient) {
   this.isAuthenticated$ = this.store.select(AuthSelectors.selectIsAuthenticated);
    this.token$ = this.store.select(AuthSelectors.selectToken)
    this.username$ = this.store.select(AuthSelectors.selectUsername)    
  }


  login(username: string, password: string): any {
    let object = {}
    if (username.indexOf('@') > -1) {
      object = { 'email': username, password }
    }
    else {
      object = { username, password }
    }
    this.http.post(`${this.BASE_URL}/login`, object).subscribe({
      next: (response: any) => {
        if (response && response.data && response.data.token) {
          this.store.dispatch(AuthActions.login({
            token: response.data.token,
            username: response.data.username,
            userId: response.data.id,
          }))
        }
        return response
      },
      error: (error) => {
        console.error('Error fetching shops:', error);
        
      },
    });
  }

  register(username: string, email: string, password: string, confirmPassword: string) {
    return this.http.post(`${this.BASE_URL}/register`, {
      username,
      email,
      password,
      'confirm_password': confirmPassword,
    });
  }

  activate(username: string, token: string) {
    return this.http.post(`${this.BASE_URL}/activate?username=${username}&token=${token}`, {});
  }

  logout() {
    this.store.dispatch(AuthActions.logout())
    this.store.dispatch(AuthActions.clearToken())
  }

  getAllShops() {
    return this.http.get(`${this.BASE_URL}/shops`);
  }

  createNewShop(name: string, domain: string) {
    return this.http.post(`${this.BASE_URL}/shops`, {
      name,
      slug: domain,
    });
  }
}
