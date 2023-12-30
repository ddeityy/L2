package main

/*
Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов

Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
В случае остальных ошибок сервер должен возвращать HTTP 500.
Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

import (
	"calendar/api"
	"calendar/http/server"
	"os"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	server := server.NewServer(logger)

	logger.Info("Connecting to database")

	db, err := api.ConnectDB(logger)
	if err != nil {
		logger.Panic(err.Error())
	}

	api := api.NewAPI(db)

	server.AddRoute("/create_event", api.CreateEvent)
	server.AddRoute("/update_event", api.UpdateEvent)
	server.AddRoute("/delete_event", api.DeleteEvent)
	server.AddRoute("/events_for_day", api.GetEventsForDay)
	server.AddRoute("/events_for_week", api.GetEventsForWeek)
	server.AddRoute("/events_for_month", api.GetEventsForMonth)

	server.Run()

}
