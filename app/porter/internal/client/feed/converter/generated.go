// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package converter

import (
	gofeed "github.com/mmcdole/gofeed"
	modelfeed "github.com/tuihub/librarian/internal/model/modelfeed"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ToPBFeed(source *gofeed.Feed) *modelfeed.Feed {
	var pModelfeedFeed *modelfeed.Feed
	if source != nil {
		var modelfeedFeed modelfeed.Feed
		modelfeedFeed.Title = (*source).Title
		modelfeedFeed.Description = (*source).Description
		modelfeedFeed.Link = (*source).Link
		var pModelfeedPersonList []*modelfeed.Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*modelfeed.Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pGofeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedFeed.Authors = pModelfeedPersonList
		modelfeedFeed.Language = (*source).Language
		modelfeedFeed.Image = c.pGofeedImageToPModelfeedImage((*source).Image)
		var pModelfeedItemList []*modelfeed.Item
		if (*source).Items != nil {
			pModelfeedItemList = make([]*modelfeed.Item, len((*source).Items))
			for j := 0; j < len((*source).Items); j++ {
				pModelfeedItemList[j] = c.ToPBFeedItem((*source).Items[j])
			}
		}
		modelfeedFeed.Items = pModelfeedItemList
		modelfeedFeed.FeedType = (*source).FeedType
		modelfeedFeed.FeedVersion = (*source).FeedVersion
		pModelfeedFeed = &modelfeedFeed
	}
	return pModelfeedFeed
}
func (c *ConverterImpl) ToPBFeedItem(source *gofeed.Item) *modelfeed.Item {
	var pModelfeedItem *modelfeed.Item
	if source != nil {
		var modelfeedItem modelfeed.Item
		modelfeedItem.Title = (*source).Title
		modelfeedItem.Description = (*source).Description
		modelfeedItem.Content = (*source).Content
		modelfeedItem.Link = (*source).Link
		modelfeedItem.Updated = (*source).Updated
		modelfeedItem.UpdatedParsed = TimeToTime((*source).UpdatedParsed)
		modelfeedItem.Published = (*source).Published
		modelfeedItem.PublishedParsed = TimeToTime((*source).PublishedParsed)
		var pModelfeedPersonList []*modelfeed.Person
		if (*source).Authors != nil {
			pModelfeedPersonList = make([]*modelfeed.Person, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pModelfeedPersonList[i] = c.pGofeedPersonToPModelfeedPerson((*source).Authors[i])
			}
		}
		modelfeedItem.Authors = pModelfeedPersonList
		modelfeedItem.GUID = (*source).GUID
		modelfeedItem.Image = c.pGofeedImageToPModelfeedImage((*source).Image)
		var pModelfeedEnclosureList []*modelfeed.Enclosure
		if (*source).Enclosures != nil {
			pModelfeedEnclosureList = make([]*modelfeed.Enclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pModelfeedEnclosureList[j] = c.pGofeedEnclosureToPModelfeedEnclosure((*source).Enclosures[j])
			}
		}
		modelfeedItem.Enclosures = pModelfeedEnclosureList
		pModelfeedItem = &modelfeedItem
	}
	return pModelfeedItem
}
func (c *ConverterImpl) pGofeedEnclosureToPModelfeedEnclosure(source *gofeed.Enclosure) *modelfeed.Enclosure {
	var pModelfeedEnclosure *modelfeed.Enclosure
	if source != nil {
		var modelfeedEnclosure modelfeed.Enclosure
		modelfeedEnclosure.URL = (*source).URL
		modelfeedEnclosure.Length = (*source).Length
		modelfeedEnclosure.Type = (*source).Type
		pModelfeedEnclosure = &modelfeedEnclosure
	}
	return pModelfeedEnclosure
}
func (c *ConverterImpl) pGofeedImageToPModelfeedImage(source *gofeed.Image) *modelfeed.Image {
	var pModelfeedImage *modelfeed.Image
	if source != nil {
		var modelfeedImage modelfeed.Image
		modelfeedImage.URL = (*source).URL
		modelfeedImage.Title = (*source).Title
		pModelfeedImage = &modelfeedImage
	}
	return pModelfeedImage
}
func (c *ConverterImpl) pGofeedPersonToPModelfeedPerson(source *gofeed.Person) *modelfeed.Person {
	var pModelfeedPerson *modelfeed.Person
	if source != nil {
		var modelfeedPerson modelfeed.Person
		modelfeedPerson.Name = (*source).Name
		modelfeedPerson.Email = (*source).Email
		pModelfeedPerson = &modelfeedPerson
	}
	return pModelfeedPerson
}