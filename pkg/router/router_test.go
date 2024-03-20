package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultRouter(t *testing.T) {
	r := NewDefaultRouter()

	assert.NotNil(t, r)
}

func TestDefaultRouterApplyMiddlewares(t *testing.T) {
	r := NewDefaultRouter()

	var called int
	middleware := func(h HandlerFunc) HandlerFunc {
		return func(c *Context) {
			called++
			h(c)
		}
	}

	r.Use(middleware)

	handler := func(c *Context) {
		called++
	}
	r.applyMiddlewares(handler)(&Context{})

	assert.Equal(t, 2, called)
}
