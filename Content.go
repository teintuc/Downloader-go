/**
 * Downadload the content of an url and return it has a byte
 */
package downloader

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func DownloadContent(url string) (body []byte, err error) {
	response, err := NewGetClient(url)
	defer response.Body.Close()

	if err == nil && response.StatusCode == http.StatusOK {
		body, err = ioutil.ReadAll(response.Body)
	} else {
		err = errors.New("Error: " + string(response.Status))
	}
	return body, err
}
