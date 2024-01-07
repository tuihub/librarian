package internal_test

import (
	"testing"

	"github.com/tuihub/librarian/logger"
	"github.com/tuihub/librarian/pkg/porter-rss/internal"

	"github.com/stretchr/testify/require"
)

func getURL() string {
	return "https://github.com/TuiHub/Librarian/releases.atom"
}

func TestRSS(t *testing.T) {
	r := internal.NewRSS()
	data, err := r.Get(getURL())
	require.NoError(t, err)
	res, err := r.Parse(data)
	logger.Infof("res: %+v", res)
	require.NoError(t, err)
}
