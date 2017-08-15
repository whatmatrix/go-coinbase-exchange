package coinbase

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Currency struct {
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	MinSize decimal.Decimal `json:"min_size,string"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}
