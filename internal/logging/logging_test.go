package logging

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetLogger(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
		Configure("info", "text", "")
	})
}

func TestConfigure_UsesConfigLevelWhenCLILevelIsEmpty(t *testing.T) {
	resetLogger(t)

	Configure("warn", "text", "")

	assert.Equal(t, log.WarnLevel, log.GetLevel())
}

func TestConfigure_CLILevelOverridesConfig(t *testing.T) {
	resetLogger(t)

	Configure("warn", "text", "debug")

	assert.Equal(t, log.DebugLevel, log.GetLevel())
}

func TestConfigure_InvalidCLILevelUsesConfig(t *testing.T) {
	resetLogger(t)
	var output bytes.Buffer
	log.SetOutput(&output)

	Configure("warn", "text", "verbose")

	assert.Equal(t, log.WarnLevel, log.GetLevel())
	assert.Contains(t, output.String(), "invalid CLI log level")
}

func TestConfigure_Formatters(t *testing.T) {
	tests := []struct {
		name       string
		format     string
		assertBody func(t *testing.T, output string)
	}{
		{
			name:   "text",
			format: "text",
			assertBody: func(t *testing.T, output string) {
				assert.Contains(t, output, "formatter-test")
			},
		},
		{
			name:   "logfmt",
			format: "logfmt",
			assertBody: func(t *testing.T, output string) {
				assert.Contains(t, output, `msg=formatter-test`)
			},
		},
		{
			name:   "json",
			format: "json",
			assertBody: func(t *testing.T, output string) {
				var record map[string]any
				require.NoError(t, json.Unmarshal([]byte(strings.TrimSpace(output)), &record))
				assert.Equal(t, "formatter-test", record["msg"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetLogger(t)
			var output bytes.Buffer
			log.SetOutput(&output)
			Configure("info", tt.format, "")

			log.Info("formatter-test")

			tt.assertBody(t, output.String())
		})
	}
}
