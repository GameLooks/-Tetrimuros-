
package coincap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var (
	r      *mux.Router
	server *httptest.Server
	client *Client
)

func setup() func() {
	r = mux.NewRouter()
	server = httptest.NewServer(r)

	client = NewClient(nil)
	client.SetBaseURL(server.URL)

	return func() {
		server.Close()
	}
}

// load test data
func fixture(path string) string {
	f, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(f)
}

func TestRates(t *testing.T) {
	teardown := setup()
	defer teardown()

	r.HandleFunc("/rates", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("rates.json"))
	})

	rates, _, err := client.Rates()
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) < 0 {
		t.Fatalf("No rates returned")
	}
	got := rates[0]
	expected := Rate{
		ID:             "romanian-leu",
		Symbol:         "RON",
		CurrencySymbol: "lei",
		Type:           "fiat",
		RateUSD:        "0.2505529076289101",
	}
	if *got != expected {
		t.Errorf("Expected %s, Got %s", expected, got)
	}
}

func TestRatesBadURL(t *testing.T) {
	teardown := setup()
	defer teardown()

	client.SetBaseURL("bad.fail:")
	_, _, err := client.Rates()
	if err == nil {
		t.Errorf("Expected client to fail on Rates() because of bad base path")
	}
	_, _, err = client.RateByID("bitcoin")
	if err == nil {
		t.Errorf("Expected client to fail on RateByID() because of bad base path")
	}

}

func TestRateByID(t *testing.T) {
	teardown := setup()