# File Sharing with Golang

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

### Why are we using \*pointer in the return type instead of value of the object?

Snippet example:

```go
// why are we using this
func (r *repository) LoginRepository(input *models.UserEntity) (*models.UserEntity, int) {}

// instead of this
func (r *repository) LoginRepository(input models.UserEntity) (models.UserEntity, int) {}

```

In the first function, the return type is `*models.UserEntity`. It means it returns the pointer to the `UserEnitity` type. 
This allows us to modify the object directly, and any changes made to the returned object will affect the original object. This can be useful because we want to maintain a single shared instance of `UserEnitity` and to avoid unnecessary object copies.

In the second function, the return type is `models.UserEntity`.
It means new object will be created each time you call the function, and any modifications made to the returned `UserEntity` object will not affect the original object passed as an argument.
**Note:** This can be useful if you want to create multiple independent instances of `UserEntity`