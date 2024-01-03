package feed_test

import (
	"testing"

	"github.com/tuihub/librarian/app/porter/internal/client"
	"github.com/tuihub/librarian/app/porter/internal/client/feed"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/stretchr/testify/require"
)

func getURL() string {
	return "https://github.com/TuiHub/Librarian/releases.atom"
}

func TestRSS(t *testing.T) {
	r, _ := feed.NewRSSRepo(client.NewColly())
	data, err := r.Get(getURL())
	require.NoError(t, err)
	res, err := r.Parse(data)
	logger.Infof("res: %+v", res)
	require.NoError(t, err)
}
