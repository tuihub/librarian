package internal_test

import (
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-rss/internal"

	"github.com/stretchr/testify/require"
)

func getURL() string {
	return "https://www.miit.gov.cn/api-gateway/jpaas-plugins-web-server/front/rss/getinfo?webId=8d828e408d90447786ddbe128d495e9e&columnIds=d3e2bede1bc045e2875fc7161c01db7d"
}

func TestRSS(t *testing.T) {
	r := internal.NewRSS()
	data, err := r.Get(getURL())
	require.NoError(t, err)
	res, err := r.Parse(data)
	logger.Infof("res: %+v", res)
	require.NoError(t, err)
}
