import { Login, Register } from "@/types/auth";
import api from ".";

export const AuthService = {
  login(username: string, password: string) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    const body: Login = {
      password,
    };
    if (emailRegex.test(username)) {
      body["email"] = username;
    } else {
      body["username"] = username;
    }
    return api.post("/iam/login", body);
  },
  register(data: Register) {
    const body = {
      username: data.username,
      password: data.password,
      email: data.email,
      confirm_password: data.confirmPassword,
    };
    return api.post("/iam/register", body);
  },
};
