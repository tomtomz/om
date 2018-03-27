package acceptance

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("config-template command", func() {
	const expectedOutput = `---
product-properties:
  .properties.with_default:
    value: some-default
  .properties.without_default:
    value: ""
  .some-instance-group.with_default:
    value: some-default
  .some-instance-group.without_default:
    value: ""
  .properties.with_named_manifest:
    value: enable
  .properties.with_named_manifest_without_default:
    value: ""
`

	It("outputs a configuration template based on the given product file", func() {
		command := exec.Command(pathToMain,
			"config-template",
			"--product",
			"fixtures/example.pivotal")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "10s").Should(gexec.Exit(0))

		Expect(string(session.Out.Contents())).To(MatchYAML(expectedOutput))
	})
})
