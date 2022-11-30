package services

import (
	"Go-cup/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Response   string
}

func Get(url string, header http.Header) (Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	if header == nil {
		config, err := util.LoadConfig()
		if err != nil {
			log.Fatalln(err)
			return Response{}, err
		}
		req.Header.Add("X-Auth-Token", config.ApiKey)
	} else {
		req.Header = header
	}

	client := http.Client{Timeout: time.Duration(time.Second * 5)}

	resp, reqErr := client.Do(req)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	responseBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	response := Response{
		StatusCode: 200,
		Response:   fmt.Sprintf("%s", responseBody),
	}

	return response, nil
}
