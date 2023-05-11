package coincap

import (
	"strconv"
	"time"
)

// Timestamp is wrapper around time.Time with custom marshaling behaviour
// specific to the format returned by the CoinCap API
type Timestamp struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
// Custom unmarshaller to handle that the timestamp is not in a standard format
func (t *Timesta