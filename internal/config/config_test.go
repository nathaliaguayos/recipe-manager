package config_test

import (
	"github.com/recipe-manager/internal/config"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	getEnv := func() map[string]string {
		envs := make(map[string]string)
		for _, e := range os.Environ() {
			kv := strings.SplitN(e, "=", 2)
			envs[kv[0]] = kv[1]
		}
		return envs
	}
	initialEnv := getEnv()
	resetEnv := func(t *testing.T) {
		for k := range getEnv() {
			_, present := initialEnv[k]
			if !present {
				require.NoError(t, os.Unsetenv(k))
			}
		}
		for k, v := range initialEnv {
			require.NoError(t, os.Setenv(k, v))
		}
	}
	tests := []struct {
		it     string
		envs   func(t *testing.T) map[string]string
		assert func(t *testing.T, c *config.Config, err error)
	}{
		{
			it: "valid env vars should return hydrated config",
			envs: func(_ *testing.T) map[string]string {
				return map[string]string{
					config.EnvPrefix + "_SERVICE_NAME": "recipe-manager",
					config.EnvPrefix + "_PORT":         "8080",
					config.EnvPrefix + "_LOG_LEVEL":    "info",
				}
			},
			assert: func(t *testing.T, c *config.Config, err error) {
				require.NoError(t, err, "no error should be returned")
				require.NotNil(t, c, "config should not be nil on success")
				require.Equal(t, &config.Config{
					ServiceName: "recipe-manager",
					Port:        8080,
					Host:        "0.0.0.0",
					LogLevel:    "info",
				}, c, "invalid config returned")
			},
		},
	}
	for _, tt := range tests {
		for k, v := range tt.envs(t) {
			require.NoError(t, os.Setenv(k, v), "cannot set env var")
		}
		t.Run(tt.it, func(t *testing.T) {
			cfg, err := config.Get()
			tt.assert(t, cfg, err)
		})
		resetEnv(t)
	}
}
