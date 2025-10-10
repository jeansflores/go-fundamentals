package currencyapi_test

import (
	"testing"

	"example.com/crypto/currencyapi"
)

func TestGetRate(t *testing.T) {
	_, err := currencyapi.GetRate("")
	if err == nil {
		t.Errorf("Empty currencies should return error")
	}
}
