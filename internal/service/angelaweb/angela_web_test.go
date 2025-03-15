package angelaweb

import (
	"context"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"
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

	admin := model.User{
		Username: "admin",
		Password: "admin123",
		Email:    "admin@example.com",
		Role:     "admin",
	}
	err = db.Create(&admin).Error
	if err != nil {
		t.Fatal(err)
	}

	handler := api.NewHandler(db)
	builder := page.NewBuilder(db)
	angelaWeb := NewAngelaWeb(handler, builder)

	err = angelaWeb.Start(context.Background())

	assert.NoError(t, err)
}
