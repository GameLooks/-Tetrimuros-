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
	Rank              string `json:"rank"`  