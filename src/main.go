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
