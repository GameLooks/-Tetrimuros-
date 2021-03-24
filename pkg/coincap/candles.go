
package coincap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CandlesRequest contains the parameters you can use to customize a request for candle data from the /candles endpoint
type CandlesRequest struct {
	ExchangeID string   `json:"exchangeId"`       // search by unique exchange ID
	BaseID     string   `json:"baseId"`           // return all results with this base id
	QuoteID    string   `json:"quoteId"`          // return all results with this quote ID
	Interval   Interval `json:"interval"`         // candle interval
	Start      int      `json:"start,omitempty"`  // start time in unix milliseconds TODO: I should probably use time.Time or Timestamp here
	End        int      `json:"end,omitempty"`    // end time in unix milliseconds TODO: same as above
	Limit      int      `json:"limit,omitempty"`  // limit number of returned results (Max: 2000)
	Offset     int      `json:"offset,omitempty"` // skip the first N entries of the result set
}

// Candle represets historic market performance for an asset over a given timeframe