package tuihubbangumi_test

import (
	"context"
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/stretchr/testify/require"
)

func TestNewPorter(t *testing.T) {
	p, err := tuihubbangumi.NewPorter("test-version")
	require.NoError(t, err)
	require.NotNil(t, p)
	
	// Test that the porter service can be retrieved
	service := p.GetPorterService()
	require.NotNil(t, service)
	
	// Test that GetPorterInformation works
	info, err := service.GetPorterInformation(context.Background(), &porter.GetPorterInformationRequest{})
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Equal(t, "tuihub-bangumi", info.BinarySummary.Name)
	require.Equal(t, "test-version", info.BinarySummary.Version)
	require.Equal(t, "github.com/tuihub/librarian/pkg/tuihub-bangumi", info.GlobalName)
	
	// Test that it has the Bangumi app info source feature
	require.NotNil(t, info.FeatureSummary)
	require.NotEmpty(t, info.FeatureSummary.AppInfoSources)
	
	bangumiFeature := info.FeatureSummary.AppInfoSources[0]
	require.Equal(t, "bangumi", bangumiFeature.Id)
	require.Equal(t, "Bangumi", bangumiFeature.Name)
	require.True(t, bangumiFeature.RequireContext)
	
	// Test that it has a context schema
	require.NotNil(t, info.ContextJsonSchema)
}