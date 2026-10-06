# billAI-Category-service

HTTP-сервис для получения списка категорий

## Требования 

- Docker и Docker Compose
- Go 1.26+
- Docker network

## Быстрый запуск 

Поднять сервис вместе с бд

```bash
make up
```

По умолчанию сервис запускается на `http://localhost:8081`.

## Конфигурация 

```env.example
POSTGRES_DSN=postgres://postgres:postgres@categorydb:5432/categorydb?sslmode=disable
HTTP_PORT=8081
REDIS_DSN=redis-broker:6379
```

## Использование API

Посмотреть список категорий пользователя

```bash
curl http://localhost:8081/load \
  -H "X-UserID:d478dee6-efbb-4614-9730-e63358b5b93a"
```

Пример успешного ответа:

```bash
{
    "categories": [
        {
            "Id": 1,
            "UserID": "d478dee6-efbb-4614-9730-e63358b5b93a",
            "name": "Продукты питания"
        }
    ]
}
```

Если мы ввели неправильный заголовок 

```bash
{
    "categories": null
}

```

## Полезные Команды

```bash
make up       # поднять Docker-сервисы
make down     # удалить Docker-сервисы
```
