package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	c := NewContext(w, r)
	assert.NotNil(t, c)
	assert.Equal(t, w, c.Writer)
	assert.Equal(t, r, c.Request)
}

func TestPathParam(t *testing.T) {
	r := httptest.NewRequest("GET", "/test/value", nil)
	w := httptest.NewRecorder()

	r.SetPathValue("param", "value")

	c := NewContext(w, r)

	assert.Equal(t, "value", c.PathParam("param"))
}

func TestQueryParam(t *testing.T) {
	r := httptest.NewRequest("GET", "/test?param=value", nil)
	w := httptest.NewRecorder()

	c := NewContext(w, r)

	assert.Equal(t, "value", c.QueryParam("param"))
}

func TestBodyParam(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", strings.NewReader("param=value"))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	c := NewContext(w, r)

	assert.Equal(t, "value", c.BodyParam("param"))
}

func TestHeaderParam(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	r.Header.Add("param", "value")
	w := httptest.NewRecorder()

	c := NewContext(w, r)

	assert.Equal(t, "value", c.HeaderParam("param"))
}

func TestBindJSON(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", strings.NewReader(`{"param": "value"}`))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c := NewContext(w, r)

	var body map[string]string

	err := c.BindJSON(&body)

	assert.Nil(t, err)
	assert.Equal(t, "value", body["param"])
}

func TestSendJSON(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	c := NewContext(w, r)
	c.SendJSON(http.StatusOK, map[string]string{"param": "value"})

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	assert.Equal(t, `{"param":"value"}`+"\n", w.Body.String())
}

func TestSendString(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	c := NewContext(w, r)
	c.SendString(http.StatusOK, "test")

	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))

	assert.Equal(t, "test", w.Body.String())
}

func TestSendError(t *testing.T) {
	w := httptest.NewRecorder()

	c := &Context{Writer: w}

	status := http.StatusInternalServerError
	err := errors.New("test error")

	c.SendError(status, err)

	assert.Equal(t, status, w.Code)

	contentType := w.Header().Get("Content-Type")

	assert.Equal(t, "application/json", contentType)

	expectedBody := `{"error": "` + err.Error() + `"}`

	assert.Equal(t, expectedBody, w.Body.String())
}
