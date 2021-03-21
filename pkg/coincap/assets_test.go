
package coincap

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestAssets(t *testing.T) {
	teardown := setup()
	defer teardown()
	r.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("assets.json"))
	})

	params := &AssetsRequest{
		Search: "BTC",
		Limit:  4,
		Offset: 1,
	}
	assets, _, err := client.Assets(params)
	if err != nil {
		t.Fatal(err)
	}
	if len(assets) < 1 {
		t.Fatalf("No assets were returned")
	}

	asset := assets[0]
	if asset.ID != "bitcoin-private" {
		t.Errorf("Expected assetID to be bitcoin-private but was: %s", asset.ID)
	}

}