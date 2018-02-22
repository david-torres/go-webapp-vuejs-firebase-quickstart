package controllers

import (
	"net/http"

	"github.com/david-torres/go-webapp-vuejs-firebase-quickstart/services"
	"github.com/labstack/echo"
)

// HelloWorldController contains methods pertaining to rendering and manipulating resources via http
type HelloWorldController struct {
	hello    *services.HelloService
	firebase *services.FirebaseService
}

// NewHelloWorldController creates an instance of HelloWorldController
func NewHelloWorldController(hello *services.HelloService, firebase *services.FirebaseService) *HelloWorldController {
	return &HelloWorldController{hello: hello, firebase: firebase}
}

// Hello returns a hello world message
func (hc HelloWorldController) Hello(c echo.Context) error {
	uid, _ := c.Get("UID").(string) // UserID set by firebase middleware
	user := hc.firebase.GetUser(uid)
	return c.JSON(http.StatusOK, hc.hello.SayHello(user.DisplayName))
}
