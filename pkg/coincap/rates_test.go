
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