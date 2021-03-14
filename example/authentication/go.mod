module github.com/muhfaris/facebook-sdk-go/example/authentication

go 1.14

require (
	github.com/labstack/echo/v4 v4.2.1
	github.com/muhfaris/facebook-sdk-go v0.0.0-20200708085008-e6a63a17767d
	github.com/muhfaris/request v0.0.3 // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93 // indirect
)

replace github.com/muhfaris/facebook-sdk-go => ./../../
