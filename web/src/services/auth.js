import axiosInstance from './axios';

class AuthService {
  login(username, password) {
    return axiosInstance
      .post('/api/auth/login', {
        username,
        password
      })
      .then(response => {
        debugger;
        if (response.data.token) {
          localStorage.setItem('user', JSON.stringify(response.data));
          localStorage.setItem('username', username);
        }
        return response.data;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  getCurrentUser() {
    const userStr = localStorage.getItem('user');
    if (userStr) {
      return JSON.parse(userStr);
    }
    return null;
  }

  getToken() {
    const user = this.getCurrentUser();
    if (user && user.token) {
      return user.token;
    }
    return null;
  }

  isAuthenticated() {
    return this.getToken() !== null;
  }
}

export default new AuthService(); 