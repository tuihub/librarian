package modelyesod_test

import (
	"testing"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
)

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
