package commands_test

import (
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeployedManifest", func() {
	var (
		command                commands.DeployedManifest
		logger                 *fakes.Logger
		deployedProductsLister *fakes.DeployedProductsLister
	)

	BeforeEach(func() {
		logger = &fakes.Logger{}
		deployedProductsLister = &fakes.DeployedProductsLister{}
		deployedProductsLister.ListReturns([]api.DeployedProductOutput{
			{Type: "other-product", GUID: "other-product-guid"},
			{Type: "some-product", GUID: "some-product-guid"},
		}, nil)

		command = commands.NewDeployedManifest(logger, deployedProductsLister)
	})

	It("prints the manifest of the deployed product", func() {
		err := command.Execute([]string{
			"--product-name", "some-product",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(deployedProductsLister.ListCallCount()).To(Equal(1))

		Expect(deployedProductsLister.ManifestCallCount()).To(Equal(1))
		Expect(deployedProductsLister.ManifestArgsForCall(0)).To(Equal("some-product-guid"))

		Expect(logger.PrintCallCount()).To(Equal(1))
		Expect(logger.PrintArgsForCall(0)[0]).To(MatchYAML(`---
name: some-product
key: value
`))
	})
})
