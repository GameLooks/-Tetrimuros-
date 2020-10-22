package coincap

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// AssetsRequest contains the paramaters for modifying a query to
// the "/assets" endpoint. Search can be a symbol (BTC) or an asset id (bitcoin)
type AssetsRequest struct {
	Search string `json:"search,omitempty"` // search by asset id (bitcoin) or symbol (BTC)
	Limit  int    `json:"limit,omitempty"`  // limit number of returned results (Max: 2000)
	Offset int    `json:"offset,omitempty"` // skip the first N entries of the result set
}

// Asset contains various information about a given CoinCap asset such as Bitcoin
type Asset struct {
	ID                string `json:"id"`                // unique identifier for asset
	Rank              string `json:"rank"`              // rank in terms of the asset's market cap
	Symbol            string `json:"symbol"`            // common symbol to identify the asset
	Name              string `json:"name"`              // proper name for asset
	Supply            string `json:"supply"`            // available supply for trading
	MaxSupply         string `json:"maxSupply"`         // total quantity of asset issued
	MarketCapUsd      string `json:"marketCapUsd"`      // supply x price
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`     // quantity of trading volume in USD over last 24 hours
	PriceUsd          string `json:"priceUsd"`          // volume weighted price of the asset in USD
	ChangePercent24Hr str