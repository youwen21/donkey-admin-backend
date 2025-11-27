# CORS 跨域资源共享配置说明

## 📋 快速对比：Access-Control-Allow-Headers vs Access-Control-Expose-Headers

| 对比项 | Access-Control-Allow-Headers | Access-Control-Expose-Headers |
|--------|------------------------------|-------------------------------|
| **作用方向** | 控制**请求头**（浏览器 → 服务器） | 控制**响应头**（服务器 → 浏览器） |
| **使用时机** | **OPTIONS 预检请求**时返回 | **实际请求响应**时返回 |
| **解决的问题** | "前端可以发送哪些请求头？" | "前端可以读取哪些响应头？" |
| **示例场景** | 允许发送 `Authorization`、`X-Custom-Header` | 允许读取 `X-Total-Count`、`X-Auth-Token` |
| **默认行为** | 简单请求头（如 `Content-Type: text/plain`）无需设置 | 部分标准响应头（如 `Content-Type`）默认可访问 |
| **设置位置** | 预检请求（OPTIONS）的响应中 | 实际请求的响应中 |

### 🎯 一句话总结
- **Access-Control-Allow-Headers**：告诉浏览器"你可以发送这些请求头给我"
- **Access-Control-Expose-Headers**：告诉浏览器"你可以让 JavaScript 读取这些响应头"

---

## CORS 响应头说明

### 1. Access-Control-Allow-Origin
**作用**：指定哪些源（域名）可以访问该资源

**常用值**：
- `*` - 允许所有源（不推荐，特别是使用 credentials 时）
- `https://example.com` - 允许特定域名
- `http://localhost:5173` - 允许本地开发域名

**注意**：
- 如果设置了 `Access-Control-Allow-Credentials: true`，则不能使用 `*`
- 必须明确指定允许的域名

**示例**：
```http
Access-Control-Allow-Origin: http://localhost:5173
Access-Control-Allow-Origin: https://admin.example.com
```

---

### 2. Access-Control-Allow-Methods
**作用**：指定允许的 HTTP 请求方法

**常用值**：
- `GET, POST, PUT, DELETE, PATCH, OPTIONS`
- `*` - 允许所有方法（不推荐）

**示例**：
```http
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, PATCH, OPTIONS
```

---

### 3. Access-Control-Allow-Credentials
**作用**：是否允许发送 Cookie 和认证信息

**值**：
- `true` - 允许携带凭证（Cookie、Authorization 等）
- `false` - 不允许（默认）

**注意**：
- 设置为 `true` 时，`Access-Control-Allow-Origin` 不能为 `*`
- 前端需要设置 `withCredentials: true` 或 `credentials: 'include'`

**示例**：
```http
Access-Control-Allow-Credentials: true
```

---

### 4. Access-Control-Allow-Headers
**作用**：指定哪些**请求头**可以被浏览器发送到服务器（用于预检请求）

**使用场景**：当浏览器发送"非简单请求"时，会先发送 OPTIONS 预检请求，询问服务器是否允许发送这些请求头

**简单请求 vs 非简单请求**：
- **简单请求**：只使用标准请求头（如 `Content-Type: text/plain`、`Content-Type: application/x-www-form-urlencoded`）
- **非简单请求**：使用自定义请求头（如 `Authorization`、`X-Custom-Header`）或 `Content-Type: application/json`

**常用值**：
- `Content-Type, Authorization, X-Requested-With`
- `*` - 允许所有请求头（不推荐，特别是使用 credentials 时）

**示例**：
```http
Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With, X-Custom-Header
```

**前端发送请求时**：
```js
// 如果请求头包含 Authorization，浏览器会先发送 OPTIONS 预检请求
fetch('http://api.example.com/users', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer token123',  // 自定义请求头，触发预检请求
    'X-Custom-Header': 'value'
  }
})
```

