package tokenizer

import (
	"strconv"
	"strings"
)

type URI struct {
	Scheme      string
	Hostname    string
	Port        int
	Path        string
	QueryParams string
}

const (
	defaultScheme = "http"
	defaultPath   = "/"
	defaultPort   = 80
)

func (u URI) Repr() string {
	return u.Scheme + "://" + u.Hostname + u.Path
}

func (u URI) DialAddress() string {
	return u.Hostname + ":" + strconv.Itoa(u.Port)
}

func Tokenize(url string) URI {
	uri := URI{Scheme: defaultScheme, Path: defaultPath, Port: defaultPort}
	schemeParts := strings.Split(url, "://")

	if len(schemeParts) > 1 {
		uri.Scheme = schemeParts[0]
	}

	urlParts := strings.Split(schemeParts[len(schemeParts)-1], "/")
	uri.Hostname = urlParts[0]
	uri.Path = "/" + strings.Join(urlParts[1:], "/")

	return uri
}
