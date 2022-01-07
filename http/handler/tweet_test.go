package handler

import (
	"fmt"
	"testing"
	"twitter/creds"
	"twitter/http/client"

	"github.com/onsi/gomega"
)

// This test isn't really testing anything. It's just there to see if
// Twitter's API has changed or not.
func TestResponseOnTwitterAPI(t *testing.T) {
	creds.Init("../../creds/auth.json")

	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{"Test driven development"}, client.Default()}
	res, err := handler.Handle()

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	fmt.Println(res)
}

func TestRetErrIfHandlerNotRecieve200(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}, stubClient{status: 300}}
	_, err := handler.Handle()

	gmg.Ω(err).Should(gomega.HaveOccurred())
}

func TestResReturnsNotFoundStatusThenHandlerReturnsNotFoundString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}, stubClient{status: 404}}
	res, err := handler.Handle()

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(not_found))
}

func TestResReturnsUnAuthorizedStatusThenHandlerReturnsNotAuthorizedString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}, stubClient{status: 401}}
	res, err := handler.Handle()

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(not_authorized))
}

func TestResReturnsForbiddenStatusThenHandlerReturnsForbiddenString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}, stubClient{status: 403}}
	res, err := handler.Handle()

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(forbidden))
}

func TestResReturn500AsStatusThenHandlerReturnsWhateverClientReturns(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	rsBody := "Internal server problem"
	handler := TweetHandler{TweetURLReq{""}, stubClient{body: rsBody, status: 500}}

	res, err := handler.Handle()

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(rsBody))
}
