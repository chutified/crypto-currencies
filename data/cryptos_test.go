package data

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestService(t *testing.T) {

	// >>> New()
	s := New()

	// >>> Update()
	tests0 := []struct {
		name string
		url  string
		err  string
	}{
		{
			name: "ok",
			url:  "https://coinmarketcap.com/all/views/all",
			err:  "",
		},
		{
			name: "invalid content",
			url:  "https://coinmarketcap.com/rankings/exchanges",
			err:  "parsing records",
		},
		{
			name: "invalid url",
			url:  "invalid",
			err:  "fetching records",
		},
	}

	for _, test := range tests0 {
		t.Run(test.name, func(t1 *testing.T) {

			err := s.Update(test.url)
			if err != nil {

				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
			}
		})
	}

	// >>> GetCurrency()
	tests1 := []struct {
		name     string
		currency string
		err      string
	}{
		{
			name:     "full name",
			currency: "BITCOIN",
			err:      "",
		},
		{
			name:     "symbol",
			currency: "BTC",
			err:      "",
		},
		{
			name:     "not found",
			currency: "invalid",
			err:      "currency .* not found",
		},
	}

	for _, test := range tests1 {
		t.Run(test.name, func(t1 *testing.T) {

			m, err := s.GetCurrency(test.currency)
			if err != nil {

				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
				assert.NotEqual(t1, m, nil)
			}
		})
	}

	// >>> MonitorData()
	tests2 := []struct {
		name string
		url  string
		err  bool
	}{
		{
			name: "ok",
			url:  "https://coinmarketcap.com/all/views/all", err: false,
		},
		{
			name: "invalid ",
			url:  "invalid",
			err:  true,
		},
	}

	for _, test := range tests2 {
		t.Run(test.name, func(t1 *testing.T) {

			_, errs := s.MonitorData(500*time.Millisecond, test.url)

			var hasErr bool
			go func() {
				_ = <-errs
				hasErr = true
			}()

			time.Sleep(2 * time.Second)
			assert.Equal(t1, hasErr, test.err)
		})
	}
}
