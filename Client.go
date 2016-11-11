package downloader

import (
	"fmt"
	"net/http"
)

func NewGetClient(url string) (response *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	return response, err
}
