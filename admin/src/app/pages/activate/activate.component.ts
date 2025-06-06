/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/

import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { UserService } from '../../../services/user.service';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-activate',
  standalone: true,
  imports: [NgIf],
  templateUrl: './activate.component.html',
})
export class ActivateComponent implements OnInit {
  errorMessage: string | null = null;

  constructor(
    private readonly userService: UserService,
    private readonly router: Router,
    private readonly route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    const username = this.route.snapshot.queryParamMap.get('username');
    const token = this.route.snapshot.queryParamMap.get('token');

    // Optionally remove query parameters from the URL
    this.router.navigate([], {
      queryParams: {},
      replaceUrl: true,
    });

    if (username && token) {
      this.userService.activate(username, token).subscribe({
        next: () => {
          this.router.navigate(['/login']);
        },
        error: (err: any) => {
          // console.error('Activation error:', err);
          this.errorMessage = err.error.error;
        },
      });
    } else {
      this.errorMessage = 'Missing verification details in URL.';
    }
  }
}
