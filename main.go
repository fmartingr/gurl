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

	http "github.com/fmartingr/gurl/pkg/http"
)

const (
	userAgent      = "Gurl/0.0.1a1"
	port           = 80
	connectionType = "tcp"
)

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

	outgoingRequest := http.NewRequest([]http.Header{
		{Name: "Host", Value: hostname},
		{Name: "User-Agent", Value: userAgent},
		{Name: "Connection", Value: "close"},
	})
	if displayRequest {
		fmt.Println(outgoingRequest.Repr())
	}

	conn.Write([]byte(outgoingRequest.Repr()))

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	response := http.Response{}
	readingHeaders := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingHeaders = false
		}

		if response.StatusCode == 0 {
			firstLine := strings.Split(line, " ")
			response.Protocol = firstLine[0]
			response.StatusCode, _ = strconv.Atoi(firstLine[1])
			response.StatusText = strings.Join(firstLine[2:], " ")
			readingHeaders = true
			continue
		} else if readingHeaders {
			splitLine := strings.Split(line, ":")
			header := http.Header{Name: splitLine[0], Value: splitLine[1]}
			response.Headers = append(response.Headers, header)
		} else {
			response.Body += line
		}

	}

	fmt.Println(response.Repr())

	conn.Close()
}
