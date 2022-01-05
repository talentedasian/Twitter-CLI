package creds

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/onsi/gomega"
)

func TestNonExistingFileOnExpectedDirCausePanic(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	err := Init("../not_found.json")
	gmg.Ω(err).Should(gomega.Equal(CONFIG_NOT_FOUND))
}

func TestSuccessfulInitLoadAuthorizationCreds(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	var cred auth
	authFl, _ := os.Open("auth.json")

	byteV, _ := io.ReadAll(authFl)
	json.Unmarshal(byteV, &cred)

	err := Init("auth.json")
	gmg.Ω(err).ShouldNot(gomega.HaveOccurred())
	gmg.Ω(cred.Token).Should(gomega.Equal(Token()))
	gmg.Ω(cred.Key).Should(gomega.Equal(Key()))
	gmg.Ω(cred.Secret).Should(gomega.Equal(Secret()))

	defer UnInit()
	defer authFl.Close()
}

func TestAccessAuthCredsPanicIfNotInitialized(t *testing.T) {
	gmg := gomega.NewGomegaWithT(t)

	gmg.Ω(func() { Token() }).Should(gomega.Panic())
	gmg.Ω(func() { Key() }).Should(gomega.Panic())
	gmg.Ω(func() { Secret() }).Should(gomega.Panic())
}