**后端响应预检请求时**：
```http
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: http://localhost:5173
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Headers: Content-Type, Authorization, X-Custom-Header  ← 告诉浏览器允许这些请求头
Access-Control-Allow-Credentials: true
```

---

### 5. Access-Control-Expose-Headers
**作用**：指定哪些**响应头**可以被前端 JavaScript 访问（用于实际响应）

**使用场景**：默认情况下，浏览器只允许前端访问部分响应头，自定义响应头需要显式暴露

**默认可访问的响应头**（无需暴露）：
- Cache-Control
- Content-Language
- Content-Type
- Expires
- Last-Modified
- Pragma

**需要暴露的自定义响应头**（如 `X-Total-Count`、`X-Auth-Token`）：
```http
Access-Control-Expose-Headers: X-Total-Count, X-Auth-Token, X-Request-Id
```

**前端访问响应头时**：
```js
const response = await fetch('http://api.example.com/users')
const totalCount = response.headers.get('X-Total-Count')  // 需要后端设置 Access-Control-Expose-Headers
const contentType = response.headers.get('Content-Type')   // 默认可访问，无需暴露
```

**后端响应实际请求时**：
```http
HTTP/1.1 200 OK
Content-Type: application/json
X-Total-Count: 100          ← 自定义响应头
X-Auth-Token: newtoken123   ← 自定义响应头
Access-Control-Expose-Headers: X-Total-Count, X-Auth-Token  ← 告诉浏览器允许前端访问这些响应头
```

---

## ⚠️ 重要区别对比

### Access-Control-Allow-Headers（请求头控制）
| 特性 | 说明 |
|------|------|
| **方向** | 控制**浏览器 → 服务器**的请求头 |
| **时机** | 在 **OPTIONS 预检请求**中返回 |
| **作用** | 告诉浏览器："你可以发送这些请求头给我" |
| **示例** | 允许前端发送 `Authorization`、`X-Custom-Header` 等 |

### Access-Control-Expose-Headers（响应头控制）
| 特性 | 说明 |
|------|------|
| **方向** | 控制**服务器 → 浏览器**的响应头 |
| **时机** | 在 **实际请求响应**中返回 |
| **作用** | 告诉浏览器："你可以让 JavaScript 访问这些响应头" |
| **示例** | 允许前端读取 `X-Total-Count`、`X-Auth-Token` 等 |

### 完整请求流程示例

```
1. 浏览器发送 OPTIONS 预检请求
   ↓
   Request Headers:
     Origin: http://localhost:5173
     Access-Control-Request-Method: POST
     Access-Control-Request-Headers: Authorization, Content-Type
   ↓
2. 服务器响应预检请求
   ↓
   Response Headers:
     Access-Control-Allow-Origin: http://localhost:5173
     Access-Control-Allow-Methods: GET, POST, PUT, DELETE
     Access-Control-Allow-Headers: Authorization, Content-Type  ← 允许这些请求头
     Access-Control-Allow-Credentials: true
   ↓
3. 浏览器发送实际请求
   ↓
   Request Headers:
     Origin: http://localhost:5173
     Authorization: Bearer token123
     Content-Type: application/json
   ↓
4. 服务器响应实际请求
   ↓
   Response Headers:
     Access-Control-Allow-Origin: http://localhost:5173
     Access-Control-Expose-Headers: X-Total-Count, X-Auth-Token  ← 允许访问这些响应头
     X-Total-Count: 100
     X-Auth-Token: newtoken123
     Content-Type: application/json
   ↓
5. 前端 JavaScript 可以访问暴露的响应头
   ↓
   const totalCount = response.headers.get('X-Total-Count')  // ✅ 可以访问
   const authToken = response.headers.get('X-Auth-Token')     // ✅ 可以访问
   const contentType = response.headers.get('Content-Type')   // ✅ 默认可访问
```

---

## 前端配置（Vite 开发环境）

### 方案一：使用 Vite 代理（推荐）

在 `vite.config.js` 中配置代理，避免跨域问题：

