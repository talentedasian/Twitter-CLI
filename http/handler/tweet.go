package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	NULL_URL_REQ string = "URL REQ interface passed cannot be null."
)

type Handler interface {
	handle() (string, error)
}

type URLReq interface {
	URI() (string, error)
}

type TweetURLReq struct {
	keyword string
}

type TweetHandler struct {
	req URLReq
}

func (uReq TweetURLReq) URI() (string, error) {
	if &uReq == nil {
		return "", errors.New(NULL_URL_REQ)
	}
	return fmt.Sprintf("https://api.twitter.com/2/tweets/search?query=%s", uReq.keyword), nil
}

func (h TweetHandler) handle() (string, error) {
	uri, err := h.req.URI()
	if err != nil {
		return "", errors.New("There was a problem parsing the URL to request to.")
	}

	req, reqErr := http.NewRequest("GET", uri, nil)
	if reqErr != nil {
		return "", reqErr
	}

	client := &http.Client{}

	res, resErr := client.Do(req)
	if resErr != nil {
		return "", resErr
	}

	defer res.Body.Close()

	resByte, strErr := ioutil.ReadAll(res.Body)
	if strErr != nil {
		return "", strErr
	}

	return string(resByte), nil
}
