package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VMExtensionsService struct {
	client httpClient
}

type CreateVMExtension struct {
	Name            string          `json:"name"`
	CloudProperties json.RawMessage `json:"cloud_properties"`
}

func NewVMExtensionsService(client httpClient) VMExtensionsService {
	return VMExtensionsService{
		client: client,
	}
}

type VMExtensionInput struct {
	Name            string `json:"name"`
	CloudProperties string `json:"cloud_properties"`
}

func (v VMExtensionsService) Create(input CreateVMExtension) error {
	jsonData, err := json.Marshal(&input)

	println(string(jsonData))
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/api/v0/staged/vm_extensions", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := v.client.Do(req)
	if err != nil {
		return err
	}

	if err = ValidateStatusOK(resp); err != nil {
		return err
	}

	return nil
}
