# FastMail Gateway

<img src="logo.png" width="200">

一个使用 Go 编写的轻量级、独立邮件网关服务。它提供了一个简单的 HTTP API，通过 SMTP（例如 Gmail）发送邮件，支持附件和 HTML 内容。设计目标是易于使用，并可部署为单个二进制文件。

[English](README.md) | [中文文档](README_CN.md)

---

### 功能特性

- **简单 HTTP API**: 通过 RESTful `POST` 接口发送邮件。
- **SMTP 支持**: 支持任意 SMTP 服务商 (针对 Gmail 优化)。
- **附件支持**: 支持发送多个附件。
- **HTML 内容**: 支持富文本 HTML 邮件。
- **安全验证**: 基于 Token 的 API 访问验证。
- **灵活配置**: 支持 `config.yaml`、环境变量或命令行参数配置。
- **单文件部署**: 编译为单个二进制文件，无外部运行时依赖。

## 架构说明

### 系统架构
本服务作为一个轻量级中间件，连接您的应用与 SMTP 服务商。
1. **API 层**: 使用 Gin Gonic 提供的 RESTful 接口。
2. **认证层**: 通过 Bearer Token 进行请求验证。
3. **服务层**: 负责邮件内容的构建及附件处理。
4. **集成层**: 与外部 SMTP 服务器通信，完成最终投递。

### 技术栈
- **核心语言**: [Go](https://go.dev/) (1.25+)
- **Web 框架**: [Gin Gonic](https://gin-gonic.com/)
- **邮件库**: [Gomail](https://github.com/go-gomail/gomail)
- **配置管理**: YAML 与 环境变量

### 快速开始

#### 前置要求

- Go 1.25 或更高版本 (用于从源码编译)

#### 安装步骤

1.  **克隆代码仓库**
    ```bash
    git clone https://github.com/yourusername/fastmail.git
    cd fastmail
    ```

2.  **编译二进制文件**
    ```bash
    make build
    # 或者手动编译: go build -o bin/server cmd/server/main.go
    ```

#### 配置指南

1.  复制示例配置文件：
    ```bash
    cp config.example.yaml config.yaml
    ```

2.  编辑 `config.yaml` 填入你的 SMTP 设置。
    *   **Gmail 用户注意**: 你必须使用 **应用专用密码 (App Password)** 而不是你的登录密码。请在 [Google 账号安全设置](https://myaccount.google.com/apppasswords) 中生成。

    ```yaml
    server:
      port: 8080
      token: "your-secret-token" # 用于 API 验证的 Token

    smtp:
      host: "smtp.gmail.com"
      port: 587
      user: "your-email@gmail.com"
      pass: "your-app-password"
    ```

3.  **环境变量** (可选覆盖):
    - `FASTMAIL_SERVER_PORT`
    - `FASTMAIL_SERVER_TOKEN`
    - `FASTMAIL_SMTP_HOST`
    - `FASTMAIL_SMTP_PORT`
    - `FASTMAIL_SMTP_USER`
    - `FASTMAIL_SMTP_PASS`

#### 使用说明

启动服务：
```bash
./bin/server
# 或者指定参数启动:
./bin/server -port 9090 -config myconfig.yaml
```

#### API 参考

**接口地址**: `POST /api/v1/send`

**请求头 (Headers)**:
- `Authorization: Bearer <your-token>`
- `Content-Type: multipart/form-data`

**参数说明**:

| 字段名 | 类型 | 必填 | 描述 |
| :--- | :--- | :--- | :--- |
| `to` | string | 是 | 收件人邮箱地址，多个地址用逗号分隔。 |
| `subject` | string | 是 | 邮件主题。 |
| `body` | string | 是 | 邮件正文 (支持 HTML)。 |
| `attachments` | file | 否 | 附件文件 (支持多个)。 |

**调用示例 (cURL)**:
```bash
curl -X POST http://localhost:8080/api/v1/send \
  -H "Authorization: Bearer your-secret-token" \
  -F "to=recipient@example.com" \
  -F "subject=来自 FastMail 的问候" \
  -F "body=<h1>这是一封测试邮件</h1>" \
  -F "attachments=@/path/to/document.pdf"
```

## 联系方式

- QQ: 546253846
- 邮箱: sherlock.yin1994@gmail.com

## 许可证

MIT 许可证
