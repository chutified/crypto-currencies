package data

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestParseRecords(t *testing.T) {

	tests := []struct {
		name string
		rec  record
		err  string
	}{
		{
			name: "ok",
			rec: record{
				name:              "Bitcoin",
				symbol:            "BTC",
				marketCap:         "$202,198,068,939",
				price:             "$10,962.33",
				circulatingSupply: "18,444,800 BTC",
				volume:            "$36,608,208,388",
				changeHour:        "1.78%",
				changeDay:         "7.02%",
				changeWeek:        "7.02%",
			},
			err: "",
		},
		{
			name: "invalid marketCap",
			rec: record{
				name:              "Bitcoin",
				symbol:            "BTC",
				marketCap:         "$invalid",
				price:             "$10,962.33",
				circulatingSupply: "18,444,800 BTC*",
				volume:            "$36,608,208,388",
				changeHour:        "1.78%",
				changeDay:         "7.02%",
				changeWeek:        "7.02%",
			},
			err: "parse marketCap",
		},
		{
			name: "invalid price",
			rec: record{
				name:              "Bitcoin",
				symbol:            "BTC",
				marketCap:         "$202,198,068,939",
				price:             "$invalid",
				circulatingSupply: "18,444,800 BTC",
				volume:            "$36,608,208,388",
				changeHour:        "1.78%",
				changeDay:         "7.02%",
				changeWeek:        "7.02%",
			},
			err: "parse price",
		},
		{
			name: "invalid circulating supply",
			rec: record{
				name:              "Bitcoin",
				symbol:            "BTC",
				marketCap:         "$202,198,068,939",
				price:             "$10,962.33",
				circulatingSupply: "invalid",
				volume:            "$36,608,208,388",
				changeHour:        "1.78%",
				changeDay:         "7.02%",
				changeWeek:        "7.02%",
			},
			err: "parse circulatingSupply",
		},
		{
			name: "invalid volume",
			rec: record{
				name:              "Bitcoin",
				symbol:            "BTC",
				marketCap:         "$202,198,068,939",
				price:             "$10,962.33",
				circulatingSupply: "18,444,800 BTC",
				volume:            "$invalid",
				changeHour:        "1.78%",
				changeDay:         "7.02%",
				changeWeek:        "7.02%",
			},
			err: "parse volume",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			m, err := parseRecords([]record{test.rec})
			if err != nil {

				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
				assert.NotEqual(t1, m, nil)
			}

		})
	}
}
