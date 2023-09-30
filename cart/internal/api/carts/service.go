package carts

import (
	"context"
	"route256/cart/internal/models"
	servicepb "route256/cart/pkg/api/carts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Deps struct {
	ItemAddService
	CheckoutService
	ItemDeleteService
	ListItemService
	ClearService
}

type Service struct {
	servicepb.UnimplementedCartsServer
	itemAddService    ItemAddService
	checkoutService   CheckoutService
	itemDeleteService ItemDeleteService
	listItemService   ListItemService
	clearService      ClearService
}

func NewServer(d Deps) *Service {
	return &Service{
		itemAddService:    d.ItemAddService,
		checkoutService:   d.CheckoutService,
		itemDeleteService: d.ItemDeleteService,
		listItemService:   d.ListItemService,
		clearService:      d.ClearService,
	}
}

type (
	ItemAddService interface {
		AddItem(ctx context.Context, userID models.UserID, sku models.SKU, count uint16) error
	}

	CartService interface {
		Clear(ctx context.Context, userID models.UserID) error
	}

	CheckoutService interface {
		Checkout(ctx context.Context, userID models.UserID) (models.OrderID, error)
	}

	ClearService interface {
		Clear(ctx context.Context, userID models.UserID) error
	}

	ItemDeleteService interface {
		DeleteItem(ctx context.Context, userID models.UserID, sku models.SKU) error
	}

	ListItemService interface {
		ListItem(ctx context.Context, userID models.UserID) (totalPrice uint32, items []models.CartItem, err error)
	}
)

func (s Service) ItemAdd(ctx context.Context, req *servicepb.ItemAddRequest) (*emptypb.Empty, error) {
	if err := s.itemAddService.AddItem(
		ctx,
		models.UserID(req.User),
		models.SKU(req.Sku),
		uint16(req.Count),
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (s Service) ItemDelete(ctx context.Context, req *servicepb.ItemDeleteRequest) (*emptypb.Empty, error) {
	if err := s.itemDeleteService.DeleteItem(
		ctx,
		models.UserID(req.User),
		models.SKU(req.Sku),
	); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Service) Clear(ctx context.Context, req *servicepb.ClearRequest) (*emptypb.Empty, error) {
	if err := s.clearService.Clear(
		ctx,
		models.UserID(req.User),
	); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Service) Checkout(ctx context.Context, req *servicepb.CheckoutRequest) (*servicepb.CheckoutResponse, error) {
	orderID, err := s.checkoutService.Checkout(
		ctx,
		models.UserID(req.User),
	)
	if err != nil {
		return nil, err
	}

	return &servicepb.CheckoutResponse{
		OrderId: int64(orderID),
	}, nil
}

func (s Service) List(ctx context.Context, req *servicepb.ListRequest) (*servicepb.ListResponse, error) {
	totalPrice, list, err := s.listItemService.ListItem(
		ctx,
		models.UserID(req.User),
	)
	if err != nil {
		return nil, err
	}

	items := make([]*servicepb.ListItem, 0, len(list))
	for _, item := range list {
		items = append(items, &servicepb.ListItem{
			Sku:   uint32(item.SKU),
			Count: uint32(item.Count),
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &servicepb.ListResponse{
		TotalPrice: totalPrice,
		Items:      items,
	}, nil
}
