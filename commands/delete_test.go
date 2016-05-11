package commands_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
	"github.com/pivotal-cf/cm-cli/config"
)

var _ = Describe("Get", func() {
	It("displays help", func() {
		session := runCommand("delete", "-h")

		Eventually(session).Should(Exit(1))
		Expect(session.Err).To(Say("delete"))
		Expect(session.Err).To(Say("--name"))
	})

	It("deletes a secret", func() {
		server.AppendHandlers(
			CombineHandlers(
				VerifyRequest("DELETE", "/api/v1/secret/my-secret"),
				RespondWith(http.StatusOK, ""),
			),
		)

		session := runCommand("delete", "-n", "my-secret")

		Eventually(session).Should(Exit(0))
		Eventually(session.Out).Should(Say("Secret successfully deleted"))
	})

	It("prints an error when API URL is not set", func() {
		cfg := config.ReadConfig()
		cfg.ApiURL = ""
		config.WriteConfig(cfg)

		session := runCommand("delete", "-n", "my-secret")

		Eventually(session).Should(Exit(1))
		Eventually(session.Err).Should(Say("API location is not set"))
	})
})