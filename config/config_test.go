package config

import "testing"

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

		})
	}
}
