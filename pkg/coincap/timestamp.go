package coincap

import (
	"strconv"
	"time"
)

// Timestamp is wrapper around time.Time with custom marshaling behaviour
// specific to the format returned by the CoinCap API
type Timesta