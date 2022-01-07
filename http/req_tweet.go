package http

import (
	"twitter/http/client"
	"twitter/http/handler"
	"twitter/marshal"
)

func ReqTweet(h handler.Handler, cl client.Client) (*marshal.Tweet, error) {
	res, err := h.Handle(cl)
	if err != nil {
		return nil, err
	}
	return marshal.Parse(res)
}

func ReqTweets(h handler.Handler, cl client.Client) (*marshal.Tweets, error) {
	res, err := h.Handle(cl)
	if err != nil {
		return nil, err
	}
	return marshal.ParseTweets(res)
}
