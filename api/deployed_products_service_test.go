package api_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/api/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeployedProductsService", func() {
	Describe("Manifest", func() {
		var (
			client  *fakes.HttpClient
			service api.DeployedProductsService
		)

		BeforeEach(func() {
			client = &fakes.HttpClient{}
			service = api.NewDeployedProductsService(client)

			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				switch req.URL.Path {
				case "/api/v0/deployed/products/some-product-guid/manifest":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`{
							"key-1": {
								"key-2": "value-1"
							},
							"key-3": "value-2",
							"key-4": 2147483648
						}`)),
					}
				}
				return resp, nil
			}
		})

		It("returns a manifest of a product", func() {
			manifest, err := service.Manifest("some-product-guid")
			Expect(err).NotTo(HaveOccurred())
			Expect(manifest).To(MatchYAML(`---
key-1:
  key-2: value-1
key-3: value-2
key-4: 2147483648
`))
		})

		Context("failure cases", func() {
			Context("when the request object is invalid", func() {
				It("returns an error", func() {
					_, err := service.Manifest("invalid-guid-%%%")
					Expect(err).To(MatchError(ContainSubstring("invalid URL escape")))
				})
			})

			Context("when the client request fails", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{}, errors.New("nope"))

					_, err := service.Manifest("some-product-guid")
					Expect(err).To(MatchError("could not make api request to staged products manifest endpoint: nope"))
				})
			})

			Context("when the server returns a non-200 status code", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusTeapot,
						Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					}, nil)

					_, err := service.Manifest("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})

			Context("when the returned JSON is invalid", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString("%%%")),
					}, nil)

					_, err := service.Manifest("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("could not parse json")))
				})
			})
		})
	})

	Describe("List", func() {
		var (
			client *fakes.HttpClient
		)

		BeforeEach(func() {
			client = &fakes.HttpClient{}
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				resp = &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
				}
				switch req.URL.Path {
				case "/api/v0/deployed/products":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`[{
							"guid":"some-product-guid",
							"type":"some-type"
						},
						{
							"guid":"some-other-product-guid",
							"type":"some-other-type"
						}]`)),
					}
				}
				return resp, nil
			}
		})

		It("retrieves a list of deployed products from the Ops Manager", func() {
			service := api.NewDeployedProductsService(client)

			output, err := service.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(Equal([]api.DeployedProductOutput{
				{
					GUID: "some-product-guid",
					Type: "some-type",
				},
				{
					GUID: "some-other-product-guid",
					Type: "some-other-type",
				},
			},
			))

			Expect(client.DoCallCount()).To(Equal(1))

			By("checking for deployed products")
			avReq := client.DoArgsForCall(0)
			Expect(avReq.URL.Path).To(Equal("/api/v0/deployed/products"))
		})

		Context("failure cases", func() {
			Context("when the request fails", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{}, errors.New("nope"))
				})

				It("returns an error", func() {
					service := api.NewDeployedProductsService(client)

					_, err := service.List()
					Expect(err).To(MatchError("could not make api request to deployed products endpoint: nope"))
				})
			})

			Context("when the server returns a non-200 status code", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusTeapot,
						Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					}, nil)
				})

				It("returns an error", func() {
					service := api.NewDeployedProductsService(client)

					_, err := service.List()
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})

			Context("when the server returns invalid JSON", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString("%%")),
					}, nil)
				})

				It("returns an error", func() {
					service := api.NewDeployedProductsService(client)

					_, err := service.List()
					Expect(err).To(MatchError(ContainSubstring("could not unmarshal deployed products response:")))
				})
			})
		})
	})
})
