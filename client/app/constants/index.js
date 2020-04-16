export const API_DOMAIN = !process.env.NODE_ENV || process.env.NODE_ENV === 'development'
    ? 'http://127.0.0.1:3001'
    : 'http://127.0.0.1:3001';
export const SESSION_TOKEN_KEY = 'token';
export const USER_KEY = '__buk__identifier';

