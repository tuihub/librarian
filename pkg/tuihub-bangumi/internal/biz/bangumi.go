package biz

import (
	"context"
	"strings"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/client"
	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
)

const (
	defaultSearchLimit = 10
)

type BangumiUseCase struct {
	client *client.Client
}

func NewBangumiUseCase(token string) *BangumiUseCase {
	return &BangumiUseCase{
		client: client.NewClient(token),
	}
}

func (b *BangumiUseCase) GetSubject(ctx context.Context, subjectID string) (*AppInfo, error) {
	subject, err := b.client.GetSubject(ctx, subjectID)
	if err != nil {
		return nil, err
	}

	return b.convertSubjectToAppInfo(subject), nil
}

func (b *BangumiUseCase) SearchSubjects(ctx context.Context, query string) ([]*AppInfo, error) {
	searchResp, err := b.client.SearchSubjects(ctx, query, defaultSearchLimit)
	if err != nil {
		return nil, err
	}

	apps := make([]*AppInfo, len(searchResp.List))
	for i, subject := range searchResp.List {
		apps[i] = b.convertSubjectToAppInfo(&subject)
	}

	return apps, nil
}

func (b *BangumiUseCase) convertSubjectToAppInfo(subject *model.Subject) *AppInfo {
	appType := b.getAppType(model.SubjectType(subject.Type))

	// Extract developer/publisher from infobox
	var developer, publisher []string
	for _, info := range subject.InfoBox {
		switch strings.ToLower(info.Key) {
		case "原作", "原案", "脚本", "監督", "制作":
			if info.Value != "" {
				developer = append(developer, info.Value)
			}
		case "製作", "出版社", "发行":
			if info.Value != "" {
				publisher = append(publisher, info.Value)
			}
		}
	}

	// Create image URLs list
	var imageURLs []string
	if subject.Images.Large != "" {
		imageURLs = append(imageURLs, subject.Images.Large)
	}
	if subject.Images.Medium != "" {
		imageURLs = append(imageURLs, subject.Images.Medium)
	}

	return &AppInfo{
		ID:                 subject.ID,
		Name:               subject.Name,
		NameCN:             subject.NameCN,
		Type:               appType,
		ShortDescription:   subject.Summary,
		Description:        subject.Summary,
		ReleaseDate:        subject.Date,
		Developer:          strings.Join(developer, ", "),
		Publisher:          strings.Join(publisher, ", "),
		CoverImageURL:      subject.Images.Large,
		BackgroundImageURL: subject.Images.Common,
		ImageURLs:          imageURLs,
		StoreURL:           subject.URL,
		Tags:               b.extractTags(subject.Tags),
	}
}

func (b *BangumiUseCase) getAppType(subjectType model.SubjectType) AppType {
	switch subjectType {
	case model.SubjectTypeGame:
		return AppTypeGame
	case model.SubjectTypeAnime:
		return AppTypeAnime
	case model.SubjectTypeBook:
		return AppTypeBook
	case model.SubjectTypeMusic:
		return AppTypeMusic
	default:
		return AppTypeUnspecified
	}
}

func (b *BangumiUseCase) extractTags(tags []model.Tag) []string {
	result := make([]string, len(tags))
	for i, tag := range tags {
		result[i] = tag.Name
	}
	return result
}

// AppInfo represents processed Bangumi subject data.
type AppInfo struct {
	ID                 int
	Name               string
	NameCN             string
	Type               AppType
	ShortDescription   string
	Description        string
	ReleaseDate        string
	Developer          string
	Publisher          string
	CoverImageURL      string
	BackgroundImageURL string
	ImageURLs          []string
	StoreURL           string
	Tags               []string
}

// AppType represents the type of app/media.
type AppType int

const (
	AppTypeUnspecified AppType = iota
	AppTypeGame
	AppTypeAnime
	AppTypeBook
	AppTypeMusic
	AppTypeReal
)
