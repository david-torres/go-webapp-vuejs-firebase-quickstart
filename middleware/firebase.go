package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/david-torres/go-webapp-vuejs-firebase-quickstart/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// FirebaseAuthMiddleware contains methods pertaining to firebase authentication
type FirebaseAuthMiddleware struct {
	Skipper  middleware.Skipper
	firebase *services.FirebaseService
}

// NewFirebaseAuthMiddleware creates an instance of FirebaseAuthMiddleware
func NewFirebaseAuthMiddleware(firebase *services.FirebaseService) *FirebaseAuthMiddleware {
	return &FirebaseAuthMiddleware{firebase: firebase, Skipper: firebaseAuthSkipper}
}

// Verify is the middleware function, validates a Firebase id token
func (f *FirebaseAuthMiddleware) Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if f.Skipper(c) {
			return next(c)
		}

		r := c.Request()
		token := token(r)

		if token == "" {
			err := errors.New("Invalid token")
			return err
		}

		authToken, err := f.firebase.Verify(token)
		if err != nil {
			err := errors.New("Invalid token")
			return err
		}

		// set UID for controllers
		c.Set("UID", authToken.UID)
		return next(c)
	}
}

// add conditions to skip middleware here
func firebaseAuthSkipper(c echo.Context) bool {
	return false
}

// fetch the token from the header
func token(r *http.Request) string {
	authHeader := r.Header.Get(echo.HeaderAuthorization)
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return token
}
