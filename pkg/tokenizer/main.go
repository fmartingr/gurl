package tokenizer

import (
	"log"
	"os"
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

var supporttedSchemes = [...]string{"http"}

const (
	defaultScheme = "http"
	defaultPath   = "/"
	defaultPort   = 80
)

func isSupportedScheme(scheme string) bool {
	for _, a := range supporttedSchemes {
		if a == scheme {
			return true
		}
	}
	return false
}

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
		if !isSupportedScheme(schemeParts[0]) {
			log.Print("Unssuported scheme " + schemeParts[0])
			os.Exit(1)
		}
		uri.Scheme = schemeParts[0]
	}

	urlParts := strings.Split(schemeParts[len(schemeParts)-1], "/")
	uri.Hostname = urlParts[0]
	uri.Path = "/" + strings.Join(urlParts[1:], "/")

	return uri
}
