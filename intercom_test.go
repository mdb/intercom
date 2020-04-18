package intercom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLogger_silent_level(t *testing.T) {
	logger := NewLogger("silent")

	assert.Equal(t, logger.Level, silentLevel)
}

func Test_NewLogger_error_level(t *testing.T) {
	logger := NewLogger("error")

	assert.Equal(t, logger.Level, errorLevel)
}

func Test_NewLogger_warn_level(t *testing.T) {
	logger := NewLogger("warn")

	assert.Equal(t, logger.Level, warnLevel)
}

func Test_NewLogger_info_level(t *testing.T) {
	logger := NewLogger("info")

	assert.Equal(t, logger.Level, infoLevel)
}

func Test_NewLogger_unrecognized_level(t *testing.T) {
	logger := NewLogger("foo")

	assert.Equal(t, logger.Level, infoLevel)
}

func Test_NewLogger_debug_level(t *testing.T) {
	logger := NewLogger("debug")

	assert.Equal(t, logger.Level, debugLevel)
}

func Test_Errorf(t *testing.T) {
	logger := NewLogger("info")
	bar := "bar"

	output := logger.Errorf("foo %s", bar)

	assert.Equal(t, output, "foo bar")
}
