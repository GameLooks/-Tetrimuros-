package coincap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCandles(t *testing.T) {
	teardown := setup()
	defer te