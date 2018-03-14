package commands_test

import (
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StagedManifest", func() {

	var (
		stagedProductsFinder  *fakes.StagedProductsFinder
		stagedManifestService *fakes.StagedManifestService
		command               commands.StagedManifest
	)

	BeforeEach(func() {
		stagedProductsFinder = &fakes.StagedProductsFinder{}
		stagedManifestService = &fakes.StagedManifestService{}
		command = commands.NewStagedManifest(stagedManifestService, stagedProductsFinder)
	})

	Describe("Execute", func() {
		It("returns the manifest", func() {
			stagedManifestService.ShowReturns("manifest output", nil)

			stagedProductsFinder.FindReturns(api.StagedProductsFindOutput{
				Product: api.StagedProduct{
					Type: "some-product-name",
					GUID: "some-product-id",
				},
			}, nil)

			err := command.Execute([]string{"--product-name", "some-product-name"})
			Expect(err).NotTo(HaveOccurred())

			Expect(stagedProductsFinder.FindCallCount()).To(Equal(1))
			Expect(stagedProductsFinder.FindArgsForCall(0)).To(Equal("some-product-name"))

			Expect(stagedManifestService.ShowCallCount()).To(Equal(1))
			Expect(stagedManifestService.ShowArgsForCall(0)).To(Equal("some-product-id"))

			//test output here
		})
	})
})
