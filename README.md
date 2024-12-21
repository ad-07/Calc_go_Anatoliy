# calc_go
api calc

запуск сервера:
go run ".\cmd\main.go"
запросы:
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2"
}'
вывод: result: 4.000000 и статус 200 OK

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "6/0"
}'
вывод: error: Internal server error и статус 500 Internal Server Error

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "6+0"
}'




реализация логов:
запуск сервера: 2021/10/20 12:00:00: "Starting server..."
Лог о каждом запросе: 2021/10/20 12:00:00: Received request: r.Method r.URL.Path r.RemoteAddr
при ошибке net.SplitHostPort: 2021/10/20 12:00:00: Failed to parse RemoteAddr: err 
и возврат "Internal Server Error", с кодом http.StatusInternalServerError
при успешном net.SplitHostPort: 2021/10/20 12:00:00: Received request: r.Method r.URL.Path from IP: host, Port: port


примеры ошибок
-errInternalServer( "Internal server error") и кодом 500
Пример 1: Проблема с JSON-декодированием
Пример 2: Ошибка при парсинге или вычислении
Пример 3: Неожиданный сбой в процессе вычислений
Пример 4: Проблемы с ресурсами сервера
-errInvalidExpression("Expression is not valid") и кодом 422
пример 1:  входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.
пример 2: непарность скобок
пример 3: длина выражения меньше 3
-ErrMethodNotAllowed и кодом 405
пример 1: метод не post


реализация файла application_test.go
с проверками
-TestCalcHandler_Success тест без ошибок: тело запроса `{"expression": "2+2"}`, должен возвращать "result: 4.000000"
-TestCalcHandler_Success2 тест без ошибок: тело запроса `{"expression": "1/2"}`, должен возвращать "result: 0.500000"
-TestCalcHandler_InvalidExpression: при некорректном выражении, ошибка Expression is not valid
-TestCalcHandler_MethodNotAllowed: при методе GET, ошибка Method not allowed
-TestCalcHandler_InternalServerError: при делении на ноль, ошибка Internal server error