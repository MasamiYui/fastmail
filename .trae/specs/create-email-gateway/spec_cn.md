# 邮件网关设计文档 (Email Gateway Spec)

## 背景 (Why)
当前需要一个简单、独立的 HTTP 服务，用于通过 Gmail SMTP 协议发送邮件。该服务需要支持编译为单个二进制文件，并通过配置文件或命令行参数进行灵活配置。

## 变更内容 (What Changes)
- 创建一个新的 Go 项目结构。
- 实现一个 HTTP 服务器，提供发送邮件的 API 接口。
- 集成 Gmail SMTP 用于邮件投递。
- 实现简单的基于 Token 的身份验证机制。
- 支持通过文件（如 `config.yaml`）和命令行参数进行配置。
- 支持发送带附件的邮件。

## 影响范围 (Impact)
- **新服务**: 一个独立的二进制程序 `email-gateway`。
- **API**: 暴露新的 HTTP API 接口。

## 新增需求 (ADDED Requirements)

### 需求：配置管理
系统必须支持从以下途径加载配置：
1.  配置文件（例如 `config.yaml`）。
2.  环境变量。
3.  命令行参数（例如 `--port`, `--config`）。

**关键配置项**:
- 服务器端口 (Server Port)
- API Token (用于身份验证)
- SMTP 主机 (例如 `smtp.gmail.com`)
- SMTP 端口 (例如 `587`)
- SMTP 用户名/密码 (发件人凭据)

### 需求：邮件发送 HTTP API
系统必须暴露 `POST /api/v1/send` 接口。
- **身份验证**: Bearer Token 或 `X-Api-Token` 请求头，需匹配配置的 Token。
- **Content-Type**: `multipart/form-data` (以支持附件上传)。
- **参数**:
    - `to`: 收件人邮箱地址。
    - `subject`: 邮件标题。
    - `body`: 邮件正文 (HTML 或 纯文本)。
    - `attachments`: 文件附件。

#### 场景：成功发送
- **当** 客户端发送一个包含正确 Token 和邮件数据的有效 POST 请求时。
- **则** 系统连接到 SMTP 服务器并发送邮件。
- **则** 系统返回 `200 OK` 及成功消息。

#### 场景：认证失败
- **当** 客户端发送请求时 Token 无效或缺失。
- **则** 系统返回 `401 Unauthorized`。

### 需求：部署
系统必须编译为单个二进制文件，除配置文件外不依赖外部运行时环境。

### 需求：可观测性与健壮性 (其他考量)
- **健康检查**: 提供 `GET /health` 接口以检查服务状态。
- **日志记录**: 使用结构化日志 (JSON) 记录请求追踪和错误报告。
- **优雅停机**: 处理操作系统信号 (SIGINT, SIGTERM)，在退出前完成正在处理的请求。
- **超时控制**: 为 SMTP 连接设置合理的超时时间，避免请求挂起。
