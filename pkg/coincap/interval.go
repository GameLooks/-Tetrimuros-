package coincap

// Interval represents point-in-time intervals for retrieving historical market data
type Interval string

// Valid Intervals for historical market data
// Used when requesting Asset History and Candles
const (
	// AssetHistory Intervals
	Minute         Interval = "m1"
	FifteenMinutes Interval = "m15"
	Hour           Interval = "h1"
	Day            Interval = "d1"

	// Candle Intervals (includes all the above intervals)
	FiveMinutes   Interval = "m5"
	Th