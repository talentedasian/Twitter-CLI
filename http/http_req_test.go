package http

import (
	"testing"
	"twitter/marshal"

	"github.com/onsi/gomega"
)

type StubTweetHandler struct {
	cnt string
}

func (h StubTweetHandler) handle() (string, error) {
	return h.cnt, nil
}

func TestReqTweetReturnTweetStruct(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	resRaw :=
		`
		{
			 "text": "Irrelevant content",
			 "author_id": "2244994945321",
			 "id": "213121",
			 "conversation_id": "213213212332",
			 "created_at": "2021-03-19T19:59:10.000Z"
		}
	`

	tw, _ := marshal.Parse(resRaw)

	handler := StubTweetHandler{resRaw}
	res, err := ReqTweet(handler)

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(tw))
}

func TestReqTweetsReturnSliceTweetStruct(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	resRaw := `
		{
			"data": [
				 {
					 "text": "Irrelevant content",
					 "author_id": "2244994945321",
					 "id": "213121",
				 "conversation_id": "213213212332",
					 "created_at": "2021-03-19T19:59:10.000Z"
				 },
				 {
					 "text": "Another irrelevant content",
					 "author_id": "092813015",
					 "id": "2190373285y2",
					 "lang": "en",
					 "conversation_id": "0912109320",
					 "created_at": "2021-03-19T19:59:10.000Z"
				 }
			]
		}
	`

	tw, _ := marshal.ParseTweets(resRaw)

	handler := StubTweetHandler{resRaw}
	res, err := ReqTweets(handler)

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(tw))
}

func TestMalformedJsonRetErrAndNilResponseOnTweets(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	resRaw := `
		 {
			 "text": "Irrelevant content:,
			 "author_id": "2244994945321",
			 "id": "213121",
			 "conversation_id": "213213212332",
			 "created_at": "2021-03-19T19:59:10.000Z"
		 }
	`

	h := StubTweetHandler{resRaw}
	res, err := ReqTweets(h)

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.BeNil())
}

func TestMalformedJsonRetErrAndNilResponseOnTweet(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	resRaw := `
		{
			"data": [
				 {
					 "text": "Irrelevant content",
					 "author_id": "2244994945321",
					 "id": "213121",
				 "conversation_id": "213213212332",
					 "created_at": "2021-03-19T19:59:10.000Z"
				 }
				 {
					 "text": "Another irrelevant content",
					 "author_id": "092813015",
					 "id": "2190373285y2",
					 "lang": "en",
					 "conversation_id": "0912109320",
					 "created_at": "2021-03-19T19:59:10.000Z"
				 }
			]
		}
	`

	h := StubTweetHandler{resRaw}
	res, err := ReqTweets(h)

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.BeNil())
}
