/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/
// auth.interceptor.ts
import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { Router } from '@angular/router';
import { catchError, switchMap, take } from 'rxjs/operators';
import { selectToken } from './store/auth/auth.selector';
import * as AuthActions from './store/auth/auth.action';

export const AuthInterceptor: HttpInterceptorFn = (req, next) => {
  const store = inject(Store);
  const router = inject(Router);

  return store.select(selectToken).pipe(
    take(1),
    switchMap(token => {
      if (token) {
        req = req.clone({
          setHeaders: {
            Authorization: `Bearer ${token}`
          }
        });
      }

      return next(req).pipe(
        catchError((error) => {
          if (error.status === 401) {
            store.dispatch(AuthActions.logout());
            store.dispatch(AuthActions.clearToken());
            router.navigate(['/login']);
          }
          throw error;
        })
      );
    })
  );
};
