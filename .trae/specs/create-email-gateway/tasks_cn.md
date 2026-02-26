# 任务列表 (Tasks)

- [x] 任务 1: 项目初始化
  - [x] 初始化 Go 模块 (`go mod init fastmail`)
  - [x] 创建项目结构 (`cmd/`, `internal/config`, `internal/api`, `internal/service`)
  - [x] 添加 `.gitignore` 和 `README.md`。

- [x] 任务 2: 配置管理
  - [x] 实现 `internal/config` 包，使用 `viper` 读取 `config.yaml`、环境变量和命令行标志。
  - [x] 定义 `Config` 结构体 (包含服务器端口、Auth Token、SMTP 主机/端口/用户/密码)。
  - [x] 添加命令行标志支持 (例如 `-config`, `-port`)。

- [x] 任务 3: SMTP 服务实现
  - [x] 实现 `internal/service/email.go`，使用 `gopkg.in/gomail.v2` (或类似库)。
  - [x] 创建 `EmailService` 结构体和 `SendEmail` 方法。
  - [x] 支持 HTML 邮件正文和文件附件。
  - [x] 实现连接池/重用（可选），或简单地为每个请求创建新连接。

- [x] 任务 4: HTTP API 实现
  - [x] 实现 `internal/api/handler.go`，使用 `gin` 框架。
  - [x] 创建 `POST /api/v1/send` 处理器。
  - [x] 解析 `multipart/form-data` 请求体：`to`, `subject`, `body`, `attachments`。
  - [x] 验证必填字段。
  - [x] 调用 `EmailService.SendEmail`。
  - [x] 返回 JSON 响应 (成功/错误)。

- [ ] 任务 5: 认证与中间件
  - [x] 实现 `AuthMiddleware`，验证 `Authorization` 头是否匹配配置的 Token。
  - [ ] 实现 `RequestLogger` 中间件，用于结构化日志记录。
  - [x] 添加 `GET /health` 接口用于健康检查。

- [x] 任务 6: 主程序入口
  - [x] 在 `cmd/server/main.go` 中连接 `config`、`service` 和 `api`。
  - [x] 实现优雅停机处理 (Context cancellation on SIGINT/SIGTERM)。

- [x] 任务 7: 构建与文档
  - [x] 创建 `Makefile` 用于构建二进制文件 (`make build`).
  - [x] 创建示例配置文件 `config.yaml`.
  - [x] 在 `README.md` 中编写使用文档 (中文+英文).
