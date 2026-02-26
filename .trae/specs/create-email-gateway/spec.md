# Email Gateway Spec

## Why
Currently, there is a need for a simple, standalone service to send emails via HTTP API using Gmail's SMTP protocol. This service should be easy to deploy as a single binary and configurable via file or command-line arguments.

## What Changes
- Create a new Go project structure.
- Implement an HTTP server with endpoints for sending emails.
- Integrate with Gmail SMTP for email delivery.
- Implement a simple token-based authentication mechanism.
- Support configuration via file (e.g., `config.yaml`) and command-line flags.
- Support attachments in email sending.

## Impact
- **New Service**: A standalone binary `email-gateway`.
- **API**: New HTTP API endpoints exposed.

## ADDED Requirements

### Requirement: Configuration Management
The system SHALL support loading configuration from:
1.  A configuration file (e.g., `config.yaml`).
2.  Environment variables.
3.  Command-line flags (e.g., `--port`, `--config`).

**Key Config Items**:
- Server Port
- API Token (for authentication)
- SMTP Host (e.g., `smtp.gmail.com`)
- SMTP Port (e.g., `587`)
- SMTP Username/Password (Sender email credentials)

### Requirement: HTTP API for Email Sending
The system SHALL expose a `POST /api/v1/send` endpoint.
- **Authentication**: Bearer Token or `X-Api-Token` header matching the configured token.
- **Content-Type**: `multipart/form-data` (to support attachments).
- **Parameters**:
    - `to`: Recipient email address(es).
    - `subject`: Email subject.
    - `body`: Email body (HTML or Text).
    - `attachments`: File uploads.

#### Scenario: Success case
- **WHEN** a client sends a valid POST request with correct token and email data.
- **THEN** the system connects to the SMTP server and sends the email.
- **THEN** the system returns `200 OK` with a success message.

#### Scenario: Auth Failure
- **WHEN** a client sends a request with an invalid or missing token.
- **THEN** the system returns `401 Unauthorized`.

### Requirement: Deployment
The system SHALL compile to a single binary with no external runtime dependencies (other than the config file).

### Requirement: Observability & Robustness (Additional Considerations)
- **Health Check**: `GET /health` endpoint to verify service status.
- **Logging**: Structured logging (JSON) for request tracking and error reporting.
- **Graceful Shutdown**: Handle OS signals (SIGINT, SIGTERM) to finish in-flight requests before exiting.
- **Timeouts**: Set reasonable timeouts for SMTP connections to avoid hanging requests.
