package http

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
}

// NewRequest Creates a new request
func NewRequest(path string, headers []Header) Request {
	return Request{
		Verb:     "GET",
		Path:     path,
		Protocol: "HTTP/1.1",
		Headers:  headers,
	}
}

// Repr Return the string represntation of the request
func (r Request) Repr() string {
	result := r.Verb + " " + r.Path + " " + r.Protocol + "\r\n"
	for _, h := range r.Headers {
		result += h.Name + ": " + h.Value + "\r\n"
	}
	return result + "\r\n"
}
