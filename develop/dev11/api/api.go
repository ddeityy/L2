package api

import (
	"net/http"

	"gorm.io/gorm"
)

type API struct {
	DB *gorm.DB
}

func NewAPI(DB *gorm.DB) API {
	return API{
		DB: DB,
	}
}

func (a *API) CreateEvent(http.ResponseWriter, *http.Request) {

}

func (a *API) UpdateEvent(http.ResponseWriter, *http.Request) {

}

func (a *API) DeleteEvent(http.ResponseWriter, *http.Request) {

}

func (a *API) GetEventsForDay(http.ResponseWriter, *http.Request) {

}

func (a *API) GetEventsForWeek(http.ResponseWriter, *http.Request) {

}

func (a *API) GetEventsForMonth(http.ResponseWriter, *http.Request) {

}
