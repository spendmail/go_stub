package app

import (
	_ "image/jpeg"
	"testing"

	internalconfig "github.com/spendmail/stub/internal/config"
	internallogger "github.com/spendmail/stub/internal/logger"
	"github.com/stretchr/testify/require"
)

func TestApplication(t *testing.T) {
	t.Run("do test", func(t *testing.T) {
		config, err := internalconfig.New("../../configs/stub.toml")
		require.NoError(t, err, "should be without errors")

		logger, err := internallogger.New(config)
		require.NoError(t, err, "should be without errors")

		app, err := New(logger, config)
		require.NoError(t, err, "should be without errors")

		bytes, err := app.StubMethod(100, "stringStubParam", "anyStubParam", map[string][]string{})
		require.NoError(t, err, "should be without errors")
		require.Equal(t, bytes, []byte{})
	})
}
