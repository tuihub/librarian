package angelaweb

import (
	"context"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_AngelaWeb(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	handler := api.NewHandler(db)
	angelaWeb := NewAngelaWeb(handler)

	err = angelaWeb.Start(context.Background())

	assert.NoError(t, err)
}
