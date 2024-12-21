# calc_go
api calc

запуск сервера:
go run ".\cmd\main.go"

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
