package gottp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"path"
	"strings"
)

type GoTTPFunc func(request *HttpRequest, response *HttpResponse)

type GoTTPServer struct {
	Port           int
	MaxRequestSize int64
	Gets           map[string]GoTTPFunc
	Posts          map[string]GoTTPFunc
	Puts           map[string]GoTTPFunc
	Updates        map[string]GoTTPFunc
	Deletes        map[string]GoTTPFunc
}

func CreateApp() *GoTTPServer {
	app := new(GoTTPServer)
	app.Port = 80
	app.MaxRequestSize = 1 << 24
	app.Gets = make(map[string]GoTTPFunc)
	app.Posts = make(map[string]GoTTPFunc)
	app.Puts = make(map[string]GoTTPFunc)
	app.Updates = make(map[string]GoTTPFunc)
	app.Deletes = make(map[string]GoTTPFunc)
	return app
}

func (s *GoTTPServer) Start() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))

	if err != nil {
		fmt.Printf("Could not start server: %s\n", err)
		return
	}

	fmt.Printf("GoTTP listening on port %d\n", s.Port)

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn, s)
	}
}

func handleConnection(conn net.Conn, app *GoTTPServer) {
	defer conn.Close()

	limitreader := io.LimitReader(conn, app.MaxRequestSize)
	reader := bufio.NewReader(limitreader)

	request := ParseHttpRequest(reader)

	response := new(HttpResponse)
	response.Version = request.Version
	response.StatusCode = 200
	response.Message = "OK"
	response.Headers = make(map[string]string)
	response.Cookies = make([]string, 0)

	fmt.Println(request.Path)
	//TODO extract and format path of request
	url, err := url.Parse(request.Path)
	if err != nil {
		fmt.Printf("Could not parse url %s\nError: %s\n", request.Path, err)
		response.StatusCode = 503
		response.Message = "Bad Request"
		conn.Write([]byte(response.BuildResponseHeader()))
		return
	}
	cleanPath := path.Clean(url.Path)

	//Check if path is a static file and serve it
	if stat, err := os.Stat(fmt.Sprintf("static/%s", cleanPath)); err == nil && !stat.IsDir() {
		file, err := os.ReadFile(fmt.Sprintf("static/%s", cleanPath))
		//file exists serve file
		if err == nil {
			response.StatusCode = 200
			response.Message = "OK"
			response.Body = file
			extension := path.Ext(cleanPath)
			response.ContentType = GetContentType(extension)

			conn.Write([]byte(response.BuildResponseHeader()))
			if response.Body != nil {
				conn.Write(response.Body)
			}
			return
			//file doesnt exist and is the 404 pages, write 404 error
		} else if strings.ToLower(cleanPath) == "404.html" {
			response.StatusCode = 404
			response.Message = "Not Found"
			conn.Write([]byte(response.BuildResponseHeader()))
			return
			//file simply doesnt exist, redirect to 404 page
		} else {
			response.StatusCode = 301
			response.Message = "Moved Permanently"
			response.Headers["Location"] = "/404.html"
			conn.Write([]byte(response.BuildResponseHeader()))
			return
		}
	}

	//Handle server functions
	gottpfunction := findFunc(request.Method, cleanPath, app)

	if gottpfunction == nil {
		response.StatusCode = 301
		response.Message = "Moved Permanently"
		response.Headers["Location"] = "/404.html"
		conn.Write([]byte(response.BuildResponseHeader()))
		return
	}

	gottpfunction(request, response)

	conn.Write([]byte(response.BuildResponseHeader()))
	if response.Body != nil {
		conn.Write(response.Body)
	}
}

func findFunc(method string, path string, app *GoTTPServer) GoTTPFunc {
	switch method {
	case "GET":
		return app.Gets[path]
	case "PUT":
		return app.Puts[path]
	case "UPDATE":
		return app.Updates[path]
	case "DELETE":
		return app.Deletes[path]
	}
	return nil
}

func (s *GoTTPServer) Get(path string, f GoTTPFunc) {
	s.Gets[path] = f
}

func (s *GoTTPServer) Post(path string, f GoTTPFunc) {
	s.Posts[path] = f
}

func (s *GoTTPServer) Put(path string, f GoTTPFunc) {
	s.Puts[path] = f
}

func (s *GoTTPServer) Update(path string, f GoTTPFunc) {
	s.Updates[path] = f
}

func (s *GoTTPServer) Delete(path string, f GoTTPFunc) {
	s.Deletes[path] = f
}
