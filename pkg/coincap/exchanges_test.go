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
