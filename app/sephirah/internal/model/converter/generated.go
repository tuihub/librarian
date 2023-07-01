// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package converter

import (
	modelbinah "github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	modelgebura "github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	modelnetzach "github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	modeltiphereth "github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	modelyesod "github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	libauth "github.com/tuihub/librarian/internal/lib/libauth"
	model "github.com/tuihub/librarian/internal/model"
	modelfeed "github.com/tuihub/librarian/internal/model/modelfeed"
	v11 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	v1 "github.com/tuihub/protos/pkg/librarian/v1"
)

type toBizConverterImpl struct{}

func (c *toBizConverterImpl) ToBizApp(source *v1.App) *modelgebura.App {
	var pModelgeburaApp *modelgebura.App
	if source != nil {
		var modelgeburaApp modelgebura.App
		modelgeburaApp.ID = ToBizInternalID((*source).Id)
		modelgeburaApp.Source = ToBizAppSource((*source).Source)
		modelgeburaApp.SourceAppID = (*source).SourceAppId
		modelgeburaApp.SourceURL = PtrToString((*source).SourceUrl)
		modelgeburaApp.Name = (*source).Name
		modelgeburaApp.Type = ToBizAppType((*source).Type)
		modelgeburaApp.ShortDescription = (*source).ShortDescription
		modelgeburaApp.ImageURL = (*source).ImageUrl
		modelgeburaApp.Details = c.pV1AppDetailsToPModelgeburaAppDetails((*source).Details)
		pModelgeburaApp = &modelgeburaApp
	}
	return pModelgeburaApp
}
func (c *toBizConverterImpl) ToBizAppPackage(source *v1.AppPackage) *modelgebura.AppPackage {
	var pModelgeburaAppPackage *modelgebura.AppPackage
	if source != nil {
		var modelgeburaAppPackage modelgebura.AppPackage
		modelgeburaAppPackage.ID = ToBizInternalID((*source).Id)
		modelgeburaAppPackage.Source = ToBizAppPackageSource((*source).Source)
		modelgeburaAppPackage.SourceID = ToBizInternalID((*source).SourceId)
		modelgeburaAppPackage.Name = (*source).Name
		modelgeburaAppPackage.Description = (*source).Description
		modelgeburaAppPackage.Binary = c.ToBizAppPackageBinary((*source).Binary)
		modelgeburaAppPackage.Public = (*source).Public
		pModelgeburaAppPackage = &modelgeburaAppPackage
	}
	return pModelgeburaAppPackage
}
func (c *toBizConverterImpl) ToBizAppPackageBinary(source *v1.AppPackageBinary) *modelgebura.AppPackageBinary {
	var pModelgeburaAppPackageBinary *modelgebura.AppPackageBinary
	if source != nil {
		var modelgeburaAppPackageBinary modelgebura.AppPackageBinary
		modelgeburaAppPackageBinary.Name = (*source).Name
		modelgeburaAppPackageBinary.SizeBytes = (*source).SizeBytes
		modelgeburaAppPackageBinary.PublicURL = (*source).PublicUrl
		var byteList []uint8
		if (*source).Sha256 != nil {
			byteList = make([]uint8, len((*source).Sha256))
			for i := 0; i < len((*source).Sha256); i++ {
				byteList[i] = (*source).Sha256[i]
			}
		}
		modelgeburaAppPackageBinary.Sha256 = byteList
		pModelgeburaAppPackageBinary = &modelgeburaAppPackageBinary
	}
	return pModelgeburaAppPackageBinary
}
func (c *toBizConverterImpl) ToBizAppPackageBinaryList(source []*v1.AppPackageBinary) []*modelgebura.AppPackageBinary {
	var pModelgeburaAppPackageBinaryList []*modelgebura.AppPackageBinary
	if source != nil {
		pModelgeburaAppPackageBinaryList = make([]*modelgebura.AppPackageBinary, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaAppPackageBinaryList[i] = c.ToBizAppPackageBinary(source[i])
		}
	}
	return pModelgeburaAppPackageBinaryList
}
func (c *toBizConverterImpl) ToBizAppPackageSourceList(source []v1.AppPackageSource) []modelgebura.AppPackageSource {
	var modelgeburaAppPackageSourceList []modelgebura.AppPackageSource
	if source != nil {
		modelgeburaAppPackageSourceList = make([]modelgebura.AppPackageSource, len(source))
		for i := 0; i < len(source); i++ {
			modelgeburaAppPackageSourceList[i] = ToBizAppPackageSource(source[i])
		}
	}
	return modelgeburaAppPackageSourceList
}
func (c *toBizConverterImpl) ToBizAppSourceList(source []v1.AppSource) []modelgebura.AppSource {
	var modelgeburaAppSourceList []modelgebura.AppSource
	if source != nil {
		modelgeburaAppSourceList = make([]modelgebura.AppSource, len(source))
		for i := 0; i < len(source); i++ {
			modelgeburaAppSourceList[i] = ToBizAppSource(source[i])
		}
	}
	return modelgeburaAppSourceList
}
func (c *toBizConverterImpl) ToBizAppTypeList(source []v1.AppType) []modelgebura.AppType {
	var modelgeburaAppTypeList []modelgebura.AppType
	if source != nil {
		modelgeburaAppTypeList = make([]modelgebura.AppType, len(source))
		for i := 0; i < len(source); i++ {
			modelgeburaAppTypeList[i] = ToBizAppType(source[i])
		}
	}
	return modelgeburaAppTypeList
}
func (c *toBizConverterImpl) ToBizFeedConfig(source *v11.FeedConfig) *modelyesod.FeedConfig {
	var pModelyesodFeedConfig *modelyesod.FeedConfig
	if source != nil {
		var modelyesodFeedConfig modelyesod.FeedConfig
		modelyesodFeedConfig.ID = ToBizInternalID((*source).Id)
		modelyesodFeedConfig.Name = (*source).Name
		modelyesodFeedConfig.FeedURL = (*source).FeedUrl
		var stringList []string
		if (*source).Tags != nil {
			stringList = make([]string, len((*source).Tags))
			for i := 0; i < len((*source).Tags); i++ {
				stringList[i] = (*source).Tags[i]
			}
		}
		modelyesodFeedConfig.Tags = stringList
		modelyesodFeedConfig.AuthorAccount = ToBizInternalID((*source).AuthorAccount)
		modelyesodFeedConfig.Source = ToBizFeedConfigSource((*source).Source)
		modelyesodFeedConfig.Status = ToBizFeedConfigStatus((*source).Status)
		modelyesodFeedConfig.PullInterval = DurationPBToDuration((*source).PullInterval)
		pModelyesodFeedConfig = &modelyesodFeedConfig
	}
	return pModelyesodFeedConfig
}
func (c *toBizConverterImpl) ToBizFeedConfigSourceList(source []v11.FeedConfigSource) []modelyesod.FeedConfigSource {
	var modelyesodFeedConfigSourceList []modelyesod.FeedConfigSource
	if source != nil {
		modelyesodFeedConfigSourceList = make([]modelyesod.FeedConfigSource, len(source))
		for i := 0; i < len(source); i++ {
			modelyesodFeedConfigSourceList[i] = ToBizFeedConfigSource(source[i])
		}
	}
	return modelyesodFeedConfigSourceList
}
func (c *toBizConverterImpl) ToBizFeedConfigStatusList(source []v11.FeedConfigStatus) []modelyesod.FeedConfigStatus {
	var modelyesodFeedConfigStatusList []modelyesod.FeedConfigStatus
	if source != nil {
		modelyesodFeedConfigStatusList = make([]modelyesod.FeedConfigStatus, len(source))
		for i := 0; i < len(source); i++ {
			modelyesodFeedConfigStatusList[i] = ToBizFeedConfigStatus(source[i])
		}
	}
	return modelyesodFeedConfigStatusList
}
func (c *toBizConverterImpl) ToBizFileMetadata(source *v11.FileMetadata) *modelbinah.FileMetadata {
	var pModelbinahFileMetadata *modelbinah.FileMetadata
	if source != nil {
		var modelbinahFileMetadata modelbinah.FileMetadata
		modelbinahFileMetadata.ID = ToBizInternalID((*source).Id)
		modelbinahFileMetadata.Name = (*source).Name
		modelbinahFileMetadata.SizeBytes = (*source).SizeBytes
		modelbinahFileMetadata.Type = ToBizFileType((*source).Type)
		var byteList []uint8
		if (*source).Sha256 != nil {
			byteList = make([]uint8, len((*source).Sha256))
			for i := 0; i < len((*source).Sha256); i++ {
				byteList[i] = (*source).Sha256[i]
			}
		}
		modelbinahFileMetadata.Sha256 = byteList
		pModelbinahFileMetadata = &modelbinahFileMetadata
	}
	return pModelbinahFileMetadata
}
func (c *toBizConverterImpl) ToBizInternalIDList(source []*v1.InternalID) []model.InternalID {
	var modelInternalIDList []model.InternalID
	if source != nil {
		modelInternalIDList = make([]model.InternalID, len(source))
		for i := 0; i < len(source); i++ {
			modelInternalIDList[i] = ToBizInternalID(source[i])
		}
	}
	return modelInternalIDList
}
func (c *toBizConverterImpl) ToBizNotifyFlow(source *v11.NotifyFlow) *modelnetzach.NotifyFlow {
	var pModelnetzachNotifyFlow *modelnetzach.NotifyFlow
	if source != nil {
		var modelnetzachNotifyFlow modelnetzach.NotifyFlow
		modelnetzachNotifyFlow.ID = ToBizInternalID((*source).Id)
		modelnetzachNotifyFlow.Name = (*source).Name
		modelnetzachNotifyFlow.Description = (*source).Description
		modelnetzachNotifyFlow.Source = c.ToBizNotifyFlowSource((*source).Source)
		var pModelnetzachNotifyFlowTargetList []*modelnetzach.NotifyFlowTarget
		if (*source).Targets != nil {
			pModelnetzachNotifyFlowTargetList = make([]*modelnetzach.NotifyFlowTarget, len((*source).Targets))
			for i := 0; i < len((*source).Targets); i++ {
				pModelnetzachNotifyFlowTargetList[i] = c.ToBizNotifyFlowTarget((*source).Targets[i])
			}
		}
		modelnetzachNotifyFlow.Targets = pModelnetzachNotifyFlowTargetList
		modelnetzachNotifyFlow.Status = ToBizNotifyFlowStatus((*source).Status)
		pModelnetzachNotifyFlow = &modelnetzachNotifyFlow
	}
	return pModelnetzachNotifyFlow
}
func (c *toBizConverterImpl) ToBizNotifyFlowSource(source *v11.NotifyFlowSource) *modelnetzach.NotifyFlowSource {
	var pModelnetzachNotifyFlowSource *modelnetzach.NotifyFlowSource
	if source != nil {
		var modelnetzachNotifyFlowSource modelnetzach.NotifyFlowSource
		modelnetzachNotifyFlowSource.FeedIDFilter = c.ToBizInternalIDList((*source).FeedIdFilter)
		pModelnetzachNotifyFlowSource = &modelnetzachNotifyFlowSource
	}
	return pModelnetzachNotifyFlowSource
}
func (c *toBizConverterImpl) ToBizNotifyFlowTarget(source *v11.NotifyFlowTarget) *modelnetzach.NotifyFlowTarget {
	var pModelnetzachNotifyFlowTarget *modelnetzach.NotifyFlowTarget
	if source != nil {
		var modelnetzachNotifyFlowTarget modelnetzach.NotifyFlowTarget
		modelnetzachNotifyFlowTarget.TargetID = ToBizInternalID((*source).TargetId)
		modelnetzachNotifyFlowTarget.ChannelID = (*source).ChannelId
		pModelnetzachNotifyFlowTarget = &modelnetzachNotifyFlowTarget
	}
	return pModelnetzachNotifyFlowTarget
}
func (c *toBizConverterImpl) ToBizNotifyTarget(source *v11.NotifyTarget) *modelnetzach.NotifyTarget {
	var pModelnetzachNotifyTarget *modelnetzach.NotifyTarget
	if source != nil {
		var modelnetzachNotifyTarget modelnetzach.NotifyTarget
		modelnetzachNotifyTarget.ID = ToBizInternalID((*source).Id)
		modelnetzachNotifyTarget.Name = (*source).Name
		modelnetzachNotifyTarget.Description = (*source).Description
		modelnetzachNotifyTarget.Type = ToBizNotifyTargetType((*source).Type)
		modelnetzachNotifyTarget.Status = ToBizNotifyTargetStatus((*source).Status)
		modelnetzachNotifyTarget.Token = (*source).Token
		pModelnetzachNotifyTarget = &modelnetzachNotifyTarget
	}
	return pModelnetzachNotifyTarget
}
func (c *toBizConverterImpl) ToBizNotifyTargetStatusList(source []v11.NotifyTargetStatus) []modelnetzach.NotifyTargetStatus {
	var modelnetzachNotifyTargetStatusList []modelnetzach.NotifyTargetStatus
	if source != nil {
		modelnetzachNotifyTargetStatusList = make([]modelnetzach.NotifyTargetStatus, len(source))
		for i := 0; i < len(source); i++ {
			modelnetzachNotifyTargetStatusList[i] = ToBizNotifyTargetStatus(source[i])
		}
	}
	return modelnetzachNotifyTargetStatusList
}
func (c *toBizConverterImpl) ToBizNotifyTargetTypeList(source []v11.NotifyTargetType) []modelnetzach.NotifyTargetType {
	var modelnetzachNotifyTargetTypeList []modelnetzach.NotifyTargetType
	if source != nil {
		modelnetzachNotifyTargetTypeList = make([]modelnetzach.NotifyTargetType, len(source))
		for i := 0; i < len(source); i++ {
			modelnetzachNotifyTargetTypeList[i] = ToBizNotifyTargetType(source[i])
		}
	}
	return modelnetzachNotifyTargetTypeList
}
func (c *toBizConverterImpl) ToBizTimeRange(source *v1.TimeRange) *model.TimeRange {
	var pModelTimeRange *model.TimeRange
	if source != nil {
		var modelTimeRange model.TimeRange
		modelTimeRange.StartTime = ToBizTime((*source).StartTime)
		modelTimeRange.Duration = DurationPBToDuration((*source).Duration)
		pModelTimeRange = &modelTimeRange
	}
	return pModelTimeRange
}
func (c *toBizConverterImpl) ToBizUser(source *v11.User) *modeltiphereth.User {
	var pModeltipherethUser *modeltiphereth.User
	if source != nil {
		var modeltipherethUser modeltiphereth.User
		modeltipherethUser.ID = ToBizInternalID((*source).Id)
		modeltipherethUser.UserName = (*source).Username
		modeltipherethUser.PassWord = (*source).Password
		modeltipherethUser.Type = ToLibAuthUserType((*source).Type)
		modeltipherethUser.Status = ToBizUserStatus((*source).Status)
		pModeltipherethUser = &modeltipherethUser
	}
	return pModeltipherethUser
}
func (c *toBizConverterImpl) ToBizUserStatusList(source []v11.UserStatus) []modeltiphereth.UserStatus {
	var modeltipherethUserStatusList []modeltiphereth.UserStatus
	if source != nil {
		modeltipherethUserStatusList = make([]modeltiphereth.UserStatus, len(source))
		for i := 0; i < len(source); i++ {
			modeltipherethUserStatusList[i] = ToBizUserStatus(source[i])
		}
	}
	return modeltipherethUserStatusList
}
func (c *toBizConverterImpl) ToLibAuthUserTypeList(source []v11.UserType) []libauth.UserType {
	var libauthUserTypeList []libauth.UserType
	if source != nil {
		libauthUserTypeList = make([]libauth.UserType, len(source))
		for i := 0; i < len(source); i++ {
			libauthUserTypeList[i] = ToLibAuthUserType(source[i])
		}
	}
	return libauthUserTypeList
}
func (c *toBizConverterImpl) pV1AppDetailsToPModelgeburaAppDetails(source *v1.AppDetails) *modelgebura.AppDetails {
	var pModelgeburaAppDetails *modelgebura.AppDetails
	if source != nil {
		var modelgeburaAppDetails modelgebura.AppDetails
		modelgeburaAppDetails.Description = (*source).Description
		modelgeburaAppDetails.ReleaseDate = (*source).ReleaseDate
		modelgeburaAppDetails.Developer = (*source).Developer
		modelgeburaAppDetails.Publisher = (*source).Publisher
		modelgeburaAppDetails.Version = (*source).Version
		pModelgeburaAppDetails = &modelgeburaAppDetails
	}
	return pModelgeburaAppDetails
}

