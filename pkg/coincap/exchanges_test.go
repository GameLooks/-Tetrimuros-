package coincap

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestExchanges(t *testing.T) {

	teardown := setup()
	defer teardown()

	r.HandleFunc("/exchanges", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("exchange.json"))
	})

	exchanges, _, err := client.Exchanges()
	if err != nil {
		t.Fatal(err)
	}
	if len(exchanges) < 0 {
		t.Fatalf("No Exchanges returned")
	}
	got := exchanges[0]
	ts := Timestamp{
		Time: time.Unix(0, 1536336916333*1e6),
	}

	expected := Exchange{
		ID:                 "binance",
		Name:               "Binance",
		Rank:               