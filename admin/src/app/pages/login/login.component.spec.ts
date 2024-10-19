import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { of, throwError } from 'rxjs';

import { LoginComponent } from './login.component';
import { UserService } from '../../services/user.service';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let userService: UserService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        LoginComponent,
        ReactiveFormsModule,
        MatButtonModule,
        MatFormFieldModule,
        MatIconModule,
        MatInputModule,
        FormsModule,
      ],
      providers: [
        UserService,
        provideHttpClient(),
        provideHttpClientTesting(),
        provideAnimationsAsync(),
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    userService = TestBed.inject(UserService);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize with default values', () => {
    expect(component.hide()).toBeTrue();
    expect(component.username).toBe('');
    expect(component.password).toBe('');
    expect(component.error).toBe('');
  });

  it('should toggle password visibility', () => {
    component.onClickHidePassword(new MouseEvent('click'));
    expect(component.hide()).toBeFalse();
    component.onClickHidePassword(new MouseEvent('click'));
    expect(component.hide()).toBeTrue();
  });

  it('should show error for invalid username and password', () => {
    component.username = 'user';
    component.password = 'pass';
    component.onSubmitForm(new MouseEvent('click'));
    expect(component.error).toBe('Username or password is incorrect format');
    expect(component.username).toBe('');
    expect(component.password).toBe('');
  });

  it('should call userService.login on valid form submission', () => {
    spyOn(userService, 'login').and.returnValue(of({ data: { 'api-token': 'token' } }));
    spyOn(userService, 'setApiToken');

    component.username = 'validUser';
    component.password = 'validPassword';
    component.onSubmitForm(new MouseEvent('click'));

    expect(userService.login).toHaveBeenCalledWith('validUser', 'validPassword');
    expect(userService.setApiToken).toHaveBeenCalledWith('token');
  });

  it('should handle login error', () => {
    spyOn(userService, 'login').and.returnValue(throwError({ error: { error: 'Invalid credentials' } }));

    component.username = 'validUser';
    component.password = 'validPassword';
    component.onSubmitForm(new MouseEvent('click'));

    expect(component.error).toBe('Invalid credentials');
  });

  it('should handle missing api-token in response', () => {
    spyOn(userService, 'login').and.returnValue(of({ data: {} }));

    component.username = 'validUser';
    component.password = 'validPassword';
    component.onSubmitForm(new MouseEvent('click'));

    expect(component.error).toBe('There was an error while login. Please try again');
  });
});