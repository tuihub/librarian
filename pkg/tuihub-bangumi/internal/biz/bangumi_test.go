package biz_test

import (
	"context"
	"os"
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/biz"
)

func getToken() string {
	return os.Getenv("BANGUMI_API_TOKEN")
}

func getSubjectID() string {
	return "12"  // Cowboy Bebop - a well-known anime for testing
}

func TestBangumiUseCase_GetSubject(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}
	
	uc := biz.NewBangumiUseCase(token)
	res, err := uc.GetSubject(context.Background(), getSubjectID())
	logger.Infof("res %+v, err: %+v", res, err)
}

func TestBangumiUseCase_SearchSubjects(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}
	
	uc := biz.NewBangumiUseCase(token)
	res, err := uc.SearchSubjects(context.Background(), "cowboy bebop")
	logger.Infof("res %+v, err: %+v", res, err)
}