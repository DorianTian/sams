my_project/
│
├── app/                  # 业务逻辑层（用于存放控制器、服务等）
│   ├── controller/       # 控制器层：处理请求并返回响应
│   │   ├── user_controller.go
│   │   └── auth_controller.go
│   ├── service/          # 服务层：处理业务逻辑
│   │   ├── user_service.go
│   │   └── auth_service.go
│   └── model/            # 数据模型层（存放结构体和数据库模型）
│       ├── user.go
│       └── auth.go
│
├── config/               # 配置文件：配置文件、常量等
│   ├── config.go
│   └── app_config.yaml
│
├── routes/               # 路由层：定义 API 路由
│   └── routes.go
│
├── middleware/           # 中间件：自定义中间件，如认证、日志等
│   └── auth_middleware.go
│
├── repository/           # 数据库交互层：与数据库交互的代码
│   ├── user_repository.go
│   └── auth_repository.go
│
├── util/                 # 工具类：通用工具函数
│   ├── logger.go
│   └── hash_util.go
│
├── public/               # 静态文件目录：如图片、CSS、JS 等
│   └── index.html
│
├── ui/                   # 前端代码
│   ├── package.json      # 前端依赖和脚本
│   ├── pnpm-lock.yaml    # pnpm 锁文件
│   ├── tsconfig.json     # TypeScript 配置
│   ├── tailwind.config.js # Tailwind CSS 配置
│   ├── postcss.config.js # PostCSS 配置
│   ├── .eslintrc.json     # ESLint 配置文件
│   └── src/              # 前端代码存放
│       └── pages/        # Next.js 页面文件夹
│
├── .npmrc                # 限制使用 pnpm 的配置
├── .gitignore            # Git 忽略文件
├── go.mod                # Go 项目模块文件
├── go.sum                # Go 模块的校验文件
├── pnpm-workspace.yaml   # 此文件让所有子项目成为一个 pnpm 工作空间
└── main.go               # 项目入口：启动 Gin 应用，加载路由和中间件