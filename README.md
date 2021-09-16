# fibonacci

## Установка
В папке с проектом запустить команду
<code>docker-compose up --build fibonacci</code>

## Как пользоваться
http REST API запрос: <code>http://localhost:8000/api/fibonacci?from=1&to=5 </code> 
где "from" это начиная с какого порядкового номера возвратить числа (должен быть положительным целым числом),
а "to" это до какого числа включительно (должен быть больше или равен "from" и не более чем 90)

gRPC: В качестве тестового клиента используется утилита evans, которую можно скачать в виде готового бинарника отсюда
<code> https://github.com/ktr0731/evans/releases </code>
Далее в папке с проектом запустить команду <code> evans proto/fibonacci.proto -p 8080 </code>,
сделать вызов к сервису <code> call Fibonacci </code> и ввести параметры.

## Кэш
Redis доступен из локального окружения по адресу localhost:63790 (внутри приложения redis:6379)

## Тесты
<code> go test ./... </code>