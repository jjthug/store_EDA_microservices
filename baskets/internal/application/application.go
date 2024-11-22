package application

import (
	"EDA_GO/baskets/internal/domain"
	"EDA_GO/internal/ddd"
	"context"
)

type (
	StartBasket struct {
		ID         string
		CustomerID string
	}

	CancelBasket struct {
		ID string
	}

	CheckoutBasket struct {
		ID        string
		PaymentID string
	}

	AddItem struct {
		ID        string
		ProductID string
		Quantity  int
	}

	RemoveItem struct {
		ID        string
		ProductID string
		Quantity  int
	}

	GetBasket struct {
		ID string
	}

	App interface {
		StartBasket(ctx context.Context, start StartBasket) error
		CancelBasket(ctx context.Context, cancel CancelBasket) error
		CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
		AddItem(ctx context.Context, addItem AddItem) error
		RemoveItem(ctx context.Context, removeItem RemoveItem) error
		GetBasket(ctx context.Context, getBasket GetBasket) error
	}

	Application struct {
		baskets   domain.BasketRepository
		stores    domain.StoreRepository
		products  domain.ProductRepository
		publisher ddd.EventPublisher[ddd.Event]
	}
)

var _ App = (*Application)(nil)
