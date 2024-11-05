# GoTTP

## A from scratch implementation of an HTTP server in Go.

```
package main

import (
	"lukaskofler.dev/gottp/src/pkg/gottp"
)

func main() {
	app := gottp.CreateApp()
	app.Port = 8050

	app.Get("/", func(request *gottp.HttpRequest, response *gottp.HttpResponse) {
		response.StatusCode = 200
		response.Message = "OK"
		response.ContentType = "text/plain"
		response.Body = []byte("Hello World!")
	})

	app.Start()
}
```

## METHODS

It is possible to add routes for **GET**, **PUT**, **POST** and **DELETE** as seen above.


## Headers

Headers can be read and set
```
val, err := request.Headers["Host"]
response.Headers["Content-Type"] = "text/plain"
//this one could also be done with response.ContentType = "text/plain"
```

## Cookies

Cookies can be read and set
```
cookies = request.GetCookies()
response.SetCookie("mycookie", "value", map[string]string{"max-age": "3600"})
```
The cookie options are optional and the following can be set: **max-age**, **domain**, **path**, **secure**, **http-only**, **same-site**, **priority**.
the secure and http-only options don't care about their value and will be set if present in the map.

## Query parameters

The query parameters of a request can be read
```
request.GetQueryParameters()
```

## Static files

Static files can be served from a folder named "static" in the project root folder.
If the path is recognized as a file, the file will be served if present, otherwise the server will redirect to /404.html.
If static/404.html doesn't exist it, the server will respond with a 404 error.

Normal paths will be handled by the functions the user defines, if a path is not found, the server will redirect to /404.html.

