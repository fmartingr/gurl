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

	"github.com/fmartingr/gurl/pkg/cli"
	http "github.com/fmartingr/gurl/pkg/http"
	"github.com/fmartingr/gurl/pkg/tokenizer"
)

const (
	userAgent = "Gurl/0.0.1a1"
)

func main() {
	displayRequest := flag.Bool("request", false, "Display request")
	displayResponse := flag.Bool("response", true, "Display response")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Print("Missing URL to connect to")
		os.Exit(cli.ErrorNoUrl)
	}
	uri := tokenizer.Tokenize(flag.Args()[0])

	conn, err := net.Dial("tcp", uri.DialAddress())
	if err != nil {
		log.Print("Error connecting:", err.Error())
		os.Exit(cli.ErrorConnectionOpen)
	}

	outgoingRequest := http.NewRequest(
		uri.Path,
		[]http.Header{
			{Name: "Host", Value: uri.Hostname},
			{Name: "User-Agent", Value: userAgent},
			{Name: "Connection", Value: "close"},
		})
	if *displayRequest {
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
			splitLine := strings.SplitN(line, ":", 2)
			header := http.Header{Name: splitLine[0], Value: splitLine[1]}
			response.Headers = append(response.Headers, header)
		} else {
			response.Body += line
		}

	}

	if *displayResponse {
		fmt.Println(response.Repr())
	}

	conn.Close()
}
