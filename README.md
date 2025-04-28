# JWT-Verify

## 项目简介

JWT-Verify 是一个基于 JSON Web Token (JWT) 的认证和授权示例项目。它包含后端服务（使用 Go 编写）和前端服务（使用 JavaScript 和 Axios 实现），展示了如何生成、验证和使用 JWT 进行用户认证。

### 功能特性

- **JWT 生成**：通过用户名生成带有过期时间的 JWT。
- **JWT 验证**：验证客户端发送的 JWT 是否有效。
- **前端集成**：通过 Axios 实现前端与后端的交互，支持自动附加 JWT 到请求头。
- **中间件支持**：使用 Gin 框架实现 JWT 验证中间件。

---

## 使用方法

### 环境要求

- **后端**：需要安装 Go 1.23.5 或更高版本。
- **前端**：需要安装 Node.js 和 npm。
- **依赖**：确保安装了 `github.com/golang-jwt/jwt/v5` 和 `github.com/gin-gonic/gin` 等依赖。

### 后端运行

1. **克隆项目**：

   ```bash
   git clone https://github.com/Ashside/JWT-Verify.git
   cd JWT-Verify
   ```

2. **配置密钥**：
   - 在 `verify/jwt.go` 中设置 `secretKey` 或 `filePath`。
   - 在 `filePath` 文件中存储用于签名的密钥。

3. **运行后端服务**：

- 将 `verify` 目录下的 `jwt.go` 文件加入你的代码
- 为需要使用 JWT 的路由添加中间件，例如：

```go
authGroup.Use(middleware.AuthMiddleware())/ 添加 JWT 中间件
```

- 启动服务：

```bash
go run main.go
```

### 前端运行

1. **添加文件**：
   - 将 `web` 目录下的 `src` 文件夹复制到你的前端项目中。
   - 在 `web/src/services/axios.js` 中添加 `baseURL`，指向后端服务地址。
   - 在 `web/src/services/auth.js` 中修改你的登录、登出和 Token 管理逻辑。

2. **安装依赖**：

   ```bash
   cd web
   npm install
   ```

3. **启动前端服务**：

   ```bash
   npm run dev
   ```

## 开发指南

### 后端开发

1. **JWT 配置**：
   - 修改 `verify/jwt.go` 中的 `secretKey` 或 `filePath` 以更改密钥配置。
   - 调整 `tokenExpiration` 以更改 Token 的过期时间。

2. **中间件**：
   - 在 `middleware/auth.go` 中自定义 JWT 验证逻辑。

3. **依赖管理**：
   - 使用 `go mod tidy` 管理依赖。

### 前端开发

1. **Axios 配置**：
   - 修改 `web/src/services/axios.js` 中的 `baseURL` 以适配后端服务地址。

2. **认证服务**：
   - 在 `web/src/services/auth.js` 中扩展登录、登出和 Token 管理逻辑。

3. **调试**：
   - 使用浏览器开发者工具调试前端代码。

---

## 许可证

本项目基于 MIT 许可证，详情请参阅 [LICENSE](./LICENSE) 文件。

## Acknowledgments

- 感谢 Github Copilot 和 ChatGPT 的帮助。