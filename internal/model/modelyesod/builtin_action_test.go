package modelyesod_test

import (
	"testing"

	"github.com/tuihub/librarian/internal/model/modelyesod"
)

func TestGetSimpleKeywordFilterActionConfigSchema(t *testing.T) {
	s, err := modelyesod.GetSimpleKeywordFilterActionConfigSchema()
	t.Log(s)
	if err != nil {
		t.Errorf("GetSimpleKeywordFilterActionConfigSchema() error = %v", err)
	}
}

func TestGetKeywordFilterActionConfigSchema(t *testing.T) {
	s, err := modelyesod.GetKeywordFilterActionConfigSchema()
	t.Log(s)
	if err != nil {
		t.Errorf("GetKeywordFilterActionConfigSchema() error = %v", err)
	}
}

func TestGetDescriptionGeneratorActionConfigSchema(t *testing.T) {
	s, err := modelyesod.GetDescriptionGeneratorActionConfigSchema()
	t.Log(s)
	if err != nil {
		t.Errorf("GetDescriptionGeneratorActionConfigSchema() error = %v", err)
	}
}
