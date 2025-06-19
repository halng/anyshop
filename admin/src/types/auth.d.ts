export interface Login  {
    username?: string;
    password: string;
    email?: string;
}

export interface Register {
    username: string;
    password: string;
    email: string;
    confirmPassword: string;
}