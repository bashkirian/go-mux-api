# Тестовое задание Авито Бэкенд (пока версия без доп.запросов)

## Вопросы и ответы

Q: Как проверять баланс после резервации? Прописать ограничения в таблице, либо предусмотреть ошибку в коде?  
A: Я наложил constraint_positive на столбец user_balance в табличке balance. Если резервация существует в таблице, то принимается попытка уменьшить баланс, на сумму, указанную в резервации. Если он становится меньше нуля, выбрасывается ошибка postgresql.

Q: Какая должна быть структура базы данных?  
A: ![image](https://github.com/bashkirian/go-mux-api/blob/internship/build/docker/db/схема_дб.jpg)

Q: Нужны ли запросы для перевода с кошелька на кошелек и иные запросы, не указанные в задании?  
A: Я этого не сделал. Думаю, существующего функционала достаточно для принятия задания.

Q: Можно ли задавать пользователю самому имя хоста, номер порта, имя создаваемой БД и т.д.?  
A: Да, пользователь может задать имя хоста и номер порта в файле .env. Например, по умолчанию стоит хост для работы в докере.

```ShellSession
DB_HOST=balance-db    #with docker                   
#DB_HOST=localhost                            # when running the app without docker 
```  

Для того, чтобы запустить программу локально командой 

```ShellSession
go run .
```  
нужно раскомментировать вторую строку и закомментировать первую.  

## Запуск программы

### Порт

Используется порт 8010

### Неподсредственно запуск в докере
```ShellSession
docker-compose up --build
```

## Коллекции запросов

Коллекции запросов для REST можно импортировать [отсюда](https://github.com/bashkirian/go-mux-api/tree/internship/Postman)

## Тестирование

Пока тесты можно прогнать локально, планируется сделать прогон через Docker.  
```bash
$ go test -v
=== RUN   TestGetNonExistentProduct  
--- PASS: TestGetNonExistentProduct (0.01s)  
=== RUN   TestCreateProduct  
--- PASS: TestCreateProduct (0.01s)  
=== RUN   TestGetProduct  
--- PASS: TestGetProduct (0.01s)  
=== RUN   TestUpdateProduct  
--- PASS: TestUpdateProduct (0.01s)  
PASS  
ok      github.com/TomFern/go-mux-api   0.698s  
```

### Получение документации 

## http://localhost:8010/swagger/index.html

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/documentation.jpg)

## Обязательные запросы

### /balance/deposit

Принимает на вход json с ID пользователя и количеством зачисляемых средств. Возвращает обновленный кошелек. Если кошелька еще не существует, создает новый с заданной суммой. Если пополнение отрицательное, выбрасывается ошибка 400.

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/wallet_create.jpg)

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/negative_deposit.jpg)

### /balance/show/:id

Показывает баланс пользователя. Если пользователя не существует, выбрасывает ошибку 404. Иначе, возвращает кошелек пользователя с актуальным балансом.  

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/no_wallet.jpg)  

### /reservation

Создает резервацию с указанными параметрами. Если пользователя/резервации с таким ID не существует или же резервация уже существует, выбрасывает ошибку 500.

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/reservation_exists.jpg)

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/reservation_creation.jpg)  

### /reservation/accept

Принимает резервацию с указанными параметрами. Если пользователя/резервации с таким ID не существует или же резервация уже существует, выбрасывает ошибку 500.
Если баланс после принятия станет отрицательным, тоже будет ошибка 500.

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/negative_balance.jpg) 

![image](https://github.com/bashkirian/go-mux-api/blob/internship/requests/correct_reservation.jpg)  

## License

Copyright (c) 2022 Rendered Text

Distributed under the MIT License. See the file LICENSE.
