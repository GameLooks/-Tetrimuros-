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
	if len(e