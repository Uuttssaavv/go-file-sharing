# golang-crud-todo

Generate the go.mod file with the command `go mod init go-crud` go-crud here is the package name.
After this run the command `go mod tidy` to add module requirements and generate sum files.
`go mod tidy` also cleans the unused packages.

`main.go` contains main function which is the entry point of the application.
It contains the routes of the application.

```go
func main() {
	SetupAppRouter()
}

func SetupAppRouter() *gin.Engine {

	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.TestMode)

	api := router.Group("api/v1")

	routes.InitAuthRoutes(db, api)

	return router
}
```

The `InitAuthRoutes` takes two parameters `db *gorm.DB` and `route *gin.RouterGroup`.

```go
func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
    // creates the new instance of the Repository
	loginRepository := loginAuth.NewRepositoryLogin(db)

    // creates the new instance of the Service
	loginService := loginAuth.NewServiceLogin(loginRepository)

    // creates the new instance of the Handler
	loginHandler := loginHandler.NewHandlerLogin(loginService)

	route.POST("/login", loginHandler.LoginHandler)
}
```

The call flow for the application looks like

```
loginHandler    ->  loginService        ->      loginRepository
(parse request)     (LoginInput to UserEntity)   (perform DB operations)
```

- The `loginHandler` validates the request body and parse the request body into `LoginInput` struct. Now, if any errors occur, it returns the response with respective error message.
  Now, the `loginHandler` calls the `loginService` that takes the argument `LoginInput`.
- The `loginService` parses the `LoginInput` and uses the values to parse it into the `UserEntity` struct.
  The `loginService` calls the `loginRepository` that takes the argument `UserEntity`.
- The `loginRepository` has the access to the `*gorm.DB`. It performs the operation with the database and returns either of the `(*model.EntityUsers, string)`
  value based upon the success and failure of the operation.
