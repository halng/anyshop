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

  setApiToken(jsonObject: any) {
    const expiredTime = new Date().getTime() + 3600000; // 1 hour

    localStorage.setItem(this.authKey, JSON.stringify(jsonObject));
    localStorage.setItem(this.authExpireKey, expiredTime.toString());
  }

  isLogin() {
    // check if need to refresh token
    const expiredTime = localStorage.getItem(this.authExpireKey);
    if (expiredTime) {
      const expiredTimeInt = parseInt(expiredTime, 10);
      if (expiredTimeInt < new Date().getTime()) {
        localStorage.removeItem(this.authKey);
        localStorage.removeItem(this.authExpireKey);
        return false;
      } else {
        return !!localStorage.getItem(this.authKey);
      }
    }
    return false;
  }

  logout() {
    localStorage.removeItem(this.authKey);
    localStorage.removeItem(this.authExpireKey);
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
        'X-API-SECRET-TOKEN': `${token}`,
        'X-API-USER-ID': id,
      },
    });
  }
}
