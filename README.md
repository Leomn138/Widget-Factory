# Go Widget Factory REST API
A RESTful API example for a widget factory application with Go

## Installation & Run
```bash
# Download this project
go get github.com/Leomn138/Widget-Factory
```
Set up the latest version of Apache CouchDb (http://couchdb.apache.org/#download)

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/Leomn138/Widget-Factory/config/config.go).

You should also set the port the application will be running and the auth secret.

```go
func GetConfig() *Config {

	return &Config{
		UserDBConfig: &DBConfig{
			Name: "users",
			Port: 5984,
			Host: "127.0.0.1",
			Protocol: "http",
		},
		WidgetDBConfig: &DBConfig{
			Name: "widgets",
			Port: 5984,
			Host: "127.0.0.1",
			Protocol: "http",
		},
		Port: ":8000",
		Auth: &Auth {
			Secret: "secret",
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
│   │   ├── auth.go      // Auth functions
│   │   ├── users.go     // APIs for Users
│   │   └── widgets.go   // APIs for Widgets
│   ├── repository
│   │   └── couchdb.go   // CouchDb operations for our application
│   ├── common
│   │   └──common.go	 // Common response functions
│   └── config
│       └── config.go        // Configuration
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

