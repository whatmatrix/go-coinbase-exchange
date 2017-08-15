package coinbase

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Order struct {
	Type      string          `json:"type"`
	Size      decimal.Decimal `json:"size,string"`
	Side      string          `json:"side"`
	ProductId string          `json:"product_id"`
	ClientOID string          `json:"client_oid,omitempty"`
	Stp       string          `json:"stp,omitempty"`
	// Limit Order
	Price       decimal.Decimal `json:"price,string,omitempty"`
	TimeInForce string          `json:"time_in_force,omitempty"`
	PostOnly    bool            `json:"post_only,omitempty"`
	CancelAfter string          `json:"cancel_after,omitempty"`
	// Market Order
	Funds decimal.Decimal `json:"funds,string,omitempty"`
	// Response Fields
	Id            string          `json:"id"`
	Status        string          `json:"status,omitempty"`
	Settled       bool            `json:"settled,omitempty"`
	DoneReason    string          `json:"done_reason,omitempty"`
	CreatedAt     Time            `json:"created_at,string,omitempty"`
	FillFee       decimal.Decimal `json:"fill_fee,omitempty"`
	FilledSize    decimal.Decimal `json:"filled_size,omitempty"`
	ExecutedValue decimal.Decimal `json:"executed_value,omitempty"`
}

type ListOrdersParams struct {
	Status     string
	Pagination PaginationParams
}

func (c *Client) CreateOrder(newOrder *Order) (Order, error) {
	var savedOrder Order

	if len(newOrder.Type) == 0 {
		newOrder.Type = "limit"
	}

	url := fmt.Sprintf("/orders")
	_, err := c.Request("POST", url, newOrder, &savedOrder)
	return savedOrder, err
}

func (c *Client) CancelOrder(id string) error {
	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("DELETE", url, nil, nil)
	return err
}

func (c *Client) GetOrder(id string) (Order, error) {
	var savedOrder Order

	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.Request("GET", url, nil, &savedOrder)
	return savedOrder, err
}

func (c *Client) ListOrders(p ...ListOrdersParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
		if p[0].Status != "" {
			paginationParams.AddExtraParam("status", p[0].Status)
		}
	}

	return NewCursor(c, "GET", fmt.Sprintf("/orders"),
		&paginationParams)
}
