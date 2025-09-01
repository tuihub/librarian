package client_test

import (
	"context"
	"os"
	"testing"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/client"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
)

func getToken() string {
	return os.Getenv("BANGUMI_API_TOKEN")
}

func getSubjectID() string {
	return "12" // Cowboy Bebop - a well-known anime for testing
}

func TestClient_GetSubject(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}

	c := client.NewClient(token)
	res, err := c.GetSubject(context.Background(), getSubjectID())
	logger.Infof("res %+v, err: %+v", res, err)
}

func TestClient_SearchSubjects(t *testing.T) {
	token := getToken()
	if token == "" {
		t.Skip("BANGUMI_API_TOKEN not set")
	}

	c := client.NewClient(token)
	res, err := c.SearchSubjects(context.Background(), "cowboy bebop", 5)
	logger.Infof("res %+v, err: %+v", res, err)
}
