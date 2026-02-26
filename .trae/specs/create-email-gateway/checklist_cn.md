# 检查清单 (Checklist)

* [x] `go build` 成功创建名为 `fastmail` 的二进制文件。

* [x] 运行 `./fastmail` 不带配置文件时，回退到默认值或干净地报错。

* [x] 运行 `./fastmail` 带有配置文件时，正确加载 SMTP 设置。

* [x] API 端点 `GET /health` 返回 200 OK。

* [x] API 端点 `POST /api/v1/send` 拒绝没有 `Authorization` 头的请求 (401 Unauthorized)。

* [x] API 端点 `POST /api/v1/send` 接受有效请求并尝试通过配置的 SMTP 服务器发送邮件。

* [x] 邮件成功发送到收件人邮箱，且主题和正文正确。

* [x] 邮件附件正确附加，且收件人可查看。

* [x] 优雅停机功能正常 (在退出前完成正在处理的请求)。

