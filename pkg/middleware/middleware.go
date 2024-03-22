package middleware

import (
	"github.com/mbrunos/go-hire/pkg/middleware/jwtauth"
	"github.com/mbrunos/go-hire/pkg/middleware/requestlogger"
)

var JwtAuth = jwtauth.JwtAuthMiddleware
var Logger = requestlogger.Logger
