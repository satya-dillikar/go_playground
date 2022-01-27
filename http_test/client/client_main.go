package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	log.Printf("UpperCase: word: %v", word)
	if err != nil {
		return "", errors.Wrap(err, "unable to complete Get request")
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "unable to read response data")
	}
	result := string(out)
	log.Printf("UpperCase : result: %v", result)
	return result, nil
}

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("client connected on :1234")

	url := "http://127.0.0.1:1234"
	c := NewClient(url)
	expected := "dummydata"
	//UpperCase_real(url, expected)
	c.UpperCase(expected)

}
