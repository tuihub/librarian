package bizyesod

import (
	"context"
	"net/url"
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/PuerkitoBio/goquery"
)

const maxImgNum = 9
const maxDescLen = 128
const keywordFilterActionID = "keyword_filter"
const descriptionGeneratorActionID = "description_generator"

func RequiredStartAction(ctx context.Context, item *modelfeed.Item) (*modelfeed.Item, error) {
	return parsePublishPlatform(ctx, item)
}

func GetBuiltinActionMap(
	ctx context.Context,
) map[string]func(context.Context, modeltiphereth.FeatureRequest, *modelfeed.Item) (*modelfeed.Item, error) {
	return map[string]func(context.Context, modeltiphereth.FeatureRequest, *modelfeed.Item) (*modelfeed.Item, error){
		keywordFilterActionID:        keywordFilterAction,
		descriptionGeneratorActionID: descriptionGeneratorAction,
	}
}

func getBuiltinActionFeatureFlags() ([]*modeltiphereth.FeatureFlag, error) {
	keyword, err := modelyesod.GetKeywordFilterActionConfigSchema()
	if err != nil {
		return nil, err
	}
	desc, err := modelyesod.GetDescriptionGeneratorActionConfigSchema()
	if err != nil {
		return nil, err
	}
	return []*modeltiphereth.FeatureFlag{
		{
			ID:               keywordFilterActionID,
			Region:           "",
			Name:             "Keyword Filter",
			Description:      "Filter feed item by keyword",
			ConfigJSONSchema: keyword,
		},
		{
			ID:               descriptionGeneratorActionID,
			Region:           "",
			Name:             "Description Generator",
			Description:      "Generate description from content",
			ConfigJSONSchema: desc,
		},
	}, nil
}

func RequiredEndAction(ctx context.Context, item *modelfeed.Item) (*modelfeed.Item, error) {
	return parseDigestAction(ctx, item)
}

func parsePublishPlatform(_ context.Context, item *modelfeed.Item) (*modelfeed.Item, error) {
	if len(item.Link) > 0 {
		linkParsed, err := url.Parse(item.Link)
		if err == nil {
			item.PublishPlatform = linkParsed.Host
		}
	}
	return item, nil
}

func parseDigestAction(_ context.Context, item *modelfeed.Item) (*modelfeed.Item, error) {
	content := item.Content
	if len(content) == 0 {
		content = item.Description
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, err
	}
	digestDesc := doc.Text()
	digestDesc = strings.ReplaceAll(digestDesc, " ", "")
	digestDesc = strings.ReplaceAll(digestDesc, "\n", "")
	digestDescRune := []rune(digestDesc)
	if len(digestDescRune) > maxDescLen {
		digestDescRune = digestDescRune[:maxDescLen]
	}
	digestDesc = string(digestDescRune)
	item.DigestDescription = digestDesc

	for i, n := range doc.Find("img").Nodes {
		if i == maxImgNum {
			break
		}
		image := new(modelfeed.Image)
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				image.URL = attr.Val
			}
			if attr.Key == "alt" {
				image.Title = attr.Val
			}
		}
		item.DigestImages = append(item.DigestImages, image)
	}
	return item, nil
}

func keywordFilterAction(ctx context.Context, _ modeltiphereth.FeatureRequest, item *modelfeed.Item) (*modelfeed.Item, error) {
	// TODO: impl
	return item, nil
}

func descriptionGeneratorAction(_ context.Context, _ modeltiphereth.FeatureRequest, item *modelfeed.Item) (*modelfeed.Item, error) {
	if len(item.Description) > 0 {
		return item, nil
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		return nil, err
	}
	desc := doc.Text()
	desc = strings.ReplaceAll(desc, " ", "")
	desc = strings.ReplaceAll(desc, "\n", "")
	descRune := []rune(desc)
	if len(descRune) > maxDescLen {
		descRune = descRune[:maxDescLen]
	}
	item.Description = string(descRune)
	return item, nil
}
