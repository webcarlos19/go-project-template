# API Examples

Este arquivo contém exemplos de uso da API do Go Project Template.

## Health Check

```bash
curl -X GET http://localhost:8080/health
```

**Resposta:**
```json
{
  "status": "OK",
  "timestamp": "2025-08-08T22:10:02.037Z",
  "version": "1.0.0",
  "service": "go-project-template",
  "database": "OK"
}
```

## Usuários

### Listar todos os usuários

```bash
curl -X GET http://localhost:8080/api/v1/users
```

**Resposta:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "created_at": "2025-08-08T22:10:09.793Z",
    "updated_at": "2025-08-08T22:10:09.793Z"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane.smith@example.com",
    "created_at": "2025-08-08T22:10:09.793Z",
    "updated_at": "2025-08-08T22:10:09.793Z"
  }
]
```

### Buscar usuário por ID

```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

**Resposta:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "created_at": "2025-08-08T22:10:09.793Z",
  "updated_at": "2025-08-08T22:10:09.793Z"
}
```

### Criar um novo usuário

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Carlos dos Santos",
    "email": "carlos@example.com"
  }'
```

**Resposta:**
```json
{
  "id": 4,
  "name": "Carlos dos Santos",
  "email": "carlos@example.com",
  "created_at": "2025-08-08T22:15:30.123Z",
  "updated_at": "2025-08-08T22:15:30.123Z"
}
```

### Atualizar um usuário

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "john.smith@example.com"
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "name": "John Smith",
  "email": "john.smith@example.com",
  "created_at": "2025-08-08T22:10:09.793Z",
  "updated_at": "2025-08-08T22:16:45.456Z"
}
```

### Deletar um usuário

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

**Resposta:** Status 204 (No Content)

## Tratamento de Erros

### Usuário não encontrado

```bash
curl -X GET http://localhost:8080/api/v1/users/999
```

**Resposta:**
```json
{
  "error": "Not Found",
  "message": "User not found",
  "code": 404
}
```

### Dados inválidos

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "email": "invalid-email"
  }'
```

**Resposta:**
```json
{
  "error": "Bad Request",
  "message": "Name and email are required",
  "code": 400
}
```

## Usando com diferentes ferramentas

### HTTPie

```bash
# Listar usuários
http GET localhost:8080/api/v1/users

# Criar usuário
http POST localhost:8080/api/v1/users name="Maria Silva" email="maria@example.com"

# Atualizar usuário
http PUT localhost:8080/api/v1/users/1 name="João Santos" email="joao@example.com"
```

### Postman

Importe a seguinte coleção no Postman:

```json
{
  "info": {
    "name": "Go Project Template API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080"
    }
  ],
  "item": [
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "url": "{{baseUrl}}/health"
      }
    },
    {
      "name": "Get Users",
      "request": {
        "method": "GET",
        "url": "{{baseUrl}}/api/v1/users"
      }
    },
    {
      "name": "Create User",
      "request": {
        "method": "POST",
        "url": "{{baseUrl}}/api/v1/users",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"name\": \"Test User\",\n  \"email\": \"test@example.com\"\n}"
        }
      }
    }
  ]
}
```

### JavaScript/Fetch

```javascript
// Listar usuários
fetch('http://localhost:8080/api/v1/users')
  .then(response => response.json())
  .then(data => console.log(data));

// Criar usuário
fetch('http://localhost:8080/api/v1/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'JavaScript User',
    email: 'js@example.com'
  })
})
.then(response => response.json())
.then(data => console.log(data));
```

### Python/Requests

```python
import requests

# Listar usuários
response = requests.get('http://localhost:8080/api/v1/users')
users = response.json()
print(users)

# Criar usuário
user_data = {
    'name': 'Python User',
    'email': 'python@example.com'
}
response = requests.post('http://localhost:8080/api/v1/users', json=user_data)
new_user = response.json()
print(new_user)
```
