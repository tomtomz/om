package commands

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/pivotal-cf/jhanda"
	yaml "gopkg.in/yaml.v2"
)

type ConfigTemplate struct {
	logger  logger
	Options struct {
		Product string `long:"product"  short:"p"  required:"true" description:"path to product to generate config template for"`
	}
}

type propertyBlueprint struct {
	Name         string      `yaml:"name"`
	Optional     bool        `yaml:"optional"`
	Configurable bool        `yaml:"configurable"`
	Type         string      `yaml:"type"`
	Default      interface{} `yaml:"default"`
}

type instanceGroup struct {
	Name       string              `yaml:"name"`
	Properties []propertyBlueprint `yaml:"property_blueprints"`
}

type metadata struct {
	Properties     []propertyBlueprint `yaml:"property_blueprints"`
	InstanceGroups []instanceGroup     `yaml:"job_types"`
}

func NewConfigTemplate(logger logger) ConfigTemplate {
	return ConfigTemplate{
		logger: logger,
	}
}

func (ct ConfigTemplate) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This command generates a configuration template that can be passed in to om configure-product",
		ShortDescription: "generates a config template for the product",
		Flags:            ct.Options,
	}
}

func (ct ConfigTemplate) Execute(args []string) error {
	if _, err := jhanda.Parse(&ct.Options, args); err != nil {
		return fmt.Errorf("could not parse config-template flags: %s", err)
	}

	zipReader, err := zip.OpenReader(ct.Options.Product)
	if err != nil {
		return fmt.Errorf("could not open product file: %s: %s", ct.Options.Product, err)
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {

		matched, err := path.Match("metadata/*.yml", file.Name)

		if err != nil {
			return err
		}

		if !matched {
			continue
		}

		metadataFile, err := file.Open()
		if err != nil {
			return err
		}

		contents, err := ioutil.ReadAll(metadataFile)
		if err != nil {
			return err
		}

		input := metadata{}
		err = yaml.Unmarshal(contents, &input)
		if err != nil {
			panic(err)
		}

		productProperties := map[string]interface{}{}

		ct.fillPropertiesTemplate(productProperties, ".properties", input.Properties)

		for _, ig := range input.InstanceGroups {
			ct.fillPropertiesTemplate(productProperties, fmt.Sprintf(".%s", ig.Name), ig.Properties)
		}

		configTemplate := map[string]interface{}{
			"product-properties": productProperties,
		}

		template, err := yaml.Marshal(configTemplate)
		if err != nil {
			panic(err)
		}

		ct.logger.Printf("---\n%s\n", string(template))
	}

	return nil
}

func (ct ConfigTemplate) fillPropertiesTemplate(template map[string]interface{}, prefix string, properties []propertyBlueprint) {
	for _, property := range properties {
		if property.Configurable {
			var value interface{}
			value = ""
			if property.Default != nil {
				value = property.Default
			}
			template[fmt.Sprintf("%s.%s", prefix, property.Name)] = map[string]interface{}{
				"value": value,
			}
		}
	}
}
