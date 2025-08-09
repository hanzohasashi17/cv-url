# cv-url

## Запуск

```bash
go run ./cmd/server
```

## Тест

```bash
curl -X POST http://localhost:8080/send \
  -H "Content-Type: application/json" \
  -d '{
        "cv_url":"https://hh.kz/resume/84ed2ed2ff0dd13d0b0039ed1f7a5963625869",
        "email":"nomadbaj@gmail.com",
        "smtp_login":"you@gmail.com",
        "smtp_password":"app_pass"
    }'
```