# Тестовое задание Авито Бэкенд

## Запуск программы

### Порт

Используется порт 8010

### Неподсредственно запуск
```ShellSession
docker-compose up
```
## Коллекции запросов

Коллекции запросов для REST можно импортировать [отсюда](https://github.com/bashkirian/go-mux-api/tree/internship/Postman)
## Test

Пока тесты можно прогнать локально, планируется сделать прогон через Docker.
```bash
$ go test -v
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.00s)
=== RUN   TestGetNonExistentUser
--- PASS: TestGetNonExistentUser (0.00s)
=== RUN   TestCreateWallet
--- PASS: TestCreateWallet (0.00s)
=== RUN   TestGetWallet
--- PASS: TestGetWallet (0.00s)
=== RUN   TestUpdateBalance
--- PASS: TestUpdateBalance (0.01s)
PASS
ok      _/home/tom/r/go-mux-api 0.034s
```

## License

Copyright (c) 2021 Rendered Text

Distributed under the MIT License. See the file LICENSE.
