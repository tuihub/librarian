package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
)

const (
	baseURL     = "https://api.bgm.tv"
	userAgent   = "tuihub-bangumi/1.0"
	timeout     = 30 * time.Second
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
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&subject); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	
	return &subject, nil
}

func (c *Client) SearchSubjects(ctx context.Context, query string, limit int) (*model.SearchSubjectsResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 25 {
		limit = 25 // Bangumi API limit
	}
	
	params := url.Values{}
	params.Set("keyword", query)
	params.Set("limit", fmt.Sprintf("%d", limit))
	
	url := fmt.Sprintf("%s/v0/search/subjects?%s", baseURL, params.Encode())
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	
	return &searchResp, nil
}