# FastMail Gateway

<img src="logo.png" width="200">

A lightweight, standalone email gateway service written in Go. It provides a simple HTTP API to send emails via SMTP (e.g., Gmail), supporting attachments and HTML content. Designed for ease of use and deployment as a single binary.

[English](README.md) | [中文文档](README_CN.md)

---

### Features

- **Simple HTTP API**: Send emails via a RESTful `POST` endpoint.
- **SMTP Support**: Works with any SMTP provider (optimized for Gmail).
- **Attachments**: Supports multiple file attachments.
- **HTML Content**: Send rich HTML emails.
- **Secure**: Token-based authentication for API access.
- **Configurable**: extensive configuration via `config.yaml`, environment variables, or command-line flags.
- **Single Binary**: Easy to deploy with no external runtime dependencies.

### Getting Started

#### Prerequisites

- Go 1.18 or higher (for building from source)

#### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/yourusername/fastmail.git
    cd fastmail
    ```

2.  **Build the binary**
    ```bash
    make build
    # Or manually: go build -o bin/server cmd/server/main.go
    ```

#### Configuration

1.  Copy the example configuration file:
    ```bash
    cp config.example.yaml config.yaml
    ```

2.  Edit `config.yaml` with your SMTP settings.
    *   **Note for Gmail Users**: You must use an **App Password** instead of your login password. Generate one at [Google Account Security](https://myaccount.google.com/apppasswords).

    ```yaml
    server:
      port: 8080
      token: "your-secret-token" # Token for API authentication

    smtp:
      host: "smtp.gmail.com"
      port: 587
      user: "your-email@gmail.com"
      pass: "your-app-password"
    ```

3.  **Environment Variables** (Optional override):
    - `FASTMAIL_SERVER_PORT`
    - `FASTMAIL_SERVER_TOKEN`
    - `FASTMAIL_SMTP_HOST`
    - `FASTMAIL_SMTP_PORT`
    - `FASTMAIL_SMTP_USER`
    - `FASTMAIL_SMTP_PASS`

#### Usage

Start the server:
```bash
./bin/server
# Or with flags:
./bin/server -port 9090 -config myconfig.yaml
```

#### API Reference

**Endpoint**: `POST /api/v1/send`

**Headers**:
- `Authorization: Bearer <your-token>`
- `Content-Type: multipart/form-data`

**Parameters**:

| Field | Type | Required | Description |
| :--- | :--- | :--- | :--- |
| `to` | string | Yes | Recipient email address(es), comma-separated. |
| `subject` | string | Yes | Email subject. |
| `body` | string | Yes | Email body (HTML supported). |
| `attachments` | file | No | File(s) to attach. |

**Example (cURL)**:
```bash
curl -X POST http://localhost:8080/api/v1/send \
  -H "Authorization: Bearer your-secret-token" \
  -F "to=recipient@example.com" \
  -F "subject=Hello from FastMail" \
  -F "body=<h1>This is a test email</h1>" \
  -F "attachments=@/path/to/document.pdf"
```

## Contact

- QQ: 546253846
- Email: sherlock.yin1994@gmail.com

## License

MIT License
