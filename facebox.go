package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
)

const billieURL = "https://salvemusic.com.ua/wp-content/uploads/2019/07/Eilish3-1024x683.jpg"

func getBillieURL() *url.URL {
	u, err := url.Parse(billieURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	return u
}

func getBillie() io.Reader {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	request, err := http.NewRequest(http.MethodGet, getBillieURL().String(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	return resp.Body
}
