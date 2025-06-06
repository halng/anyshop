/*
* *****************************************************************************************
* Copyright 2024 By ANYSHOP Project 
* Licensed under the Apache License, Version 2.0;
* *****************************************************************************************
*/


import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { UserService } from './user.service';
import { Store } from '@ngrx/store';
import { MockStore, provideMockStore } from '@ngrx/store/testing';
import { environment } from '../environments/environment';
import * as AuthActions from '../store/auth/auth.action';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;
  let store: MockStore;
  const initialState = {
    auth: {
      isAuthenticated: false,
      token: null,
      username: null,
      userId: null
    }
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [
        UserService,
        provideMockStore({ initialState })
      ]
    });

    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
    store = TestBed.inject(Store) as MockStore;
    spyOn(store, 'dispatch').and.callThrough();
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('login', () => {
    it('should login with username', () => {
      const mockResponse = {
        data: {
          token: 'test-token',
          username: 'testuser',
          id: '123'
        }
      };

      service.login('testuser', 'password');

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/login`);
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({
        username: 'testuser',
        password: 'password'
      });

      req.flush(mockResponse);

      expect(store.dispatch).toHaveBeenCalledWith(
        AuthActions.login({
          token: mockResponse.data.token,
          username: mockResponse.data.username,
          userId: mockResponse.data.id
        })
      );
    });

    it('should login with email', () => {
      const mockResponse = {
        data: {
          token: 'test-token',
          username: 'testuser',
          id: '123'
        }
      };

      service.login('test@example.com', 'password');

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/login`);
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({
        email: 'test@example.com',
        password: 'password'
      });

      req.flush(mockResponse);
    });

    it('should handle login error', () => {
      spyOn(console, 'error');
      
      service.login('testuser', 'password');

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/login`);
      req.error(new ErrorEvent('Network error'));

      expect(console.error).toHaveBeenCalled();
    });
  });

  describe('register', () => {
    it('should register a new user', () => {
      const mockResponse = { success: true };

      service.register('testuser', 'test@example.com', 'password', 'password')
        .subscribe(response => {
          expect(response).toEqual(mockResponse);
        });

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/register`);
      expect(req.request.method).toBe('POST');
      expect(req.request.body).toEqual({
        username: 'testuser',
        email: 'test@example.com',
        password: 'password',
        confirm_password: 'password'
      });

      req.flush(mockResponse);
    });
  });

  describe('activate', () => {
    it('should activate user account', () => {
      const mockResponse = { success: true };

      service.activate('testuser', 'activation-token')
        .subscribe(response => {
          expect(response).toEqual(mockResponse);
        });

      const req = httpMock.expectOne(
        `${environment.HOST}/api/v1/iam/activate?username=testuser&token=activation-token`
      );
      expect(req.request.method).toBe('POST');

      req.flush(mockResponse);
    });
  });

  describe('logout', () => {
    it('should dispatch logout actions', () => {
      service.logout();

      expect(store.dispatch).toHaveBeenCalledWith(AuthActions.logout());
      expect(store.dispatch).toHaveBeenCalledWith(AuthActions.clearToken());
    });
  });

  describe('getAllShops', () => {
    it('should get all shops with auth token', () => {
      const mockResponse = { shops: [] };
      const authToken = 'test-token';

      service.getAllShops().subscribe(response => {
        expect(response).toEqual(mockResponse);
      });

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/shops`);
      expect(req.request.method).toBe('GET');
      expect(req.request.headers.get('Authorization')).toBe(`Bearer ${authToken}`);

      req.flush(mockResponse);
    });
  });

  describe('createNewShop', () => {
    it('should create a new shop', () => {
      const mockResponse = { success: true };
      const authToken = 'test-token';

      service.createNewShop('Test Shop', 'test-domain')
        .subscribe(response => {
          expect(response).toEqual(mockResponse);
        });

      const req = httpMock.expectOne(`${environment.HOST}/api/v1/iam/shops`);
      expect(req.request.method).toBe('POST');
      expect(req.request.headers.get('Authorization')).toBe(`Bearer ${authToken}`);
      expect(req.request.body).toEqual({
        name: 'Test Shop',
        slug: 'test-domain'
      });

      req.flush(mockResponse);
    });
  });

  describe('Store Selectors', () => {
    it('should select authentication state', (done) => {
      store.setState({
        auth: {
          isAuthenticated: true,
          token: 'test-token',
          username: 'testuser',
          userId: '123'
        }
      });

      service.isAuthenticated$.subscribe(isAuthenticated => {
        expect(isAuthenticated).toBe(true);
        done();
      });
    });

    it('should select token', (done) => {
      store.setState({
        auth: {
          isAuthenticated: true,
          token: 'test-token',
          username: 'testuser',
          userId: '123'
        }
      });

      service.token$.subscribe(token => {
        expect(token).toBe('test-token');
        done();
      });
    });

    it('should select username', (done) => {
      store.setState({
        auth: {
          isAuthenticated: true,
          token: 'test-token',
          username: 'testuser',
          userId: '123'
        }
      });

      service.username$.subscribe(username => {
        expect(username).toBe('testuser');
        done();
      });
    });
  });
});
