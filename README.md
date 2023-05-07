# WB_L0 Test task
В проекте реализован сервис, который подключается к каналу в Nats Streaming, получает данные и сохраняет их в базу данных PostgreSQL и в кэш.
Сервис также поднимает HTTP-сервер, который предоставляет данные по UID из кэша. Cкрины интерфейса.
___
![](https://github.com/faringet/WB_L0/blob/master/screenshots/Board.jpg)

`Publisher` также генерирует пять сообщений, согласно вводной JSON-структуре.
Перед тем как отправить структуру, в ней генерируется и прописываются такие критические поля, как:
- `orderUid`
- `Transaction`
- `Rid`


## Основные пакеты: 
- обработка запросов с помощю [**Gin Web Framework**](https://gin-gonic.com/docs/)
- работа с БД и миграциями с помощью [**GORM**](https://gorm.io/docs/)
- запись осуществляется в [**PostgreSQL**](https://www.postgresql.org/)
- [**NATS**](https://nats.io/) 


# Запуск
Для запуска нам потребуется:

- PostgreSQL
```
docker run --name psql-container-wbL0 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=password -d postgres
```

- NATS
```
docker run --rm -d -p 4222:4222 -p 8222:8222 --name nats nats
```

- После этого запускаем сам сервис
```
main.go
```

- Идем по адресу
```
http://localhost:3000/
```