```js
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
        // 如果需要，可以设置请求头
        configure: (proxy, options) => {
          proxy.on('proxyReq', (proxyReq, req, res) => {
            // 可以在这里添加自定义请求头
          })
        }
      }
    }
  }
})
```

### 方案二：开发服务器设置 CORS 头（不推荐）

如果需要直接在 Vite 开发服务器设置 CORS 头：

```js
export default defineConfig({
  server: {
    cors: {
      origin: 'http://localhost:5173',
      credentials: true,
      methods: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'],
      allowedHeaders: ['Content-Type', 'Authorization'],
      exposedHeaders: ['X-Total-Count', 'X-Auth-Token']
    }
  }
})
```

---

## 后端配置示例

### Go (Gin) 示例

```go
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        
        // 允许的源列表
        allowedOrigins := []string{
            "http://localhost:5173",
            "https://admin.example.com",
        }
        
        // 检查源是否在允许列表中
        allowed := false
        for _, o := range allowedOrigins {
            if origin == o {
                allowed = true
                break
            }
        }
        
        if allowed {
            c.Header("Access-Control-Allow-Origin", origin)
        }
        
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        c.Header("Access-Control-Expose-Headers", "X-Total-Count, X-Auth-Token, X-Request-Id")
        c.Header("Access-Control-Max-Age", "86400") // 预检请求缓存时间（秒）
        
        // 处理 OPTIONS 预检请求
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}

// 使用中间件
r.Use(corsMiddleware())

// ========== 实战示例：同时使用两个响应头 ==========
// 场景：获取用户列表，需要发送 Authorization 请求头，并接收 X-Total-Count 响应头

// 1. 前端发送请求（包含自定义请求头 Authorization）
// 前端代码：
/*
fetch('http://api.example.com/users', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer token123',  // 自定义请求头，需要 Access-Control-Allow-Headers
    'Content-Type': 'application/json'
  },
  credentials: 'include'
})
*/

// 2. 浏览器先发送 OPTIONS 预检请求
// 请求头：
//   Origin: http://localhost:5173
//   Access-Control-Request-Method: GET
//   Access-Control-Request-Headers: Authorization, Content-Type

// 3. 后端响应预检请求（必须包含 Access-Control-Allow-Headers）
// 响应头：
//   Access-Control-Allow-Origin: http://localhost:5173
//   Access-Control-Allow-Methods: GET, POST, PUT, DELETE
//   Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With  ← 允许这些请求头
//   Access-Control-Allow-Credentials: true

// 4. 浏览器发送实际 GET 请求
// 请求头：
//   Origin: http://localhost:5173
//   Authorization: Bearer token123  ← 现在可以发送了
//   Content-Type: application/json

// 5. 后端处理请求并返回数据（必须包含 Access-Control-Expose-Headers）
func getUserList(c *gin.Context) {
    users := []User{...}
    totalCount := 100
    
    // 设置自定义响应头
    c.Header("X-Total-Count", strconv.Itoa(totalCount))
    c.Header("X-Auth-Token", "newtoken456")
    
    // 注意：Access-Control-Expose-Headers 已经在中间件中设置
    // 这样前端才能访问 X-Total-Count 和 X-Auth-Token
    
    c.JSON(200, gin.H{
        "code": 0,
        "data": users,
    })
}

// 6. 前端接收响应并访问响应头
// 前端代码：
/*
const response = await fetch('http://api.example.com/users', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer token123',
    'Content-Type': 'application/json'
  },
  credentials: 'include'
})

const data = await response.json()
const totalCount = response.headers.get('X-Total-Count')  // ✅ 可以访问（因为设置了 Access-Control-Expose-Headers）
const authToken = response.headers.get('X-Auth-Token')     // ✅ 可以访问
console.log('总数:', totalCount)  // 输出: 100
*/
```

### Node.js (Express) 示例

