package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"route256/cart/internal/models"
	"route256/cart/internal/usecase/cart"
)

var (
	ErrOrderNotCreated = errors.New("order not created")
)

const (
	GetProductInfoPath = "/product/info"
)

type productService struct {
	name       string
	baseURL    string
	httpClient *http.Client
}

var _ cart.ProductService = (*productService)(nil)

func NewProductService(baseURL string) *productService {
	return &productService{
		name:       "product service",
		httpClient: &http.Client{},
		baseURL:    baseURL,
	}
}

func (c *productService) GetProductInfo(ctx context.Context, sku models.SKU) (name string, price uint32, err error) {
	reqBody := GetProductRequest{
		Token: "testtoken", // TODO: get from request and pass into ctx
		SKU:   uint32(sku),
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to encode request %w", c.name, err)
	}

	reqURL, err := url.JoinPath(c.baseURL, GetProductInfoPath)
	if err != nil {
		return "", 0, err
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(data))
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to create HTTP request: %w", c.name, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to execute HTTP request: %w", c.name, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		var response GetProductErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return "", 0, fmt.Errorf("%s: failed to decode error response: %w", c.name, err)
		}
		return "", 0, fmt.Errorf("%s: HTTP request responded with: %d , message: %s", c.name, resp.StatusCode, response.Message)
	}

	var response GetProductResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to decode error response: %w", c.name, err)
	}

	return response.Name, response.Price, nil
}
