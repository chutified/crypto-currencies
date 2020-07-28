package config

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	tests := []struct {
		name string
		file string
		err  string
	}{
		{
			name: "ok",
			file: "config/tests/config_0.yaml",
			err:  "",
		},
		{
			name: "invalid",
			file: "config/tests/config_1.yaml",
			err:  "could not decode config file content",
		},
		{
			name: "non-existing",
			file: "config/tests/config_2.yaml",
			err:  "could not read config file",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			cfg, err := GetConfig(test.file)
			if err != nil {

				exp := fmt.Sprintf("%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.err)
				assert.Equal(t1, cfg.Host, "localhost")
				assert.Equal(t1, cfg.Port, 10503)
			}
		})
	}
}
