# gurl

An attempt to learn golang my making a simpler _curl_ client.

## Usage

Retrive the code with `go get` and then run `$GOPATH/bin/gurl <url>`. But why would you...?

```
go get -u github.com/fmartingr/gurl
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
    - [ ] customize headers via cli
    - [ ] keep alive?
    - [ ] more verbs (POST, DELETE, HEAD)
- [ ] SSL support (HTTPS)
