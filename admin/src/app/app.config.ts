/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import {
  ApplicationConfig,
  provideZoneChangeDetection,
  importProvidersFrom,
} from '@angular/core';

import { routes } from './app.routes';
import {
  provideRouter,
  withEnabledBlockingInitialNavigation,
  withHashLocation,
  withInMemoryScrolling,
  withRouterConfig,
  withViewTransitions,
} from '@angular/router';
import { provideClientHydration } from '@angular/platform-browser';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { DropdownModule, SidebarModule } from '@coreui/angular';
// import { provideAnimations } from '@angular/platform-browser/animations';
// import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { IconSetService } from '@coreui/icons-angular';
import { provideToastr } from 'ngx-toastr';
import { AuthInterceptor } from '../interceptor';
import { provideStore } from '@ngrx/store';
import { authReducer } from '../store/auth/auth.reducer';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }),
    provideRouter(
      routes,
      withRouterConfig({
        onSameUrlNavigation: 'reload',
      }),
      withInMemoryScrolling({
        scrollPositionRestoration: 'top',
        anchorScrolling: 'enabled',
      }),
      withEnabledBlockingInitialNavigation(),
      withViewTransitions(),
      withHashLocation()
    ),
    provideClientHydration(),
    // provideAnimations()
    provideAnimationsAsync(),
    provideToastr(
      {
        timeOut: 5000,
        positionClass: 'toast-top-right',
        preventDuplicates: true,
      }
    ),
    provideStore({
      auth: authReducer
    }),
    provideHttpClient(withInterceptors([
      AuthInterceptor
    ])),
    IconSetService,
    importProvidersFrom(SidebarModule, DropdownModule),
  ],
};
