package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	CLIENT              *http.Client = &http.Client{}
	INVALID_NET_REQUEST error        = errors.New("The HTTP Request passed is not from net/http")
)

type Client interface {
	Do(args ...interface{}) (*Response, error)
}

type Response struct {
	Body   string
	Status int
}

type NetHttpClient struct {
	Client *http.Client
}

func Default() *NetHttpClient {
	return &NetHttpClient{CLIENT}
}

func (client *NetHttpClient) Do(args ...interface{}) (*Response, error) {
	if req, ok := args[0].(*http.Request); ok {
		res, resErr := client.Client.Do(req)
		defer res.Body.Close()
		if resErr != nil {
			return nil, resErr
		}

		resByte, strErr := ioutil.ReadAll(res.Body)
		if strErr != nil {
			return nil, strErr
		}

		return &Response{string(resByte), res.StatusCode}, nil
	}

	return nil, INVALID_NET_REQUEST
}
