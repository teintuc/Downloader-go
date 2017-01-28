package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type PassThru struct {
	io.Reader
	total         int64
	contentLength int64
	printProgress bool
}

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

func progress(current, total, bar_size int) {
	amount := int(float32(current) / (float32(total) / float32(bar_size)))
	remain := bar_size - amount
	progress := (current * 100) / total

	bar := " [" + strings.Repeat("=", amount) + strings.Repeat(" ", remain) + "] " + strconv.Itoa(current) + " " + strconv.Itoa(progress) + "%"
	fmt.Println(bar)
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	if err == nil && pt.printProgress == true && pt.contentLength > 0 {
		pt.total += int64(n)
		progress(int(pt.total), int(pt.contentLength), 25)
	}
	return n, err
}

func File(url string, filename string) (err error) {
	/* Create a GET resousrce to download the content */
	response, err := NewGetClient(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	/* Create the targeted file */
	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	fmt.Println("Downloading to: " + filename)
	/* Copy the content into the targeted file */
	src := &PassThru{Reader: response.Body, contentLength: response.ContentLength, printProgress: true}
	_, err = io.Copy(fd, src)
	if err == nil {
		fmt.Println("Done....")
	}
	return err
}
