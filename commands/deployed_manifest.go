package commands

import "github.com/pivotal-cf/jhanda"

type DeployedManifest struct {
	logger           logger
	deployedProducts deployedProductsLister
	Options          struct {
		ProductName string `long:"product-name" short:"p" required:"true" description:"name of product"`
	}
}

func NewDeployedManifest(logger logger, deployedProducts deployedProductsLister) DeployedManifest {
	return DeployedManifest{
		logger:           logger,
		deployedProducts: deployedProducts,
	}
}

func (dm DeployedManifest) Execute(args []string) error {
	output, err := dm.deployedProducts.List()
	if err != nil {
		panic(err)
	}

	var guid string
	for _, product := range output {
		if product.Type == dm.Options.ProductName {
			guid = product.GUID
			break
		}
	}

	if guid == "" {
		panic("could not find product")
	}

	return nil
}

func (dm DeployedManifest) Usage() jhanda.Usage {
	return jhanda.Usage{}
}
