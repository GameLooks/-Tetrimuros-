package coincap

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// AssetsRequest contains the paramaters for modifying a query to
// the "/assets" endp