package bizyesod

import (
	"context"
	"net/url"
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/PuerkitoBio/goquery"
)

const maxImgNum = 9
const maxDescLen = 128
const simpleKeywordFilterActionID = "simple_keyword_filter"
const keywordFilterActionID = "keyword_filter"
const descriptionGeneratorActionID = "description_generator"

func RequiredStartAction(ctx context.Context, item *modelfeed.Item) (*modelfeed.Item, error) {
	return parsePublishPlatform(ctx, item)
}

func GetBuiltinActionMap(
	ctx context.Context,
) map[string]func(context.Context, *modelsupervisor.FeatureRequest, *modelfeed.Item) (*modelfeed.Item, error) {
	return map[string]func(context.Context, *modelsupervisor.FeatureRequest, *modelfeed.Item) (*modelfeed.Item, error){
		simpleKeywordFilterActionID:  simpleKeywordFilterAction,
		keywordFilterActionID:        keywordFilterAction,
		descriptionGeneratorActionID: descriptionGeneratorAction,
	}
}

func getBuiltinActionFeatureFlags() ([]*modelsupervisor.FeatureFlag, error) {
	simple, err := modelyesod.GetSimpleKeywordFilterActionConfigSchema()
	if err != nil {
		return nil, err
	}
	// keyword, err := modelyesod.GetKeywordFilterActionConfigSchema()
	// if err != nil {
	//	return nil, err
	//}
	desc, err := modelyesod.GetDescriptionGeneratorActionConfigSchema()
	if err != nil {
		return nil, err
	}
	return []*modelsupervisor.FeatureFlag{
		{
			ID:               simpleKeywordFilterActionID,
			Name:             "Simple Keyword Filter",
			Description:      "Filter feed item by keyword",
			ConfigJSONSchema: simple,
		},
		//{
		//	ID:               keywordFilterActionID,
		//	Name:             "Keyword Filter",
		//	Description:      "Filter feed item by keyword",
		//	ConfigJSONSchema: keyword,
		// },
		{
			ID:               descriptionGeneratorActionID,
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

func simpleKeywordFilterAction(
	_ context.Context,
	request *modelsupervisor.FeatureRequest,
	item *modelfeed.Item,
) (*modelfeed.Item, error) {
	config := new(modelyesod.SimpleKeywordFilterActionConfig)
	if err := libcodec.Unmarshal(libcodec.JSON, []byte(request.ConfigJSON), config); err != nil {
		return nil, err
	}
	for _, titleInclude := range config.TitleInclude {
		if !strings.Contains(item.Title, titleInclude) {
			return nil, nil //nolint:nilnil // return nil to skip this item
		}
	}
	for _, titleExclude := range config.TitleExclude {
		if strings.Contains(item.Title, titleExclude) {
			return nil, nil //nolint:nilnil // return nil to skip this item
		}
	}
	for _, contentInclude := range config.ContentInclude {
		if !strings.Contains(item.Content, contentInclude) {
			return nil, nil //nolint:nilnil // return nil to skip this item
		}
	}
	for _, contentExclude := range config.ContentExclude {
		if strings.Contains(item.Content, contentExclude) {
			return nil, nil //nolint:nilnil // return nil to skip this item
		}
	}
	return item, nil
}

func keywordFilterAction(ctx context.Context, _ *modelsupervisor.FeatureRequest, item *modelfeed.Item) (*modelfeed.Item, error) {
	// TODO: impl
	return item, nil
}

func descriptionGeneratorAction(_ context.Context, _ *modelsupervisor.FeatureRequest, item *modelfeed.Item) (*modelfeed.Item, error) {
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
