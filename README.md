# File Sharing with Golang

This project is designed to help beginners learn backend development using Golang. It covers a wide range of beginner concepts and techniques in the backend development space. Topics covered in this project include:

- Setting up a database connection using `GORM`
- Building REST APIs using the `GIN framework`
- Implementing CRUD operations using `GORM` for data manipulation
- Understanding and implementing table associations with foreign keys
- Efficiently populating fields using the `Preload()` function
- Serialization of data into JSON format
- Customizing response and error formats for better user experience
- Uploading files to Cloudinary for storage and retrieval
- Utilizing .env files for secure configuration management
- Following **Clean Architecture** principles for maintainability and scalability

---

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


### Relationship
```go
type FileEntity struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Type      string     `gorm:"not null" json:"type"`
	Name      string     `gorm:"not null" json:"name"`
	Url       string     `gorm:"not null" json:"url"`
	AccessKey string     `gorm:"" json:"access_key"`
	CreatedAt time.Time  `gorm:"" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"" json:"updatedAt"`
	UserID    uint       `gorm:"foreignkey:UserRefer" json:"-"`
	User      UserEntity `gorm:"foreignkey:UserRefer" json:"user"`
}

type UserEntity struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"column:username;unique;not null"`
	Email     string `gorm:"column:email;unique;not null"`
	Image     string `gorm:"column:image"`
	Password  string `gorm:"column:password;not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
```
In the `FileEntity` struct, there is a foreign key field `UserID` of type uint. This field establishes the relationship between `FileEntity` and `UserEntity` using the **foreign key** constraint. The foreign key constraint ensures data integrity between the two tables when inserting or updating records.

### Populating Associations
To populate the foreign key value in GORM based on the given context, we can use the `Preload` function with the relationship defined in the `FileModel` struct. Here's an example:
```go
var files []models.FileModel
db.Select("*").Where("user_id=?", userId).Find(&files)
db.Preload("User").Find(&files)
```
In this case, the `Preload("User")` call specifies that we want to load the associated `UserEntity` record using the foreign key relationship defined in the `FileModel` struct. This will populate the User field in the `FileModel` instance with the associated `UserEntity`.
### Utilities

In the file `utils/json.go` the function definition looks like

```go
func ObjectToJson[T any](object interface{}, data *T) {
	// 
}
```

Which means the function `ObjectToJson` takes the parameter of type `T` and returns the pointer to address of same type.

**Usecase:**
```go
var data register.RegisterResponse
utils.ObjectToJson(resultLogin, &data)
```
So, when we want to convert the object of type `RegisterResponse` to json. We call the function `ObjectToJson` and it maps `resultLogin` to the object of type `register.RegisterResponse` and assign the value to `data`.