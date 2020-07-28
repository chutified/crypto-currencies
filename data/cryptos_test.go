package data

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestService(t *testing.T) {

	// >>> New()
	s := New()

	// >>> Update()
	tests := []struct {
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

	for _, test := range tests {
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
}
