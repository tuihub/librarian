package model_test

import (
	"encoding/json"
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"

	"github.com/stretchr/testify/require"
)

func TestPorterContext_Marshaling(t *testing.T) {
	config := model.PorterContext{
		Token: "test-token",
	}

	data, err := json.Marshal(config)
	require.NoError(t, err)

	var unmarshaled model.PorterContext
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	require.Equal(t, config.Token, unmarshaled.Token)
}

func TestGetAppInfoConfig_Marshaling(t *testing.T) {
	config := model.GetAppInfoConfig{
		AppID: "12345",
	}

	data, err := json.Marshal(config)
	require.NoError(t, err)

	var unmarshaled model.GetAppInfoConfig
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	require.Equal(t, config.AppID, unmarshaled.AppID)
}

func TestSearchAppInfoConfig_Marshaling(t *testing.T) {
	config := model.SearchAppInfoConfig{
		NameLike: "cowboy bebop",
	}

	data, err := json.Marshal(config)
	require.NoError(t, err)

	var unmarshaled model.SearchAppInfoConfig
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	require.Equal(t, config.NameLike, unmarshaled.NameLike)
}

func TestSubjectType_Constants(t *testing.T) {
	require.Equal(t, model.SubjectTypeBook, model.SubjectType(1))
	require.Equal(t, model.SubjectTypeAnime, model.SubjectType(2))
	require.Equal(t, model.SubjectTypeMusic, model.SubjectType(3))
	require.Equal(t, model.SubjectTypeGame, model.SubjectType(4))
	require.Equal(t, model.SubjectTypeReal, model.SubjectType(6))
}
