package marshal

import (
	"encoding/json"
)

type Tweet struct {
	Content   string `json:"text"`
	Id        string `json:"id"`
	AuthorId  string `json:"author_id"`
	TweetedOn string `json:"created_at"`
}

type Tweets struct {
	Tweets []Tweet `json:"data"`
}

func Parse(j string) (*Tweet, error) {
	var tw Tweet

	if err := json.Unmarshal([]byte(j), &tw); err != nil {
		return nil, err
	}

	return &tw, nil
}

func ParseTweets(j string) (*Tweets, error) {
	var tw Tweets

	if err := json.Unmarshal([]byte(j), &tw); err != nil {
		return nil, err
	}

	return &tw, nil
}
