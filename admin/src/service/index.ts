import axios, { AxiosInstance, InternalAxiosRequestConfig } from "axios";
import { AuthService } from "./auth";
const API_BASE_URL = "http://localhost:9000/api/v1";
axios.defaults.baseURL = API_BASE_URL;

const INGNORE_REQUEST_URLS = [
    "login",
    "register",
    "refresh",
    "logout",
    "verify",
    "forgot-password",]

const api: AxiosInstance = axios.create({
    baseURL: API_BASE_URL,
    timeout: 10000, // 10 seconds timeout
});

// interceptors
api.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        const shouldIgnore = INGNORE_REQUEST_URLS.some((url) =>
            config.url?.includes(url)
        );

        if (!shouldIgnore) {
            const token = localStorage.getItem("token");
            if (token) {
                config.headers["Authorization"] = `Bearer ${token}`;
            }
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (response) => {
        return response.data;
    },
    (error) => {
        if (error.response) {
            // Handle specific status codes
            switch (error.response.status) {
                case 401:
                    console.error("Unauthorized access - redirecting to login");
                    // Redirect to login or handle unauthorized access
                    break;
                case 403:
                    console.error("Forbidden access");
                    break;
                case 404:
                    console.error("Resource not found");
                    break;
                default:
                    console.error("An error occurred:", error.response.data);
            }
        } else {
            console.error("Network error or server is down");
        }
        return Promise.reject(error);
    }
);


export default api;
export { AuthService };