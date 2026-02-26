* [x] `go build` successfully creates a binary named `fastmail`.

* [x] Running `./fastmail` without config file falls back to defaults or errors cleanly.

* [x] Running `./fastmail` with config file loads SMTP settings correctly.

* [x] API endpoint `GET /health` returns 200 OK.

* [x] API endpoint `POST /api/v1/send` rejects requests without `Authorization` header (401 Unauthorized).

* [x] API endpoint `POST /api/v1/send` accepts valid requests and attempts to send email via configured SMTP server.

* [x] Email is successfully delivered to recipient inbox with correct subject and body.

* [x] Email attachment is correctly attached and viewable by recipient.

* [x] Graceful shutdown works (sends pending emails before exiting).

