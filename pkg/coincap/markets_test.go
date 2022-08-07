
package coincap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMarkets(t *testing.T) {

	teardown := setup()
	defer teardown()

	r.HandleFunc("/markets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)