type toPBConverterImpl struct{}

func (c *toPBConverterImpl) ToPBAccount(source *modeltiphereth.Account) *v1.Account {
	var pV1Account *v1.Account
	if source != nil {
		var v1Account v1.Account
		v1Account.Id = ToPBInternalID((*source).ID)
		v1Account.Platform = ToPBAccountPlatform((*source).Platform)
		v1Account.PlatformAccountId = (*source).PlatformAccountID
		v1Account.Name = (*source).Name
		v1Account.ProfileUrl = (*source).ProfileURL
		v1Account.AvatarUrl = (*source).AvatarURL
		pV1Account = &v1Account
	}
	return pV1Account
}
func (c *toPBConverterImpl) ToPBAccountList(source []*modeltiphereth.Account) []*v1.Account {
	var pV1AccountList []*v1.Account
	if source != nil {
		pV1AccountList = make([]*v1.Account, len(source))
		for i := 0; i < len(source); i++ {
			pV1AccountList[i] = c.ToPBAccount(source[i])
		}
	}
	return pV1AccountList
}
func (c *toPBConverterImpl) ToPBApp(source *modelgebura.App) *v1.App {
	var pV1App *v1.App
	if source != nil {
		var v1App v1.App
		v1App.Id = ToPBInternalID((*source).ID)
		v1App.Source = ToPBAppSource((*source).Source)
		v1App.SourceAppId = (*source).SourceAppID
		pString := (*source).SourceURL
		v1App.SourceUrl = &pString
		v1App.Name = (*source).Name
		v1App.Type = ToPBAppType((*source).Type)
		v1App.ShortDescription = (*source).ShortDescription
		v1App.ImageUrl = (*source).ImageURL
		v1App.Details = c.pModelgeburaAppDetailsToPV1AppDetails((*source).Details)
		pV1App = &v1App
	}
	return pV1App
}
func (c *toPBConverterImpl) ToPBAppList(source []*modelgebura.App) []*v1.App {
	var pV1AppList []*v1.App
	if source != nil {
		pV1AppList = make([]*v1.App, len(source))
		for i := 0; i < len(source); i++ {
			pV1AppList[i] = c.ToPBApp(source[i])
		}
	}
	return pV1AppList
}
func (c *toPBConverterImpl) ToPBAppPackage(source *modelgebura.AppPackage) *v1.AppPackage {
	var pV1AppPackage *v1.AppPackage
	if source != nil {
		var v1AppPackage v1.AppPackage
		v1AppPackage.Id = ToPBInternalID((*source).ID)
		v1AppPackage.Source = ToPBAppPackageSource((*source).Source)
		v1AppPackage.SourceId = ToPBInternalID((*source).SourceID)
		v1AppPackage.Name = (*source).Name
		v1AppPackage.Description = (*source).Description
		v1AppPackage.Binary = c.ToPBAppPackageBinary((*source).Binary)
		v1AppPackage.Public = (*source).Public
		pV1AppPackage = &v1AppPackage
	}
	return pV1AppPackage
}
func (c *toPBConverterImpl) ToPBAppPackageBinary(source *modelgebura.AppPackageBinary) *v1.AppPackageBinary {
	var pV1AppPackageBinary *v1.AppPackageBinary
	if source != nil {
		var v1AppPackageBinary v1.AppPackageBinary
		v1AppPackageBinary.Name = (*source).Name
		v1AppPackageBinary.SizeBytes = (*source).SizeBytes
		v1AppPackageBinary.PublicUrl = (*source).PublicURL
		var byteList []uint8
		if (*source).Sha256 != nil {
			byteList = make([]uint8, len((*source).Sha256))
			for i := 0; i < len((*source).Sha256); i++ {
				byteList[i] = (*source).Sha256[i]
			}
		}
		v1AppPackageBinary.Sha256 = byteList
		pV1AppPackageBinary = &v1AppPackageBinary
	}
	return pV1AppPackageBinary
}
func (c *toPBConverterImpl) ToPBAppPackageList(source []*modelgebura.AppPackage) []*v1.AppPackage {
	var pV1AppPackageList []*v1.AppPackage
	if source != nil {
		pV1AppPackageList = make([]*v1.AppPackage, len(source))
		for i := 0; i < len(source); i++ {
			pV1AppPackageList[i] = c.ToPBAppPackage(source[i])
		}
	}
	return pV1AppPackageList
}
func (c *toPBConverterImpl) ToPBEnclosure(source *modelfeed.Enclosure) *v1.FeedEnclosure {
	var pV1FeedEnclosure *v1.FeedEnclosure
	if source != nil {
		var v1FeedEnclosure v1.FeedEnclosure
		v1FeedEnclosure.Url = (*source).URL
		v1FeedEnclosure.Length = (*source).Length
		v1FeedEnclosure.Type = (*source).Type
		pV1FeedEnclosure = &v1FeedEnclosure
	}
	return pV1FeedEnclosure
}
func (c *toPBConverterImpl) ToPBFeed(source *modelfeed.Feed) *v1.Feed {
	var pV1Feed *v1.Feed
	if source != nil {
		var v1Feed v1.Feed
		v1Feed.Id = ToPBInternalID((*source).ID)
		v1Feed.Title = (*source).Title
		v1Feed.Link = (*source).Link
		v1Feed.Description = (*source).Description
		v1Feed.Items = c.ToPBFeedItemList((*source).Items)
		v1Feed.Language = (*source).Language
		v1Feed.Image = c.ToPBFeedImage((*source).Image)
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pV1FeedPersonList[i] = c.pModelfeedPersonToPV1FeedPerson((*source).Authors[i])
			}
		}
		v1Feed.Authors = pV1FeedPersonList
		pV1Feed = &v1Feed
	}
	return pV1Feed
}
func (c *toPBConverterImpl) ToPBFeedConfig(source *modelyesod.FeedConfig) *v11.FeedConfig {
	var pV1FeedConfig *v11.FeedConfig
	if source != nil {
		var v1FeedConfig v11.FeedConfig
		v1FeedConfig.Id = ToPBInternalID((*source).ID)
		v1FeedConfig.Name = (*source).Name
		v1FeedConfig.FeedUrl = (*source).FeedURL
		v1FeedConfig.AuthorAccount = ToPBInternalID((*source).AuthorAccount)
		v1FeedConfig.Source = ToPBFeedConfigSource((*source).Source)
		v1FeedConfig.Status = ToPBFeedConfigStatus((*source).Status)
		v1FeedConfig.PullInterval = ToPBDuration((*source).PullInterval)
		var stringList []string
		if (*source).Tags != nil {
			stringList = make([]string, len((*source).Tags))
			for i := 0; i < len((*source).Tags); i++ {
				stringList[i] = (*source).Tags[i]
			}
		}
		v1FeedConfig.Tags = stringList
		v1FeedConfig.LatestPullTime = ToPBTime((*source).LatestPullTime)
		pV1FeedConfig = &v1FeedConfig
	}
	return pV1FeedConfig
}
func (c *toPBConverterImpl) ToPBFeedImage(source *modelfeed.Image) *v1.FeedImage {
	var pV1FeedImage *v1.FeedImage
	if source != nil {
		var v1FeedImage v1.FeedImage
		v1FeedImage.Url = (*source).URL
		v1FeedImage.Title = (*source).Title
		pV1FeedImage = &v1FeedImage
	}
	return pV1FeedImage
}
func (c *toPBConverterImpl) ToPBFeedItem(source *modelfeed.Item) *v1.FeedItem {
	var pV1FeedItem *v1.FeedItem
	if source != nil {
		var v1FeedItem v1.FeedItem
		v1FeedItem.Id = ToPBInternalID((*source).ID)
		v1FeedItem.Title = (*source).Title
		var pV1FeedPersonList []*v1.FeedPerson
		if (*source).Authors != nil {
			pV1FeedPersonList = make([]*v1.FeedPerson, len((*source).Authors))
			for i := 0; i < len((*source).Authors); i++ {
				pV1FeedPersonList[i] = c.pModelfeedPersonToPV1FeedPerson((*source).Authors[i])
			}
		}
		v1FeedItem.Authors = pV1FeedPersonList
		v1FeedItem.Description = (*source).Description
		v1FeedItem.Content = (*source).Content
		v1FeedItem.Guid = (*source).GUID
		v1FeedItem.Link = (*source).Link
		v1FeedItem.Image = c.ToPBFeedImage((*source).Image)
		v1FeedItem.Published = (*source).Published
		v1FeedItem.PublishedParsed = ToPBTimePtr((*source).PublishedParsed)
		v1FeedItem.Updated = (*source).Updated
		v1FeedItem.UpdatedParsed = ToPBTimePtr((*source).UpdatedParsed)
		var pV1FeedEnclosureList []*v1.FeedEnclosure
		if (*source).Enclosures != nil {
			pV1FeedEnclosureList = make([]*v1.FeedEnclosure, len((*source).Enclosures))
			for j := 0; j < len((*source).Enclosures); j++ {
				pV1FeedEnclosureList[j] = c.ToPBEnclosure((*source).Enclosures[j])
			}
		}
		v1FeedItem.Enclosures = pV1FeedEnclosureList
		v1FeedItem.PublishPlatform = (*source).PublishPlatform
		pV1FeedItem = &v1FeedItem
	}
	return pV1FeedItem
}
func (c *toPBConverterImpl) ToPBFeedItemDigest(source *modelyesod.FeedItemDigest) *v11.FeedItemDigest {
	var pV1FeedItemDigest *v11.FeedItemDigest
	if source != nil {
		var v1FeedItemDigest v11.FeedItemDigest
		v1FeedItemDigest.FeedId = ToPBInternalID((*source).FeedID)
		v1FeedItemDigest.ItemId = ToPBInternalID((*source).ItemID)
		v1FeedItemDigest.AvatarUrl = (*source).AvatarURL
		v1FeedItemDigest.Authors = (*source).Authors
		v1FeedItemDigest.PublishedParsed = ToPBTime((*source).PublishedParsed)
		v1FeedItemDigest.Title = (*source).Title
		v1FeedItemDigest.ShortDescription = (*source).ShortDescription
		var stringList []string
		if (*source).ImageUrls != nil {
			stringList = make([]string, len((*source).ImageUrls))
			for i := 0; i < len((*source).ImageUrls); i++ {
				stringList[i] = (*source).ImageUrls[i]
			}
		}
		v1FeedItemDigest.ImageUrls = stringList
		v1FeedItemDigest.PublishPlatform = (*source).PublishPlatform
		v1FeedItemDigest.FeedConfigName = (*source).FeedConfigName
		v1FeedItemDigest.FeedAvatarUrl = (*source).FeedAvatarURL
		pV1FeedItemDigest = &v1FeedItemDigest
	}
	return pV1FeedItemDigest
}
func (c *toPBConverterImpl) ToPBFeedItemDigestList(source []*modelyesod.FeedItemDigest) []*v11.FeedItemDigest {
	var pV1FeedItemDigestList []*v11.FeedItemDigest
	if source != nil {
		pV1FeedItemDigestList = make([]*v11.FeedItemDigest, len(source))
		for i := 0; i < len(source); i++ {
			pV1FeedItemDigestList[i] = c.ToPBFeedItemDigest(source[i])
		}
	}
	return pV1FeedItemDigestList
}
func (c *toPBConverterImpl) ToPBFeedItemList(source []*modelfeed.Item) []*v1.FeedItem {
	var pV1FeedItemList []*v1.FeedItem
	if source != nil {
		pV1FeedItemList = make([]*v1.FeedItem, len(source))
		for i := 0; i < len(source); i++ {
			pV1FeedItemList[i] = c.ToPBFeedItem(source[i])
		}
	}
	return pV1FeedItemList
}
func (c *toPBConverterImpl) ToPBFeedWithConfig(source *modelyesod.FeedWithConfig) *v11.ListFeedConfigsResponse_FeedWithConfig {
	var pV1ListFeedConfigsResponse_FeedWithConfig *v11.ListFeedConfigsResponse_FeedWithConfig
	if source != nil {
		var v1ListFeedConfigsResponse_FeedWithConfig v11.ListFeedConfigsResponse_FeedWithConfig
		v1ListFeedConfigsResponse_FeedWithConfig.Feed = c.ToPBFeed((*source).Feed)
		v1ListFeedConfigsResponse_FeedWithConfig.Config = c.ToPBFeedConfig((*source).FeedConfig)
		pV1ListFeedConfigsResponse_FeedWithConfig = &v1ListFeedConfigsResponse_FeedWithConfig
	}
	return pV1ListFeedConfigsResponse_FeedWithConfig
}
func (c *toPBConverterImpl) ToPBFeedWithConfigList(source []*modelyesod.FeedWithConfig) []*v11.ListFeedConfigsResponse_FeedWithConfig {
	var pV1ListFeedConfigsResponse_FeedWithConfigList []*v11.ListFeedConfigsResponse_FeedWithConfig
	if source != nil {
		pV1ListFeedConfigsResponse_FeedWithConfigList = make([]*v11.ListFeedConfigsResponse_FeedWithConfig, len(source))
		for i := 0; i < len(source); i++ {
			pV1ListFeedConfigsResponse_FeedWithConfigList[i] = c.ToPBFeedWithConfig(source[i])
		}
	}
	return pV1ListFeedConfigsResponse_FeedWithConfigList
}
func (c *toPBConverterImpl) ToPBInternalIDList(source []model.InternalID) []*v1.InternalID {
	var pV1InternalIDList []*v1.InternalID
	if source != nil {
		pV1InternalIDList = make([]*v1.InternalID, len(source))
		for i := 0; i < len(source); i++ {
			pV1InternalIDList[i] = ToPBInternalID(source[i])
		}
	}
	return pV1InternalIDList
}
func (c *toPBConverterImpl) ToPBNotifyFlow(source *modelnetzach.NotifyFlow) *v11.NotifyFlow {
	var pV1NotifyFlow *v11.NotifyFlow
	if source != nil {
		var v1NotifyFlow v11.NotifyFlow
		v1NotifyFlow.Id = ToPBInternalID((*source).ID)
		v1NotifyFlow.Name = (*source).Name
		v1NotifyFlow.Description = (*source).Description
		v1NotifyFlow.Source = c.ToPBNotifyFlowSource((*source).Source)
		var pV1NotifyFlowTargetList []*v11.NotifyFlowTarget
		if (*source).Targets != nil {
			pV1NotifyFlowTargetList = make([]*v11.NotifyFlowTarget, len((*source).Targets))
			for i := 0; i < len((*source).Targets); i++ {
				pV1NotifyFlowTargetList[i] = c.ToPBNotifyFlowTarget((*source).Targets[i])
			}
		}
		v1NotifyFlow.Targets = pV1NotifyFlowTargetList
		v1NotifyFlow.Status = ToPBNotifyFlowStatus((*source).Status)
		pV1NotifyFlow = &v1NotifyFlow
	}
	return pV1NotifyFlow
}
func (c *toPBConverterImpl) ToPBNotifyFlowList(source []*modelnetzach.NotifyFlow) []*v11.NotifyFlow {
	var pV1NotifyFlowList []*v11.NotifyFlow
	if source != nil {
		pV1NotifyFlowList = make([]*v11.NotifyFlow, len(source))
		for i := 0; i < len(source); i++ {
			pV1NotifyFlowList[i] = c.ToPBNotifyFlow(source[i])
		}
	}
	return pV1NotifyFlowList
}
func (c *toPBConverterImpl) ToPBNotifyFlowSource(source *modelnetzach.NotifyFlowSource) *v11.NotifyFlowSource {
	var pV1NotifyFlowSource *v11.NotifyFlowSource
	if source != nil {
		var v1NotifyFlowSource v11.NotifyFlowSource
		v1NotifyFlowSource.FeedIdFilter = c.ToPBInternalIDList((*source).FeedIDFilter)
		pV1NotifyFlowSource = &v1NotifyFlowSource
	}
	return pV1NotifyFlowSource
}
func (c *toPBConverterImpl) ToPBNotifyFlowTarget(source *modelnetzach.NotifyFlowTarget) *v11.NotifyFlowTarget {
	var pV1NotifyFlowTarget *v11.NotifyFlowTarget
	if source != nil {
		var v1NotifyFlowTarget v11.NotifyFlowTarget
		v1NotifyFlowTarget.TargetId = ToPBInternalID((*source).TargetID)
		v1NotifyFlowTarget.ChannelId = (*source).ChannelID
		pV1NotifyFlowTarget = &v1NotifyFlowTarget
	}
	return pV1NotifyFlowTarget
}
func (c *toPBConverterImpl) ToPBNotifyTarget(source *modelnetzach.NotifyTarget) *v11.NotifyTarget {
	var pV1NotifyTarget *v11.NotifyTarget
	if source != nil {
		var v1NotifyTarget v11.NotifyTarget
		v1NotifyTarget.Id = ToPBInternalID((*source).ID)
		v1NotifyTarget.Name = (*source).Name
		v1NotifyTarget.Description = (*source).Description
		v1NotifyTarget.Type = ToPBNotifyTargetType((*source).Type)
		v1NotifyTarget.Status = ToPBNotifyTargetStatus((*source).Status)
		v1NotifyTarget.Token = (*source).Token
		pV1NotifyTarget = &v1NotifyTarget
	}
	return pV1NotifyTarget
}
func (c *toPBConverterImpl) ToPBNotifyTargetList(source []*modelnetzach.NotifyTarget) []*v11.NotifyTarget {
	var pV1NotifyTargetList []*v11.NotifyTarget
	if source != nil {
		pV1NotifyTargetList = make([]*v11.NotifyTarget, len(source))
		for i := 0; i < len(source); i++ {
			pV1NotifyTargetList[i] = c.ToPBNotifyTarget(source[i])
		}
	}
	return pV1NotifyTargetList
}
func (c *toPBConverterImpl) ToPBTimeRange(source *model.TimeRange) *v1.TimeRange {
	var pV1TimeRange *v1.TimeRange
	if source != nil {
		var v1TimeRange v1.TimeRange
		v1TimeRange.StartTime = ToPBTime((*source).StartTime)
		v1TimeRange.Duration = ToPBDuration((*source).Duration)
		pV1TimeRange = &v1TimeRange
	}
	return pV1TimeRange
}
func (c *toPBConverterImpl) ToPBUser(source *modeltiphereth.User) *v11.User {
	var pV1User *v11.User
	if source != nil {
		var v1User v11.User
		v1User.Id = ToPBInternalID((*source).ID)
		v1User.Username = (*source).UserName
		v1User.Type = ToPBUserType((*source).Type)
		v1User.Status = ToPBUserStatus((*source).Status)
		pV1User = &v1User
	}
	return pV1User
}
func (c *toPBConverterImpl) ToPBUserList(source []*modeltiphereth.User) []*v11.User {
	var pV1UserList []*v11.User
	if source != nil {
		pV1UserList = make([]*v11.User, len(source))
		for i := 0; i < len(source); i++ {
			pV1UserList[i] = c.ToPBUser(source[i])
		}
	}
	return pV1UserList
}
func (c *toPBConverterImpl) pModelfeedPersonToPV1FeedPerson(source *modelfeed.Person) *v1.FeedPerson {
	var pV1FeedPerson *v1.FeedPerson
	if source != nil {
		var v1FeedPerson v1.FeedPerson
		v1FeedPerson.Name = (*source).Name
		v1FeedPerson.Email = (*source).Email
		pV1FeedPerson = &v1FeedPerson
	}
	return pV1FeedPerson
}
func (c *toPBConverterImpl) pModelgeburaAppDetailsToPV1AppDetails(source *modelgebura.AppDetails) *v1.AppDetails {
	var pV1AppDetails *v1.AppDetails
	if source != nil {
		var v1AppDetails v1.AppDetails
		v1AppDetails.Description = (*source).Description
		v1AppDetails.ReleaseDate = (*source).ReleaseDate
		v1AppDetails.Developer = (*source).Developer
		v1AppDetails.Publisher = (*source).Publisher
		v1AppDetails.Version = (*source).Version
		pV1AppDetails = &v1AppDetails
	}
	return pV1AppDetails
}
