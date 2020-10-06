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

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var extraHeaders arrayFlags
	flag.Var(&extraHeaders, "header", "Extra header to append to request (repatable)")
	displayRequest := flag.Bool("request", false, "Display request in output")
	displayResponse := flag.Bool("response", true, "Display response in output")
	verb := flag.String("verb", "GET", "Verb to use on the request")
	body := flag.String("body", "", "Body to send with request")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Print("Missing URL to connect to")
		flag.Usage()
		os.Exit(cli.ErrorNoUrl)
	}
	uri := tokenizer.Tokenize(flag.Args()[0])

	conn, err := net.Dial("tcp", uri.DialAddress())
	if err != nil {
		log.Print("Error connecting to ", uri.Repr())
		log.Print(err.Error())
		os.Exit(cli.ErrorConnectionOpen)
	}

	headers := []http.Header{
		{Name: "Host", Value: uri.Hostname},
		{Name: "User-Agent", Value: userAgent},
		{Name: "Connection", Value: "close"},
	}

	for _, h := range extraHeaders {
		header := strings.SplitN(h, ":", 2)
		headers = append(headers, http.Header{Name: header[0], Value: header[1]})
	}

	outgoingRequest := http.NewRequest(
		strings.ToUpper(*verb),
		uri.Path,
		headers,
		*body,
	)
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
