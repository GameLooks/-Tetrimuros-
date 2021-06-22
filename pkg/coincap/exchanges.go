package coincap

import (
	"encoding/json"
	"net/http"
)

// Exchange contains information about a cryptocurrency exchange. This includes the exchanges
// relative rank, volume, and whether trading sockets are available
type Exchange struct {
	ID                 string    `json:"id"`                 // unique identifier for exchange
	Name               string    `json:"name"`               // proper name of exchange
	Rank            