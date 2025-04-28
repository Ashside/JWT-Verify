import axios from 'axios';
import authService from './auth';

const instance = axios.create({
  baseURL: 'http://localhost:8080/', // 根据实际后端地址修改
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器
instance.interceptors.request.use(
  config => {
    const token = authService.getToken();
    if (token) {
      // 确保使用正确的 Authorization 头部格式
      config.headers['Authorization'] = `Bearer ${token}`;
    } else {
      // 如果没有token，可以在这里处理，比如跳转到登录页
      console.warn('No authentication token found');
      window.location.href = '/login';
      
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // 未授权，清除用户信息并跳转到登录页
          authService.logout();
          // 清除输入框中的用户名和密码
          document.getElementById('username').value = '';
          document.getElementById('password').value = '';
          // 跳转到登录页
          console.warn('Unauthorized, redirecting to login');
          window.location.href = '/welcome';
          break;
        case 403:
          // 权限不足
          console.error('Access denied');
          break;
        case 500:
          // 服务器错误
          console.error('Server error');
          break;
        default:
          console.error('Request error:', error.response.status);
      }
    }
    return Promise.reject(error);
  }
);

export default instance; 