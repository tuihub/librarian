package angelaweb

import (
	"context"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/api"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb/internal/page"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_AngelaWeb(t *testing.T) {
	auth, err := libauth.NewAuth(&conf.Auth{
		PasswordSalt: "",
		JwtIssuer:    "",
		JwtSecret:    "",
	})
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	admin.Password = string(hashedPassword)
	err = db.AutoMigrate(&admin)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Create(&admin).Error
	if err != nil {
		t.Fatal(err)
	}

	handler := api.NewHandler(db, auth)
	builder := page.NewBuilder(db)
	angelaWeb := NewAngelaWeb(handler, builder, auth)

	err = angelaWeb.Start(context.Background())

	assert.NoError(t, err)
}
