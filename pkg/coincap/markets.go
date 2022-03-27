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
	AssetSymbol string `json:"assetSymbol,omitempty"` // return all results with this asset symbol
	AssetID     string `json:"assetID,omitempty"`     // return all results with this asset ID
	Limit       int    `json:"limit,omitempty"`       // limit number of returned results (Max: 2000)
	Offset      int    `json:"offset,omitempty"`      // skip the first N entries of the result set
}

// Market contains the market data response from the api
type Market struct {
	ExchangeID            string    `json:"exchangeId"`            // unique identifier for exchange
	Rank                  string    `json:"rank"`                  // rank in terms of volume transacted in this market
	BaseSymbol            string    `json:"baseSymbol"`            // most common symbol used to identify this asset
	BaseID                string    `json:"baseId"`                // unique identifier for this asset. base is the asset purchased
	QuoteSymbol           string    `json:"quoteSymbol"`           // most common symbol used to identify this asset
	QuoteID               string    `json:"quoteId"`               // unique identifier for thisasset. quote is the asset used to purchase base
	PriceQuote            string    `json:"priceQuote"`            // amount of quote asset traded for 1 unit of base asset
	PriceUsd              stri