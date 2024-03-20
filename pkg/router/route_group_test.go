package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRouteGroup(t *testing.T) {
	r := NewDefaultRouter()
	rg := NewRouteGroup(r.(*DefaultRouter), "/test")

	assert.NotNil(t, rg)
	assert.Equal(t, "/test", rg.prefix)
	assert.Equal(t, r, rg.router)
	assert.Empty(t, rg.middlewares)
}

func TestApplyMiddlewaresOrder(t *testing.T) {
	r := NewDefaultRouter()

	var order []int
	middleware1 := func(h HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order = append(order, 1)
			h(c)
		}
	}
	middleware2 := func(h HandlerFunc) HandlerFunc {
		return func(c *Context) {
			order = append(order, 2)
			h(c)
		}
	}
	rg := NewRouteGroup(r.(*DefaultRouter), "/test", middleware1, middleware2)

	handler := func(c *Context) {}
	rg.applyMiddlewares(handler)(&Context{})

	assert.Equal(t, []int{1, 2}, order)
}

func TestApplyMiddlewaresEffect(t *testing.T) {
	r := NewDefaultRouter()

	var called int
	middleware := func(h HandlerFunc) HandlerFunc {
		return func(c *Context) {
			called++
			h(c)
		}
	}
	rg := NewRouteGroup(r.(*DefaultRouter), "/test", middleware)

	handler := func(c *Context) {
		called++
	}
	rg.applyMiddlewares(handler)(&Context{})

	assert.Equal(t, 2, called)
}
