package handler

import (
	"fmt"
	"testing"
	"twitter/creds"
	"twitter/http/client"

	"github.com/onsi/gomega"
)

func TestResponseOnTwitterAPI(t *testing.T) {
	creds.Init("../../creds/auth.json")

	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{"Test driven development"}}
	res, err := handler.handle(client.Default())

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	fmt.Println(res)
}

func TestRetErrIfHandlerNotRecieve200(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}}
	_, err := handler.handle(stubClient{status: 404})

	gmg.Ω(err).Should(gomega.HaveOccurred())
}

func TestResReturnsNotFoundStatusHandlerReturnsNotFoundString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}}
	res, err := handler.handle(stubClient{status: 404})

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(not_found))
}

func TestResReturnsUnAuthorizedStatusHandlerReturnsNotAuthorizedString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}}
	res, err := handler.handle(stubClient{status: 401})

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(not_authorized))
}

func TestResReturnsForbiddenStatusHandlerReturnsForbiddenString(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}}
	res, err := handler.handle(stubClient{status: 403})

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(forbidden))
}

func TestResReturn500AsStatusHandlerReturnsWhateverClientReturns(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{""}}
	rsBody := "Internal server problem"

	res, err := handler.handle(stubClient{body: rsBody, status: 500})

	gmg.Ω(err).Should(gomega.HaveOccurred())
	gmg.Ω(res).Should(gomega.Equal(rsBody))
}
