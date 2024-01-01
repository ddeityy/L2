package api

import (
	"calendar/server"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (a *API) CreateEvent(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("can't parse user_id: %v", err))
		return
	}

	date, err := c.Get("date")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("can't parse date, use RFC3339 format: %v", err))
		return
	}

	title, err := c.Get("title")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	if title == "" {
		c.Error(http.StatusBadRequest, fmt.Errorf("no title provided"))
		return
	}

	event := Event{
		UserID: uint64(user_id),
		Title:  title,
		Date:   t,
	}

	result := a.DB.Create(&event)
	if result.Error != nil {
		c.Error(http.StatusInternalServerError, fmt.Errorf("database error: %v", err))
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (a *API) UpdateEvent(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("can't parse user_id: %v", err))
		return
	}

	date, err := c.Get("date")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("can't parse date, use RFC3339 format: %v", err))
		return
	}

	title, err := c.Get("title")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	if title == "" {
		c.Error(http.StatusBadRequest, fmt.Errorf("no title provided"))
		return
	}

	eid, err := c.Get("id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	event_id, err := strconv.Atoi(eid)
	if err != nil {
		c.Error(
			http.StatusBadRequest,
			fmt.Errorf("can't parse id: %v", err),
		)
		return
	}

	event := Event{ID: uint64(event_id), UserID: uint64(user_id)}

	result := a.DB.First(&event)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	event.Date = t
	event.Title = title

	result = a.DB.Save(&event)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	c.NoContent()
}

func (a *API) DeleteEvent(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(http.StatusBadRequest, fmt.Errorf("can't parse user_id: %v", err))
		return
	}

	eid, err := c.Get("id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}
	event_id, err := strconv.Atoi(eid)
	if err != nil {
		c.Error(
			http.StatusBadRequest,
			fmt.Errorf("can't parse id: %v", err),
		)
		return
	}

	event := Event{ID: uint64(event_id), UserID: uint64(user_id)}

	result := a.DB.Delete(&event)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	c.NoContent()
}

func (a *API) GetEventsForDay(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(
			http.StatusBadRequest,
			fmt.Errorf("can't parse user_id: %v", err),
		)
		return
	}

	events := []Event{}

	result := a.DB.Find(&events, "user_id = ?", user_id)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	if len(events) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	res := make([]Event, 0, len(events))

	for _, event := range events {
		if event.Date.Day() == time.Now().Day() {
			res = append(res, event)
		}
	}

	if len(res) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a *API) GetEventsForWeek(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(
			http.StatusBadRequest,
			fmt.Errorf("can't parse user_id: %v", err),
		)
		return
	}

	events := []Event{}

	result := a.DB.Find(&events, "user_id = ?", user_id)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	if len(events) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	res := make([]Event, 0, len(events))

	for _, event := range events {
		ey, ew := event.Date.ISOWeek()
		ny, nw := time.Now().ISOWeek()
		if ey == ny && ew == nw {
			res = append(res, event)
		}
	}

	if len(res) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a *API) GetEventsForMonth(c *server.Context) {
	uid, err := c.Get("user_id")
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	user_id, err := strconv.Atoi(uid)
	if err != nil {
		c.Error(
			http.StatusBadRequest,
			fmt.Errorf("can't parse user_id: %v", err),
		)
		return
	}

	events := []Event{}

	result := a.DB.Find(&events, "user_id = ?", user_id)
	if result.Error != nil {
		c.Error(
			http.StatusInternalServerError,
			fmt.Errorf("database error: %v", result.Error),
		)
		return
	}

	if len(events) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	res := make([]Event, 0, len(events))

	for _, event := range events {
		if event.Date.Month() == time.Now().Month() {
			res = append(res, event)
		}
	}

	if len(res) == 0 {
		c.Error(
			http.StatusNotFound,
			fmt.Errorf("no events matching user_id %v found", user_id),
		)
		return
	}

	c.JSON(http.StatusOK, res)
}
