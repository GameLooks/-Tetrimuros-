
package coincap

import (
	"encoding/json"
	"net/http"
)

// Rate contains the exchange rate of a given asset in terms of USD as well as
// common identifiers for the asset in question and whether or not it is a fiat currency
type Rate struct {
	ID             string `json:"id"`             // unique identifier for asset or fiat
	Symbol         string `json:"symbol"`         // most common symbol used to identify asset or fiat
	CurrencySymbol string `json:"currencySymbol"` // currency symbol if available
	RateUSD        string `json:"rateUsd"`        // rate conversion to USD
	Type           string `json:"type"`           // type of currency - fiat or crypto
}

// Rates returns currency rates standardized in USD.
// Fiat rates are sourced from OpenExchangeRates.org
func (c *Client) Rates() ([]*Rate, *Timestamp, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/rates", nil)
	if err != nil {
		return nil, nil, err
	}

	// make the request