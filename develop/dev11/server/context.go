package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Used for handlers
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
	mu      sync.RWMutex
}

// Parses both URL and Form data
func (c *Context) ParseParams() {
	err := c.Request.ParseForm()
	if err != nil {
		c.Error(http.StatusInternalServerError, fmt.Errorf("form parse error: %v", err))
	}
	c.Params = make(map[string]string)
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.Request.Form {
		c.Params[k] = v[0]
	}
}

// Returns a param value by key
func (c *Context) Get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value := c.Params[key]
	if value == "" {
		return "", fmt.Errorf("can't parse value %v: %v is empty", key, key)
	}
	return value, nil
}

// Sends a json response
func (c *Context) JSON(status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	b := make(map[string]interface{})

	switch v := v.(type) {
	case error:
		b["error"] = v.Error()
	default:
		b["result"] = v
	}

	if err := enc.Encode(b); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	c.Writer.WriteHeader(status)

	c.Writer.Write(buf.Bytes())
}

// Sends an error as json
func (c *Context) Error(httpStatusCode int, err error) {
	c.JSON(httpStatusCode, err)
}

// Sends no content response
func (c *Context) NoContent() {
	c.Writer.WriteHeader(http.StatusNoContent)
}
