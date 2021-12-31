package http

import (
	"twitter/marshal"
)

type Handler interface {
	handle() (string, error)
}

func ReqTweet(h Handler) (*marshal.Tweet, error) {
	res, err := h.handle()
	if err != nil {
		return nil, err
	}
	return marshal.Parse(res)
}

func ReqTweets(h Handler) (*marshal.Tweets, error) {
	res, err := h.handle()
	if err != nil {
		return nil, err
	}
	return marshal.ParseTweets(res)
}
