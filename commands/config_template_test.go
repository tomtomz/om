package commands_test

import (
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConfigTemplate", func() {
	var (
		logger *fakes.Logger
	)

	BeforeEach(func() {
		logger = &fakes.Logger{}
	})

	Describe("Usage", func() {
		It("returns usage information for the command", func() {
			command := commands.NewConfigTemplate(nil)
			Expect(command.Usage()).To(Equal(jhanda.Usage{
				Description:      "This command generates a configuration template that can be passed in to om configure-product",
				ShortDescription: "generates a config template for the product",
				Flags:            command.Options,
			}))
		})
	})

	Describe("failure cases", func() {
		Context("when an unknown flag is provided", func() {
			It("returns an error", func() {
				command := commands.NewConfigTemplate(nil)
				err := command.Execute([]string{"--badflag"})
				Expect(err).To(MatchError("could not parse config-template flags: flag provided but not defined: -badflag"))
			})
		})

		Context("when the product flag is not provided", func() {
			It("returns an error", func() {
				command := commands.NewConfigTemplate(nil)
				err := command.Execute([]string{})
				Expect(err).To(MatchError("could not parse config-template flags: missing required flag \"--product\""))
			})
		})

		Context("when the product file cannot be found", func() {
			It("returns an error", func() {
				command := commands.NewConfigTemplate(nil)
				err := command.Execute([]string{"--product", "/missing/product/file"})
				Expect(err).To(MatchError(HavePrefix("could not open product file: /missing/product/file: ")))
			})
		})

	})
})
