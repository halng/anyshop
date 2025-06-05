/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'view',
    loadComponent: () =>
      import('./main/main.component').then((m) => m.MainComponent),
  },
  {
    path: "my-tasks",
    loadComponent: () =>
      import("./my-tasks/my-tasks.component").then((m) => m.MyTasksComponent),
  }
];
