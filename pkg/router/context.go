package router

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  w,
	}
}

func (c *Context) PathParam(key string) string {
	return c.Request.PathValue(key)
}

func (c *Context) QueryParam(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) BodyParam(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) HeaderParam(key string) string {
	return c.Request.Header.Get(key)
}

// BindJSON binds the request body to a struct, and it must
// receive a pointer to the struct
func (c *Context) BindJSON(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

func (c *Context) SendJSON(status int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) SendString(status int, data string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(data))
}

func (c *Context) SendError(status int, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(`{"error": "` + err.Error() + `"}`))
}

func (c *Context) Set(key string, value interface{}) {
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}

func (c *Context) Get(key string) interface{} {
	return c.Request.Context().Value(key)
}
