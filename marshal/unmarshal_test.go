package marshal

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestMarshalFromResponseFromTwitter(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	tw := "Looking to get started with the Twitter API but new to APIs in general?"
	aId := "2244994945"
	tId := "1373001119480344583"
	tOn := "2021-03-19T19:59:10.000Z"
	marRes := Tweet{tw, tId, aId, tOn}

	twRes := `	 
			 {
				 "text": "Looking to get started with the Twitter API but new to APIs in general?",
				 "author_id": "2244994945",
				 "id": "1373001119480344583",
				 "lang": "en",
				 "conversation_id": "1373001119480344583",
				 "created_at": "2021-03-19T19:59:10.000Z"
			 }
	`

	mar, _ := Parse(twRes)

	gmg.Ω(*mar).Should(gomega.Equal(marRes))
}

func TestParseResponseFromTwitterTweets(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	tw := "Irrelevant content"
	aId := "2244994945321"
	tId := "213121"
	tOn := "2021-03-19T19:59:10.000Z"
	marRes := Tweet{tw, tId, aId, tOn}
	tw2 := "Another irrelevant content"
	aId2 := "092813015"
	tId2 := "2190373285y2"
	tOn2 := "2021-03-19T19:59:10.000Z"
	marRes2 := Tweet{tw2, tId2, aId2, tOn2}

	twRes := `
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

	mar, err := ParseTweets(twRes)

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	gmg.Ω(mar.Tweets).Should(gomega.ContainElements(marRes, marRes2))
}