```js
const express = require('express')
const app = express()

// 使用 cors 中间件
const cors = require('cors')

app.use(cors({
    origin: ['http://localhost:5173', 'https://admin.example.com'],
    credentials: true,
    methods: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'],
    allowedHeaders: ['Content-Type', 'Authorization', 'X-Requested-With'],
    exposedHeaders: ['X-Total-Count', 'X-Auth-Token', 'X-Request-Id'],
    maxAge: 86400
}))

// 或者手动设置
app.use((req, res, next) => {
    const origin = req.headers.origin
    const allowedOrigins = ['http://localhost:5173', 'https://admin.example.com']

    if (allowedOrigins.includes(origin)) {
        res.setHeader('Access-Control-Allow-Origin', origin)
    }

    res.setHeader('Access-Control-Allow-Credentials', 'true')
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, PATCH, OPTIONS')
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization, X-Requested-With')
    res.setHeader('Access-Control-Expose-Headers', 'X-Total-Count, X-Auth-Token, X-Request-Id')
    res.setHeader('Access-Control-Max-Age', '86400')

    if (req.method === 'OPTIONS') {
        res.sendStatus(204)
    } else {
        next()
    }
})
```

### Python (Flask) 示例

```python
from flask import Flask
from flask_cors import CORS

app = Flask(__name__)

CORS(app, 
     origins=['http://localhost:5173', 'https://admin.example.com'],
     supports_credentials=True,
     methods=['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS'],
     allow_headers=['Content-Type', 'Authorization', 'X-Requested-With'],
     expose_headers=['X-Total-Count', 'X-Auth-Token', 'X-Request-Id'],
     max_age=86400)
```

---

## 前端请求配置

在你的项目中，`admin-client.js` 已经配置了 `withCredentials: true`：

```js
const requestConfig = {
    baseURL: import.meta.env.VITE_API_BASE_URL || 'http://127.0.0.1:8000',
    timeout: 10000,
    withCredentials: true, // ✅ 已配置
    headers: {
        'Content-Type': 'application/json'
    },
}
```

对应的 fetch 请求中：
```js
fetchOptions.credentials = 'include' // ✅ 已配置
```

---

## 常见问题

### 1. 预检请求（OPTIONS）
浏览器在发送复杂请求前会先发送 OPTIONS 预检请求，后端必须正确处理：

```js
// 后端必须返回 200 或 204，并包含所有 CORS 头
if (req.method === 'OPTIONS') {
    res.status(200).end()
    return
}
```

### 2. 凭证问题
如果前端设置了 `withCredentials: true`，后端必须：
- 设置 `Access-Control-Allow-Credentials: true`
- `Access-Control-Allow-Origin` 不能为 `*`，必须指定具体域名

### 3. 自定义响应头访问
如果前端需要访问自定义响应头（如分页总数），后端必须通过 `Access-Control-Expose-Headers` 暴露：

```js
// 后端
res.setHeader('X-Total-Count', '100')
res.setHeader('Access-Control-Expose-Headers', 'X-Total-Count')

// 前端
const totalCount = response.headers.get('X-Total-Count')
```

---

## 测试 CORS 配置

### 使用 curl 测试

```bash
# 测试 OPTIONS 预检请求
curl -X OPTIONS http://127.0.0.1:8000/api/users \
  -H "Origin: http://localhost:5173" \
  -H "Access-Control-Request-Method: GET" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -v

# 测试实际请求
curl -X GET http://127.0.0.1:8000/api/users \
  -H "Origin: http://localhost:5173" \
  -H "Cookie: session=xxx" \
  -v
```

### 浏览器控制台检查

打开浏览器开发者工具 → Network 标签：
1. 查看请求的 Response Headers 是否包含 CORS 头
2. 检查是否有 CORS 错误信息
3. 查看 OPTIONS 预检请求是否成功

---

## 推荐配置

### 开发环境
使用 Vite 代理，避免跨域问题

### 生产环境
后端服务器正确设置 CORS 响应头

