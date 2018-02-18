package main

import (
	"os"

	"github.com/david-torres/go-webapp-vuejs-firebase-boilerplate/controllers"
	"github.com/david-torres/go-webapp-vuejs-firebase-boilerplate/middleware"
	"github.com/david-torres/go-webapp-vuejs-firebase-boilerplate/services"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	// init the web server
	e := echo.New()

	// init app-wide middleware
	e.Pre(mw.RemoveTrailingSlash())
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// security measures
	e.Use(mw.BodyLimit("2M")) // sets maximum request body size
	e.Use(mw.CSRF())          // default protection against session riding
	e.Use(mw.Secure())        // default protection against XSS, content sniffing, clickjacking, and other infections

	// init services
	helloService := services.NewHelloService()
	firebaseService := services.NewFirebaseService(os.Getenv("FIREBASE_CREDENTIALS_PATH"))

	// init controllers
	helloController := controllers.NewHelloWorldController(helloService, firebaseService)

	// init middleware
	auth := middleware.NewFirebaseAuthMiddleware(firebaseService)

	// app entrypoint
	e.File("/", "public/app/dist/index.html")

	// static assets
	e.Static("/static", "public/app/dist/static")

	// helloWorld routes
	helloWorldRoutes := e.Group("/hello", auth.Verify) // hello routes require a valid Firebase id token
	helloWorldRoutes.GET("/world", helloController.Hello)

	// start the server
	e.Logger.Fatal(e.Start(":3000"))
}
