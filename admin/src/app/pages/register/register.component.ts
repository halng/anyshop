import { UserService } from './../../services/user.service';
import { NgStyle } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import {
  ButtonDirective,
  CardBodyComponent,
  CardComponent,
  CardGroupComponent,
  ColComponent,
  ContainerComponent,
  FormControlDirective,
  FormDirective,
  InputGroupComponent,
  InputGroupTextDirective,
  RowComponent,
  TextColorDirective,
} from '@coreui/angular';
import { IconDirective } from '@coreui/icons-angular';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    FormsModule,
    ContainerComponent,
    RowComponent,
    ColComponent,
    CardGroupComponent,
    TextColorDirective,
    CardComponent,
    CardBodyComponent,
    FormDirective,
    InputGroupComponent,
    InputGroupTextDirective,
    IconDirective,
    FormControlDirective,
    ButtonDirective,
    NgStyle,
  ],
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  // Form fields

  email = '';
  username = '';
  password = '';
  confirmPassword = '';

  constructor(
    private readonly userService: UserService,
    private readonly toast: ToastrService,
    private readonly router: Router
  ) {}

  onSubmitForm(event: MouseEvent) {
    if (!this.isFormValid()) {
      this.toast.error('Please fill in all fields correctly.');
      return;
    }
    this.userService
      .register(this.username, this.email, this.password, this.confirmPassword)
      .subscribe({
        next: () => {
          this.toast.success('Registration successful!');
          this.router.navigate(['/login']);
        },
        error: (error: any) => {
          console.error('Registration error:', error);
          if (error.status === 400) {
            this.toast.error('Registration failed. Please check your input.');
            this.toast.error(error.error.error);
          }
        },
      });
    event.stopPropagation();
  }

  isFormValid(): boolean {
    return (
      !!this.username &&
      !!this.password &&
      !!this.confirmPassword &&
      !!this.email &&
      this.password === this.confirmPassword &&
      this.isValidEmail(this.email) &&
      this.isValidPassword(this.password)
    );
  }

  isValidEmail(email: string): boolean {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailPattern.test(email);
  }

  isValidPassword(password: string): boolean {
    const passwordPattern =
      /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/;
    return passwordPattern.test(password);
  }
}
