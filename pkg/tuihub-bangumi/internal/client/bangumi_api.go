package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
)

const (
	baseURL            = "https://api.bgm.tv"
	userAgent          = "tuihub-bangumi/1.0"
	timeout            = 30 * time.Second
	defaultSearchLimit = 10
	maxSearchLimit     = 25
)

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetSubject(ctx context.Context, subjectID string) (*model.Subject, error) {
	url := fmt.Sprintf("%s/v0/subjects/%s", baseURL, subjectID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var subject model.Subject
	if decodeErr := json.NewDecoder(resp.Body).Decode(&subject); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode response: %w", decodeErr)
	}

	return &subject, nil
}

func (c *Client) SearchSubjects(ctx context.Context, query string, limit int) (*model.SearchSubjectsResponse, error) {
	if limit <= 0 {
		limit = defaultSearchLimit
	}
	if limit > maxSearchLimit {
		limit = maxSearchLimit // Bangumi API limit
	}

	params := url.Values{}
	params.Set("keyword", query)
	params.Set("limit", strconv.Itoa(limit))

	url := fmt.Sprintf("%s/v0/search/subjects?%s", baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var searchResp model.SearchSubjectsResponse
	if decodeErr := json.NewDecoder(resp.Body).Decode(&searchResp); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode response: %w", decodeErr)
	}

	return &searchResp, nil
}
