# Go Web Application: VueJS Firebase Quickstart

The goal of this project is to provide a simple skeleton project for a web application using a Go backend utilizing [Firebase](https://firebase.google.com/docs/auth) for authentication. Routing and other web application niceities are provided by [Echo](https://echo.labstack.com) web framework. Firebase authentication middleware provides the abilty to secure routes.

Frontend is [VueJS](https://vuejs.org) with [Bulma](https://bulma.io) and [Firebase UI](https://github.com/firebase/firebaseui-web). A basic authentication example with a call to a secured route are included.

See .env.example for environment variables which must be set for builds and running the app to work correctly.

Include environment variables:

`source .env`

Install frontend deps:

`cd public/app`

`npm install`

Build frontend:

`npm run build`

Install Go deps:

`dep ensure`

Build and install:

`go build -o server`

Run:

`./server`