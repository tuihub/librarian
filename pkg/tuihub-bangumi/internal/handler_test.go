package internal_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal"
	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/stretchr/testify/require"
)

func getToken() string {
	return os.Getenv("BANGUMI_API_TOKEN")
}

func getSubjectID() string {
	return "12" // Cowboy Bebop - a well-known anime for testing
}

func TestHandler_EnableContext(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}

	handler := internal.NewHandler()

	config := model.PorterContext{
		Token: token,
	}
	configJSON, err := json.Marshal(config)
	require.NoError(t, err)

	res, err := handler.EnableContext(context.Background(), &porter.EnableContextRequest{
		ContextId:   &librarian.InternalID{Id: 1},
		ContextJson: string(configJSON),
	})
	logger.Infof("res %+v, err: %+v", res, err)
	require.NoError(t, err)
}

func TestHandler_EnableContext_MissingToken(t *testing.T) {
	handler := internal.NewHandler()

	config := model.PorterContext{
		Token: "", // Missing token should cause error
	}
	configJSON, err := json.Marshal(config)
	require.NoError(t, err)

	res, err := handler.EnableContext(context.Background(), &porter.EnableContextRequest{
		ContextId:   &librarian.InternalID{Id: 1},
		ContextJson: string(configJSON),
	})
	logger.Infof("res %+v, err: %+v", res, err)
	require.Error(t, err)
	require.Contains(t, err.Error(), "token is required")
}

func TestHandler_GetAppInfo(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}

	handler := internal.NewHandler()

	// First enable context
	config := model.PorterContext{
		Token: token,
	}
	configJSON, err := json.Marshal(config)
	require.NoError(t, err)

	_, err = handler.EnableContext(context.Background(), &porter.EnableContextRequest{
		ContextId:   &librarian.InternalID{Id: 1},
		ContextJson: string(configJSON),
	})
	require.NoError(t, err)

	// Test get app info
	appConfig := model.GetAppInfoConfig{
		AppID: getSubjectID(),
	}
	appConfigJSON, err := json.Marshal(appConfig)
	require.NoError(t, err)

	res, err := handler.GetAppInfo(context.Background(), &porter.GetAppInfoRequest{
		Config: &librarian.FeatureRequest{
			Id:         "bangumi",
			Region:     "",
			ContextId:  &librarian.InternalID{Id: 1},
			ConfigJson: string(appConfigJSON),
		},
	})
	logger.Infof("res %+v, err: %+v", res, err)
	require.NoError(t, err)
	require.NotNil(t, res.GetAppInfo())
}

func TestHandler_SearchAppInfo(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}

	handler := internal.NewHandler()

	// First enable context
	config := model.PorterContext{
		Token: token,
	}
	configJSON, err := json.Marshal(config)
	require.NoError(t, err)

	_, err = handler.EnableContext(context.Background(), &porter.EnableContextRequest{
		ContextId:   &librarian.InternalID{Id: 1},
		ContextJson: string(configJSON),
	})
	require.NoError(t, err)

	// Test search app info
	searchConfig := model.SearchAppInfoConfig{
		NameLike: "cowboy bebop",
	}
	searchConfigJSON, err := json.Marshal(searchConfig)
	require.NoError(t, err)

	res, err := handler.SearchAppInfo(context.Background(), &porter.SearchAppInfoRequest{
		Config: &librarian.FeatureRequest{
			Id:         "bangumi",
			Region:     "",
			ContextId:  &librarian.InternalID{Id: 1},
			ConfigJson: string(searchConfigJSON),
		},
	})
	logger.Infof("res %+v, err: %+v", res, err)
	require.NoError(t, err)
	require.NotEmpty(t, res.GetAppInfos())
}
