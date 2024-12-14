# Запуск
```bash
docker compose up --build
```
# Курлы для проверки
```bash
curl -X POST "http://localhost:8888/auth/token?user_id=12345"
```
```bash
curl -X POST "http://localhost:8888/auth/refresh" \            
-H "Content-Type: application/json" \
-d '{
  "access_token": "access токен из ответа выше",
  "refresh_token": "refresh токен из ответа выше"
}'
```
