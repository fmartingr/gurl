package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	userAgent      = "Gurl/0.0.1a1"
	port           = 80
	connectionType = "tcp"
)

type Header struct {
	name  string
	value string
}

type Request struct {
	verb     string
	path     string
	protocol string
	headers  []Header
}

func NewRequest(headers []Header) Request {
	return Request{
		verb:     "GET",
		path:     "/",
		protocol: "HTTP/1.1",
		headers:  headers,
	}
}

func (r Request) Repr() string {
	result := r.verb + " " + r.path + " " + r.protocol + "\r\n"
	for _, h := range r.headers {
		result += h.name + ": " + h.value + "\r\n"
	}
	return result + "\r\n"
}

type Response struct {
	protocol   string
	statusCode int
	statusText string
	headers    []Header
	body       string
}

func (r Response) Repr() string {
	result := r.protocol + " " + strconv.Itoa(r.statusCode) + " " + r.statusText + "\r\n"
	for _, h := range r.headers {
		result += h.name + ": " + h.value + "\r\n"
	}
	result += "\r\n" + r.body
	return result + "\r\n"
}

var displayRequest bool = false

func init() {
	flag.BoolVar(&displayRequest, "request", true, "Display request")
}

func main() {
	arguments := os.Args[1:]

	hostname := arguments[0]

	conn, err := net.Dial(connectionType, hostname+":"+strconv.Itoa(port))
	if err != nil {
		log.Print("Error connecting:", err.Error())
		os.Exit(1)
	}

	request := NewRequest([]Header{{"Host", hostname}, {"User-Agent", userAgent}, {"Connection", "close"}})
	if displayRequest {
		fmt.Println(request.Repr())
	}

	conn.Write([]byte(request.Repr()))

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	response := Response{}
	readingHeaders := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingHeaders = false
		}

		if response.statusCode == 0 {
			firstLine := strings.Split(line, " ")
			response.protocol = firstLine[0]
			response.statusCode, _ = strconv.Atoi(firstLine[1])
			response.statusText = strings.Join(firstLine[2:], " ")
			readingHeaders = true
			continue
		} else if readingHeaders {
			splitLine := strings.Split(line, ":")
			header := Header{splitLine[0], splitLine[1]}
			response.headers = append(response.headers, header)
		} else {
			response.body += line
		}

	}

	fmt.Println(response.Repr())

	conn.Close()
}
