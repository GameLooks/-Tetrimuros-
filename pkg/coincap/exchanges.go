package coincap

import (
	"encoding/json"
	"net/http"
)

// Exchange contains information about a cryptocurrency exchange. This includes the exchanges
// relative rank, volume, and whether trading sockets are available
type Exchange str