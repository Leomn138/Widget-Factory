# Go Widget Factory REST API
A RESTful API example for a simple widget factory application with Go

## Installation & Run
```bash
# Download this project
go get github.com/Leomn138/Widget-Factory
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/Leomn138/Widget-Factory/blob/master/config/config.go)
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
cd Widget-Factory
go build
./Widget-Factory

# API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── users.go     // APIs for User model
│   │   └── widgets.go   // APIs for Widget model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /users
* `GET` : Get all users

#### /users/:id
* `GET` : Get a user

#### /widgets
* `GET` : Get all widgets
* `POST` : Create a new widget

#### /widgets/:id
* `GET` : Get a widget
* `PUT` : Update a widget

## Todo

- [x] Support basic REST APIs.
- [x] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
