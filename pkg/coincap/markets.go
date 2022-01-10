package coincap

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// MarketsRequest contains the paramters you can use to tailor a request for market data from the /markets endpoint
type MarketsRequest struct {
	ExchangeID  string `json:"exchangeId,omitempty"`  // search by unique exchange ID
	BaseSymbol  string `json:"baseSymbol,omitempty"`  // return all results with this base symbol
	BaseID      string `json:"baseId,omitempty"`      // return all results with this base id
	QuoteSymbol string `json:"quoteSymbol,omitempty"` // return all results with this quote symbol
	QuoteID     string `json:"quoteId,omitempty"`     // return all results with this quote ID
	AssetSymbol string `json:"assetSymbol,omitempty"` // r