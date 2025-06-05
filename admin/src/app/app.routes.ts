/*
 * *****************************************************************************************
 * Copyright 2024 By ANYSHOP Project
 * Licensed under the Apache License, Version 2.0;
 * *****************************************************************************************
 */

import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';
import { LoginComponent } from './pages/login/login.component';
// import { HomeComponent } from './pages/home/home.component';
import { LayoutComponent } from './pages/layout/layout.component';

export const routes: Routes = [
  { path: 'login', component: LoginComponent },
  {
    path: 'register',
    loadComponent: () =>
      import('./pages/register/register.component').then(
        (m) => m.RegisterComponent
      ),
  },
  {
    path: 'activate',
    loadComponent: () =>
      import('./pages/activate/activate.component').then(
        (m) => m.ActivateComponent
      ),
  },
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: 'staff-management',
        loadChildren: () =>
          import('./pages/staff-management/staff-management.routes').then(
            (m) => m.routes
          ),
      },
      {
        path: 'home',
        loadChildren: () =>
          import('./pages/home/home.routes').then((m) => m.routes),
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
