package http

import (
	"strconv"
)

// Header interpretation for a request
type Header struct {
	Name  string
	Value string
}

// Request interpretation
type Request struct {
	Verb     string
	Path     string
	Protocol string
	Headers  []Header
	Body     string
}

// NewRequest Creates a new request
func NewRequest(verb string, path string, headers []Header, body string) Request {
	request := Request{
		Verb:     verb,
		Path:     path,
		Protocol: "HTTP/1.1",
		Headers:  headers,
		Body:     body,
	}

	// If request contains a body, append the content-length header
	if request.Body != "" {
		request.Headers = append(
			request.Headers,
			Header{Name: "Content-Length", Value: strconv.Itoa(len(request.Body))})
	}

	return request
}

// Repr Return the string represntation of the request
func (r Request) Repr() string {
	result := r.Verb + " " + r.Path + " " + r.Protocol + "\r\n"
	for _, h := range r.Headers {
		result += h.Name + ": " + h.Value + "\r\n"
	}
	return result + "\r\n" + r.Body + "\r\n\r\n"
}
