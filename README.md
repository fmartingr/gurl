# gurl

An attempt to learn golang my making a simpler _curl_ client.

## Usage

Retrive the code with `go get` and then run `$GOPATH/bin/gurl <url>`. But why would you...?

```
$ go get -u github.com/fmartingr/gurl
$ gurl -h
Usage of gurl:
  -body string
        Body to send with request
  -header value
        Extra header to append to request (repatable)
  -request
        Display request in output
  -response
        Display response in output (default true)
  -verb string
        Verb to use on the request (default "GET")
```

## Roadmap

- [ ] basic url tokenizer
    - [x] scheme
    - [x] hostname
    - [ ] port
    - [ ] username
    - [ ] password
    - [x] path
    - [ ] query parameters
- [x] http request
    - [x] strucs
    - [x] customize headers via cli
    - [ ] keep alive?
    - [x] more verbs (POST, DELETE, HEAD)
    - [x] Simple body
- [x] Flag to display request
- [ ] SSL support (HTTPS)
