package http

import (
	"twitter/http/handler"
	"twitter/marshal"
)

func ReqTweet(h handler.Handler) (*marshal.Tweet, error) {
	res, err := h.Handle()
	if err != nil {
		return nil, err
	}
	return marshal.Parse(res)
}

func ReqTweets(h handler.Handler) (*marshal.Tweets, error) {
	res, err := h.Handle()
	if err != nil {
		return nil, err
	}
	return marshal.ParseTweets(res)
}
