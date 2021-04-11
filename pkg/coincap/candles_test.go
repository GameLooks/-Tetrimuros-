package coincap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCandles(t *testing.T) {
	teardown := setup()
	defer teardown()
	r.HandleFunc("/candles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("candles.json"))
	})

	req := CandlesReq