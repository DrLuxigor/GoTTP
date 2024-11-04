package http

import (
	"bufio"
	"fmt"
	"strings"
)

type HttpRequest struct {
	Method  string
	Path    string
	Verion  string
	Headers map[string]string
	Body    []byte
}

func (request *HttpRequest) Print(withBody bool) {
	fmt.Printf("Method: %s\n", request.Method)
	fmt.Printf("Path: %s\n", request.Path)
	fmt.Printf("Version: %s\n", request.Verion)
	fmt.Printf("Headers:\n")
	for key, value := range request.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	if withBody {
		fmt.Printf("Body: %s\n", string(request.Body))
	}
}

func (request *HttpRequest) GetCookies() map[string]string {
	cookies := make(map[string]string)
	cookiesStr := request.Headers["Cookie"]

	for _, cookie := range strings.Split(cookiesStr, ";") {
		cookieParts := strings.SplitN(cookie, "=", 2)
		cookies[cookieParts[0]] = cookieParts[1]
	}
	return cookies
}

func (request *HttpRequest) GetQueryParams() map[string]string {
	params := make(map[string]string)
	queryIndex := strings.Index(request.Path, "?")
	if queryIndex == -1 {
		return params
	}

	if len(request.Path) == queryIndex+1 {
		return params
	}

	query := request.Path[queryIndex+1:]
	qParams := strings.Split(query, "&")
	for _, param := range qParams {
		paramParts := strings.SplitN(param, "=", 2)
		params[paramParts[0]] = paramParts[1]
	}

	return params
}

type HttpResponse struct {
	Version    string
	StatusCode int
	Message    string
	Headers    map[string]string
	Cookies    []string
	Body       []byte
}

func (response *HttpResponse) SetCookie(name string, value string, options map[string]string, maxAge int /*seconds*/, sameSite string /*Strict Lax None*/, secure bool, httpOnly bool, priority string /*Low Medium High*/, domain string, path string) {
	if response.Cookies == nil {
		response.Cookies = make([]string, 0)
	}
	var cookiestring = fmt.Sprintf("%s=%s;", name, value)
	if val, ok := options["max-age"]; ok {
		cookiestring += fmt.Sprintf(" Max-Age=%s;", val)
	}
	if val, ok := options["domain"]; ok {
		cookiestring += fmt.Sprintf(" Domain=%s;", val)
	}
	if val, ok := options["path"]; ok {
		cookiestring += fmt.Sprintf(" Path=%s;", val)
	}
	if _, ok := options["secure"]; ok {
		cookiestring += " Secure;"
	}
	if _, ok := options["http-only"]; ok {
		cookiestring += " HttpOnly;"
	}
	if val, ok := options["same-site"]; ok {
		cookiestring += fmt.Sprintf(" SameSite=%s;", val)
	}
	if val, ok := options["priority"]; ok {
		cookiestring += fmt.Sprintf(" Priority=%s;", val)
	}
	response.Cookies = append(response.Cookies, cookiestring)
}

func ParseHttpRequest(reader *bufio.Reader) *HttpRequest {
	line, err := reader.ReadString('\n')

	if err != nil {
		return nil
	}

	params := strings.Fields(line)

	request := new(HttpRequest)
	request.Method = params[0]
	request.Path = params[1]
	request.Verion = params[2]

	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		if line == "\r\n" || line == "\n" {
			break
		}
		header := strings.SplitN(line, ":", 2)
		headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}
	request.Headers = headers

	if contentLength, ok := headers["Content-Length"]; ok {
		var length int
		fmt.Sscanf(contentLength, "%d", &length)

		bodyBytes := make([]byte, length)
		_, err := reader.Read(bodyBytes)
		if err == nil {
			request.Body = bodyBytes
		}
	}

	return request
}
