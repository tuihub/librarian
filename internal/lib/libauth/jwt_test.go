package libauth_test

import (
	"context"
	"testing"

	"github.com/tuihub/librarian/internal/lib/libauth"
)

func Test_RawToContext(t *testing.T) {
	token := "abcd"
	ctx := context.Background()
	if token != libauth.RawFromContext(libauth.RawToContext(ctx, token)) {
		t.Error("RawToContext failed")
	}
}
