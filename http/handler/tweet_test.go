package handler

import (
	"fmt"
	"testing"
	"twitter/creds"

	"github.com/onsi/gomega"
)

func TestResponseOnTwitterAPI(t *testing.T) {
	creds.Init("../../creds/auth.json")

	gmg := gomega.NewGomegaWithT(t)

	handler := TweetHandler{TweetURLReq{"Test driven development"}}
	res, err := handler.handle()

	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	fmt.Println(res)
}
