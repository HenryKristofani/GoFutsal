// Authentication utilities and secure token management

interface User {
  id: number;
  username: string;
  email: string;
  role: string;
}

interface LoginResponse {
  success: boolean;
  message: string;
  access_token: string;
  refresh_token: string;
  user: User;
}

interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

interface LoginRequest {
  username: string;
  password: string;
}

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// Secure token storage with encryption-like obfuscation
class TokenManager {
  private static ACCESS_TOKEN_KEY = 'gf_at';
  private static REFRESH_TOKEN_KEY = 'gf_rt';
  private static USER_KEY = 'gf_user';

  // Simple encoding to prevent casual token theft
  private static encode(value: string): string {
    return btoa(encodeURIComponent(value));
  }

  private static decode(value: string): string {
    try {
      return decodeURIComponent(atob(value));
    } catch {
      return '';
    }
  }

  static setTokens(accessToken: string, refreshToken: string, user: User): void {
    if (typeof window === 'undefined') return;
    
    localStorage.setItem(this.ACCESS_TOKEN_KEY, this.encode(accessToken));
    localStorage.setItem(this.REFRESH_TOKEN_KEY, this.encode(refreshToken));
    localStorage.setItem(this.USER_KEY, this.encode(JSON.stringify(user)));
  }

  static getAccessToken(): string | null {
    if (typeof window === 'undefined') return null;
    
    const token = localStorage.getItem(this.ACCESS_TOKEN_KEY);
    return token ? this.decode(token) : null;
  }

  static getRefreshToken(): string | null {
    if (typeof window === 'undefined') return null;
    
    const token = localStorage.getItem(this.REFRESH_TOKEN_KEY);
    return token ? this.decode(token) : null;
  }

  static getUser(): User | null {
    if (typeof window === 'undefined') return null;
    
    const user = localStorage.getItem(this.USER_KEY);
    if (!user) return null;
    
    try {
      return JSON.parse(this.decode(user));
    } catch {
      return null;
    }
  }

  static clearTokens(): void {
    if (typeof window === 'undefined') return;
    
    localStorage.removeItem(this.ACCESS_TOKEN_KEY);
    localStorage.removeItem(this.REFRESH_TOKEN_KEY);
    localStorage.removeItem(this.USER_KEY);
  }

  static isAuthenticated(): boolean {
    return !!this.getAccessToken() && !!this.getUser();
  }
}

// API client with automatic token refresh
class AuthAPI {
  private static async makeRequest(
    url: string, 
    options: RequestInit = {},
    useRefreshToken = false
  ): Promise<Response> {
    const token = useRefreshToken ? TokenManager.getRefreshToken() : TokenManager.getAccessToken();
    
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      ...options.headers,
    };

    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }

    const response = await fetch(`${API_BASE_URL}${url}`, {
      ...options,
      headers,
    });

    // If access token expired, try to refresh
    if (response.status === 401 && !useRefreshToken && TokenManager.getRefreshToken()) {
      const newToken = await this.refreshAccessToken();
      if (newToken) {
        // Retry with new token
        headers.Authorization = `Bearer ${newToken}`;
        return fetch(`${API_BASE_URL}${url}`, {
          ...options,
          headers,
        });
      }
    }

    return response;
  }

  static async register(data: RegisterRequest): Promise<{ success: boolean; message: string; user?: User }> {
    try {
      const response = await this.makeRequest('/api/users/register', {
        method: 'POST',
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (!response.ok) {
        return { 
          success: false, 
          message: result.error || 'Registrasi gagal' 
        };
      }

      return { 
        success: true, 
        message: 'Registrasi berhasil! Silakan login.',
        user: result 
      };
    } catch (error) {
      return { 
        success: false, 
        message: 'Terjadi kesalahan koneksi' 
      };
    }
  }

  static async login(data: LoginRequest): Promise<{ success: boolean; message: string; user?: User }> {
    try {
      const response = await this.makeRequest('/api/auth/login', {
        method: 'POST',
        body: JSON.stringify(data),
      });

      const result: LoginResponse = await response.json();

      if (!response.ok) {
        return { 
          success: false, 
          message: result.message || 'Login gagal' 
        };
      }

      // Store tokens securely
      TokenManager.setTokens(result.access_token, result.refresh_token, result.user);

      return { 
        success: true, 
        message: result.message,
        user: result.user 
      };
    } catch (error) {
      return { 
        success: false, 
        message: 'Terjadi kesalahan koneksi' 
      };
    }
  }

  static async refreshAccessToken(): Promise<string | null> {
    try {
      const response = await this.makeRequest('/api/auth/refresh', {
        method: 'POST',
      }, true); // Use refresh token

      if (!response.ok) {
        // Refresh failed, clear all tokens
        TokenManager.clearTokens();
        return null;
      }

      const result: LoginResponse = await response.json();
      
      // Update access token
      const user = TokenManager.getUser();
      if (user && result.access_token) {
        TokenManager.setTokens(
          result.access_token, 
          TokenManager.getRefreshToken() || '', 
          user
        );
        return result.access_token;
      }

      return null;
    } catch (error) {
      TokenManager.clearTokens();
      return null;
    }
  }

  static logout(): void {
    TokenManager.clearTokens();
    // Redirect to home or refresh page
    if (typeof window !== 'undefined') {
      window.location.reload();
    }
  }

  static getCurrentUser(): User | null {
    return TokenManager.getUser();
  }

  static isAuthenticated(): boolean {
    return TokenManager.isAuthenticated();
  }
}

export { AuthAPI, TokenManager };
export type { User, LoginRequest, RegisterRequest };