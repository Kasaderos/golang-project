package item

import "errors"

type StocksProvider interface {
	GetStocks(sku uint32) (uint64, error)
}

type ProductProvider interface {
	GetProductInfo(sku uint32) (string, uint32, error)
}

type AddService struct {
	stocksProvider  StocksProvider
	productProvider ProductProvider
}

func NewAddService(stocksProvider StocksProvider, productProvider ProductProvider) *AddService {
	return &AddService{
		stocksProvider:  stocksProvider,
		productProvider: productProvider,
	}
}

var ErrInsufficientStocks = errors.New("Insufficient stocks")

func (s AddService) Add(user int64, sku uint32, count uint16) error {
	if _, _, err := s.productProvider.GetProductInfo(sku); err != nil {
		return err
	}
	stocksCount, err := s.stocksProvider.GetStocks(sku)
	if err != nil {
		return err
	}
	if uint64(count) > stocksCount {
		return ErrInsufficientStocks
	}
	return nil
}
