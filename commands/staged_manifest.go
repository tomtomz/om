package commands

import "fmt"

//go:generate counterfeiter -o ./fakes/staged_manifest_service.go --fake-name StagedManifestService . stagedManifestService
type stagedManifestService interface {
	Show(productID string) (string, error)
}

type StagedManifest struct {
	stagedManifestService stagedManifestService
	stagedProductsFinder  stagedProductsFinder
	Options               struct {
		ProductName string `long:"product-name" short:"p" required:"true" description:"name of product"`
	}
}

func NewStagedManifest(stagedManifestService stagedManifestService, stagedProductsFinder stagedProductsFinder) StagedManifest {
	return StagedManifest{
		stagedProductsFinder:  stagedProductsFinder,
		stagedManifestService: stagedManifestService,
	}
}

func (s StagedManifest) Execute(args []string) error {

	findOutput, err := s.stagedProductsFinder.Find(s.Options.ProductName)
	if err != nil {
		panic(err)
		// return fmt.Errorf("failed to find staged product %q: %s", e.Options.ProductName, err)
	}

	stagedManifestOutput, err := s.stagedManifestService.Show(findOutput.Product.GUID)
	if err != nil {
		panic(err)
	}

	fmt.Println(stagedManifestOutput)
	return nil
}
