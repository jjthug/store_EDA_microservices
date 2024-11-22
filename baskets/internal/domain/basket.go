package domain

import "github.com/stackus/errors"

const basketAggregate = "basket.Basket"

var (
	ErrBasketHasNoItems = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
)

type Basket struct {
	es.Aggregate
	CustomerID string
	PaymentID string
	Items map[string]Item
	Status BasketStatus
}
