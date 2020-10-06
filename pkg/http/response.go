package http

import (
	"strconv"
)

// Response internal representation
type Response struct {
	Protocol   string
	StatusCode int
	StatusText string
	Headers    []Header
	Body       string
}

// Repr Response text representation
func (r Response) Repr() string {
	result := r.Protocol + " " + strconv.Itoa(r.StatusCode) + " " + r.StatusText + "\r\n"
	for _, h := range r.Headers {
		result += h.Name + ": " + h.Value + "\r\n"
	}
	result += "\r\n" + r.Body
	return result + "\r\n"
}
