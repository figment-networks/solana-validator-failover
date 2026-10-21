package logging

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var logFormatters = map[string]log.Formatter{
	"text":   log.TextFormatter,
	"json":   log.JSONFormatter,
	"logfmt": log.LogfmtFormatter,
}

// New returns a logger with the given component as its prefix.
func New(component string) *log.Logger {
	return log.WithPrefix(component)
}

// Configure sets the global logger from config, with an optional CLI level override.
// An invalid CLI override is reported and the validated config level is retained.
func Configure(configLevel, format, cliLevel string) {
	parsedLevel, err := log.ParseLevel(configLevel)
	if err != nil {
		log.Error("invalid config log level, defaulting to info", "level", configLevel, "err", err)
		parsedLevel = log.InfoLevel
	}

	if cliLevel != "" {
		cliParsedLevel, cliErr := log.ParseLevel(cliLevel)
		if cliErr != nil {
			// Ensure this error is visible even if a previous logger configuration used
			// a level higher than error. The config level is applied immediately below.
			log.SetLevel(log.InfoLevel)
			log.Error("invalid CLI log level, using configured level", "invalid_level", cliLevel, "configured_level", configLevel, "err", cliErr)
		} else {
			parsedLevel = cliParsedLevel
		}
	}

	formatter, ok := logFormatters[format]
	if !ok {
		log.Error("invalid config log format, defaulting to text", "format", format)
		formatter = log.TextFormatter
	}

	log.SetLevel(parsedLevel)
	log.SetFormatter(formatter)
	log.SetTimeFunction(func(t time.Time) time.Time { return t.UTC() })
	log.SetTimeFormat("2006-01-02T15:04:05.000Z07:00")

	styles := log.DefaultStyles()
	styles.Timestamp = lipgloss.NewStyle().Faint(true)
	styles.Message = lipgloss.NewStyle().Foreground(lipgloss.Color("213"))
	styles.Value = lipgloss.NewStyle().Foreground(lipgloss.Color("105"))
	styles.Levels[log.DebugLevel] = styles.Levels[log.DebugLevel].Foreground(lipgloss.Color("86"))
	styles.Levels[log.InfoLevel] = styles.Levels[log.InfoLevel].Foreground(lipgloss.Color("82"))
	styles.Levels[log.WarnLevel] = styles.Levels[log.WarnLevel].Foreground(lipgloss.Color("226"))
	styles.Levels[log.ErrorLevel] = styles.Levels[log.ErrorLevel].Foreground(lipgloss.Color("196"))
	styles.Levels[log.FatalLevel] = styles.Levels[log.FatalLevel].Foreground(lipgloss.Color("208"))
	log.SetStyles(styles)
}
