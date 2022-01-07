package handler

import "twitter/http/client"

type stubClient struct {
	body   string
	status int
	err    error
}

func (stbCl stubClient) okBody(body string) stubClient {
	return stubClient{body, 200, nil}
}

func (stbCl stubClient) Do(args ...interface{}) (*client.Response, error) {
	return &client.Response{stbCl.body, stbCl.status}, stbCl.err
}
