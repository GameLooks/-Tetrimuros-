package coincap

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// MarketsRequest contains the paramters you can use to tailor a request for market data from the /markets endpoint
type MarketsRequest struct {
	ExchangeID  string `json:"exchangeId,omitempty"`  // search by unique exchange ID
	BaseSymbol  string `j