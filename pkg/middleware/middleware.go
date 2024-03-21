package middleware

import (
	"github.com/mbrunos/go-hire/pkg/middleware/jwtauth"
	"github.com/mbrunos/go-hire/pkg/middleware/requestlogger"
)

var JwtAuth = jwtauth.JwtAuth
var Logger = requestlogger.Logger
