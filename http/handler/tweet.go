package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"twitter/creds"
	"twitter/http/client"
)

var (
	NULL_URL_REQ string = "URL REQ interface passed cannot be null."
	baseUri      string = "https://api.twitter.com/2/tweets/search/recent?"
	baseQuery    string = "query=customQuery %s&%s"
	types        string = "-is:retweet -is:reply -is:quote"
	fields              = "tweet.fields=created_at,author_id"
	query        string = fmt.Sprintf(baseUri+baseQuery, types, fields)
)

type TweetURLReq struct {
	Keyword string
}

type TweetHandler struct {
	Req    URLReq
	Client client.Client
}

func (uReq TweetURLReq) URI() (string, error) {
	if &uReq == nil {
		return "", errors.New(NULL_URL_REQ)
	}
	return buildUri(uReq), nil
}

func buildUri(u TweetURLReq) string {
	uri := strings.Replace(query, "customQuery", u.Keyword, 1)
	url := strings.Replace(uri, " ", "%20", 99)
	return url
}

func (h TweetHandler) Handle() (string, error) {
	uri, err := h.Req.URI()
	if err != nil {
		return "", errors.New("There was a problem parsing the URL to request to.")
	}

	req, reqErr := http.NewRequest("GET", uri, nil)
	if reqErr != nil {
		return "", reqErr
	}
	req.Header.Add("Authorization", "Bearer "+creds.Token())

	res, resErr := h.Client.Do(req)
	if resErr != nil {
		return "", resErr
	}
	if res.Status != 200 {
		return not200Res(res.Status, res.Body)
	}

	return res.Body, nil
}

func not200Res(status int, body string) (string, error) {
	switch status {
	case 404:
		return not_found, not_200
	case 401:
		return not_authorized, not_200
	case 403:
		return forbidden, not_200
	case 500:
		return body, not_200
	default:
		return strconv.Itoa(status), not_200
	}
}
