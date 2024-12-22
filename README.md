# calc_go
api calc

## веб-сервис: пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.
#### У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url POST-запрос с телом:
#### {"expression": "выражение, которое ввёл пользователь"}
#### ответ пользователь получает HTTP-ответ с телом:
#### {"result": "результат выражения"}
#### и кодом 200, если выражение вычислено успешно, либо HTTP-ответ с телом:
#### { "error": "Expression is not valid"}
#### и кодом 422, если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.
#### Ещё один вариант HTTP-ответа:
#### {"error": "Internal server error"}
#### и код 500 в случае какой-либо иной ошибки («Что-то пошло не так»).


## запуск проекта:
#### go run ".\cmd\main.go"

## запросы:

### curl --location 'localhost:8080/api/v1/calculate' \ --header 'Content-Type: application/json' \ --data '{"expression": "2+2"}'

> вывод: result: 4.000000 и статус 200 OK

### curl --location 'localhost:8080/api/v1/calculate' \ --header 'Content-Type: application/json' \ --data '{"expression": "6/0"}'

> вывод: error: Internal server error и статус 500 Internal Server Error

### curl --location --request GET 'localhost:8080/api/v1/calculate' \ --header 'Content-Type: application/json' \ --data '{"expression": "(2+2)*2"}'

> вывод: error: Method not allowed и статус 405 Method not allowed

### curl --location --request GET 'localhost:8080/api/v1/calculate' \ --header 'Content-Type: application/json' \ --data '{"expression": "(2+2)**2"}'

> вывод: error: Expression is not valid и статус 422 Unprocessable Entity



## реализация логов:

* запуск сервера: 2021/10/20 12:00:00: "Starting server..."

* Лог о каждом запросе: 2021/10/20 12:00:00: Received request: r.Method r.URL.Path r.RemoteAddr

* и возврат "Internal Server Error", с кодом http.StatusInternalServerError




## примеры ошибок

1) errInternalServer( "Internal server error") и кодом 500

> Пример 1: Проблема с JSON-декодированием

> Пример 2: Ошибка при парсинге или вычислении

> Пример 3: Неожиданный сбой в процессе вычислений

> Пример 4: Проблемы с ресурсами сервера


2) errInvalidExpression("Expression is not valid") и кодом 422

> пример 1:  входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.

> пример 2: непарность скобок

> пример 3: длина выражения меньше 3


3) ErrMethodNotAllowed и кодом 405

> пример 1: метод не post


## реализация файла application_test.go с проверками

> TestCalcHandler_Success тест без ошибок: тело запроса `{"expression": "2+2"}`, должен возвращать "result: 4.000000"

> TestCalcHandler_Success2 тест без ошибок: тело запроса `{"expression": "1/2"}`, должен возвращать "result: 0.500000"

> TestCalcHandler_InvalidExpression: при некорректном выражении, ошибка Expression is not valid

> TestCalcHandler_MethodNotAllowed: при методе GET, ошибка Method not allowed

> TestCalcHandler_InternalServerError: при делении на ноль, ошибка Internal server error