package loms

import (
	"fmt"
	"net/url"
)

type Client struct {
	name string
	path string
}

func New(name string, basePath string) (*Client, error) {
	const stocksPath = "stocks"
	path, err := url.JoinPath(basePath, stocksPath)
	if err != nil {
		return nil, fmt.Errorf("%s: incorrect base path: %w", name, err)
	}
	return &Client{
		name: name,
		path: path,
	}, nil
}

func (c Client) GetStocks(sku uint32) (uint64, error) {
	return 0, nil
}
