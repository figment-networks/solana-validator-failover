package config

import (
	"fmt"

	"github.com/charmbracelet/log"
)

const (
	// DefaultLogLevel is the logging level used when no level is configured.
	DefaultLogLevel = "info"
	// DefaultLogFormat is the output format used when no format is configured.
	DefaultLogFormat = "text"
)

var validLogFormats = map[string]struct{}{
	"text":   {},
	"json":   {},
	"logfmt": {},
}

// LogConfig holds global logging settings.
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Validate checks that the configured level and formatter are supported.
func (c LogConfig) Validate() error {
	if _, err := log.ParseLevel(c.Level); err != nil {
		return fmt.Errorf("log.level must be one of debug, info, warn, error, fatal - got: %s", c.Level)
	}
	if _, ok := validLogFormats[c.Format]; !ok {
		return fmt.Errorf("log.format must be one of text, json, logfmt - got: %s", c.Format)
	}
	return nil
}
