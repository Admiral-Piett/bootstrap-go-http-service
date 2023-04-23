# Bootstrap a RESTful HTTP Service in Golang
This is a basic HTTP Service framework set up designed for Go RESTful API applications.  

It implements 2 routes by default though both return errors until they are refactored based on the user's wishes.
1. `POST /api/auth`
2. `GET|POST|PUT|PATCH|DELETE /.*` (Captures any other route not already defined)

It will log each response in the format `<status_code>: <request_method> <request_uri>` (eg. `404: POST /api/invalid`)

It also implements an optional pattern for easily routing to individual HTTP specific functions based on the 
requested HTTP Method.  See `handlers.NotFoundHandler` (in `app.handlers.not_found_handler.go`).

## To Run:
* `go get -u ./...`
* `go run app/cmd/main.go`

## Before Production
You may want to adjust 2 things before deploying
1. Implement your own Authentication Functionality.  The current `AuthHandler` isn't implemented.
2. Restrict your CORS requests - unless you're comfortable fielding every request from every host ever, you may want 
   to tighten up the headers it returns on Browsers' pre-flight CORS requests.  It currently allows requests from 
   any host and all mainstream HTTP Methods.
