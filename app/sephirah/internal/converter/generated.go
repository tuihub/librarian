// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package converter

import (
	bizgebura "github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	biztiphereth "github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	bizyesod "github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	libauth "github.com/tuihub/librarian/internal/lib/libauth"
	v11 "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	v1 "github.com/tuihub/protos/pkg/librarian/v1"
)

type toBizConverterImpl struct{}

func (c *toBizConverterImpl) ToBizApp(source *v1.App) *bizgebura.App {
	var pBizgeburaApp *bizgebura.App
	if source != nil {
		var bizgeburaApp bizgebura.App
		bizgeburaApp.InternalID = ToBizInternalID((*source).Id)
		bizgeburaApp.Source = ToBizAppSource((*source).Source)
		bizgeburaApp.SourceAppID = (*source).SourceAppId
		bizgeburaApp.SourceURL = PtrToString((*source).SourceUrl)
		bizgeburaApp.Name = (*source).Name
		bizgeburaApp.Type = ToBizAppType((*source).Type)
		bizgeburaApp.ShortDescription = (*source).ShortDescription
		bizgeburaApp.ImageURL = (*source).ImageUrl
		bizgeburaApp.Details = c.pV1AppDetailsToPBizgeburaAppDetails((*source).Details)
		pBizgeburaApp = &bizgeburaApp
	}
	return pBizgeburaApp
}
func (c *toBizConverterImpl) ToBizAppPackage(source *v1.AppPackage) *bizgebura.AppPackage {
	var pBizgeburaAppPackage *bizgebura.AppPackage
	if source != nil {
		var bizgeburaAppPackage bizgebura.AppPackage
		bizgeburaAppPackage.InternalID = ToBizInternalID((*source).Id)
		bizgeburaAppPackage.Source = ToBizAppPackageSource((*source).Source)
		bizgeburaAppPackage.SourceID = ToBizInternalID((*source).SourceId)
		bizgeburaAppPackage.SourcePackageID = (*source).SourcePackageId
		bizgeburaAppPackage.Name = (*source).Name
		bizgeburaAppPackage.Description = (*source).Description
		bizgeburaAppPackage.Binary = c.ToBizAppPackageBinary((*source).Binary)
		pBizgeburaAppPackage = &bizgeburaAppPackage
	}
	return pBizgeburaAppPackage
}
func (c *toBizConverterImpl) ToBizAppPackageBinary(source *v1.AppPackageBinary) *bizgebura.AppPackageBinary {
	var pBizgeburaAppPackageBinary *bizgebura.AppPackageBinary
	if source != nil {
		var bizgeburaAppPackageBinary bizgebura.AppPackageBinary
		bizgeburaAppPackageBinary.Name = (*source).Name
		bizgeburaAppPackageBinary.Size = (*source).Size
		bizgeburaAppPackageBinary.PublicURL = (*source).PublicUrl
		pBizgeburaAppPackageBinary = &bizgeburaAppPackageBinary
	}
	return pBizgeburaAppPackageBinary
}
func (c *toBizConverterImpl) ToBizAppPackageSourceList(source []v1.AppPackageSource) []bizgebura.AppPackageSource {
	var bizgeburaAppPackageSourceList []bizgebura.AppPackageSource
	if source != nil {
		bizgeburaAppPackageSourceList = make([]bizgebura.AppPackageSource, len(source))
		for i := 0; i < len(source); i++ {
			bizgeburaAppPackageSourceList[i] = ToBizAppPackageSource(source[i])
		}
	}
	return bizgeburaAppPackageSourceList
}
func (c *toBizConverterImpl) ToBizAppSourceList(source []v1.AppSource) []bizgebura.AppSource {
	var bizgeburaAppSourceList []bizgebura.AppSource
	if source != nil {
		bizgeburaAppSourceList = make([]bizgebura.AppSource, len(source))
		for i := 0; i < len(source); i++ {
			bizgeburaAppSourceList[i] = ToBizAppSource(source[i])
		}
	}
	return bizgeburaAppSourceList
}
func (c *toBizConverterImpl) ToBizAppTypeList(source []v1.AppType) []bizgebura.AppType {
	var bizgeburaAppTypeList []bizgebura.AppType
	if source != nil {
		bizgeburaAppTypeList = make([]bizgebura.AppType, len(source))
		for i := 0; i < len(source); i++ {
			bizgeburaAppTypeList[i] = ToBizAppType(source[i])
		}
	}
	return bizgeburaAppTypeList
}
func (c *toBizConverterImpl) ToBizFeedConfig(source *v11.FeedConfig) *bizyesod.FeedConfig {
	var pBizyesodFeedConfig *bizyesod.FeedConfig
	if source != nil {
		var bizyesodFeedConfig bizyesod.FeedConfig
		bizyesodFeedConfig.InternalID = ToBizInternalID((*source).Id)
		bizyesodFeedConfig.FeedURL = (*source).FeedUrl
		bizyesodFeedConfig.AuthorAccount = ToBizInternalID((*source).AuthorAccount)
		bizyesodFeedConfig.Source = ToBizFeedConfigSource((*source).Source)
		bizyesodFeedConfig.Status = ToBizFeedConfigStatus((*source).Status)
		bizyesodFeedConfig.PullInterval = DurationPBToDuration((*source).PullInterval)
		pBizyesodFeedConfig = &bizyesodFeedConfig
	}
	return pBizyesodFeedConfig
}
func (c *toBizConverterImpl) ToBizInternalIDList(source []*v1.InternalID) []int64 {
	var int64List []int64
	if source != nil {
		int64List = make([]int64, len(source))
		for i := 0; i < len(source); i++ {
			int64List[i] = ToBizInternalID(source[i])
		}
	}
	return int64List
}
func (c *toBizConverterImpl) ToBizUser(source *v11.User) *biztiphereth.User {
	var pBiztipherethUser *biztiphereth.User
	if source != nil {
		var biztipherethUser biztiphereth.User
		biztipherethUser.InternalID = ToBizInternalID((*source).Id)
		biztipherethUser.UserName = (*source).Username
		biztipherethUser.PassWord = (*source).Password
		biztipherethUser.Type = ToLibAuthUserType((*source).Type)
		biztipherethUser.Status = ToBizUserStatus((*source).Status)
		pBiztipherethUser = &biztipherethUser
	}
	return pBiztipherethUser
}
func (c *toBizConverterImpl) ToBizUserStatusList(source []v11.UserStatus) []biztiphereth.UserStatus {
	var biztipherethUserStatusList []biztiphereth.UserStatus
	if source != nil {
		biztipherethUserStatusList = make([]biztiphereth.UserStatus, len(source))
		for i := 0; i < len(source); i++ {
			biztipherethUserStatusList[i] = ToBizUserStatus(source[i])
		}
	}
	return biztipherethUserStatusList
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
func (c *toBizConverterImpl) pV1AppDetailsToPBizgeburaAppDetails(source *v1.AppDetails) *bizgebura.AppDetails {
	var pBizgeburaAppDetails *bizgebura.AppDetails
	if source != nil {
		var bizgeburaAppDetails bizgebura.AppDetails
		bizgeburaAppDetails.Description = (*source).Description
		bizgeburaAppDetails.ReleaseDate = (*source).ReleaseDate
		bizgeburaAppDetails.Developer = (*source).Developer
		bizgeburaAppDetails.Publisher = (*source).Publisher
		bizgeburaAppDetails.Version = (*source).Version
		pBizgeburaAppDetails = &bizgeburaAppDetails
	}
	return pBizgeburaAppDetails
}

type toPBConverterImpl struct{}

func (c *toPBConverterImpl) ToPBAccount(source biztiphereth.Account) v1.Account {
	var v1Account v1.Account
	v1Account.Id = c.biztipherethAccountToPV1InternalID(source)
	v1Account.Platform = ToPBAccountPlatform(source.Platform)
	v1Account.PlatformAccountId = source.PlatformAccountID
	v1Account.Name = source.Name
	v1Account.ProfileUrl = source.ProfileURL
	v1Account.AvatarUrl = source.AvatarURL
	return v1Account
}
func (c *toPBConverterImpl) ToPBAccountInternalID(source biztiphereth.Account) v1.InternalID {
	var v1InternalID v1.InternalID
	v1InternalID.Id = source.InternalID
	return v1InternalID
}
func (c *toPBConverterImpl) ToPBAccountList(source []*biztiphereth.Account) []*v1.Account {
	var pV1AccountList []*v1.Account
	if source != nil {
		pV1AccountList = make([]*v1.Account, len(source))
		for i := 0; i < len(source); i++ {
			pV1AccountList[i] = c.pBiztipherethAccountToPV1Account(source[i])
		}
	}
	return pV1AccountList
}
func (c *toPBConverterImpl) ToPBApp(source bizgebura.App) v1.App {
	var v1App v1.App
	v1App.Id = c.bizgeburaAppToPV1InternalID(source)
	v1App.Source = ToPBAppSource(source.Source)
	v1App.SourceAppId = source.SourceAppID
	pString := source.SourceURL
	v1App.SourceUrl = &pString
	v1App.Name = source.Name
	v1App.Type = ToPBAppType(source.Type)
	v1App.ShortDescription = source.ShortDescription
	v1App.ImageUrl = source.ImageURL
	v1App.Details = c.pBizgeburaAppDetailsToPV1AppDetails(source.Details)
	return v1App
}
func (c *toPBConverterImpl) ToPBAppInternalID(source bizgebura.App) v1.InternalID {
	var v1InternalID v1.InternalID
	v1InternalID.Id = source.InternalID
	return v1InternalID
}
func (c *toPBConverterImpl) ToPBAppList(source []*bizgebura.App) []*v1.App {
	var pV1AppList []*v1.App
	if source != nil {
		pV1AppList = make([]*v1.App, len(source))
		for i := 0; i < len(source); i++ {
			pV1AppList[i] = c.pBizgeburaAppToPV1App(source[i])
		}
	}
	return pV1AppList
}
func (c *toPBConverterImpl) ToPBAppPackage(source *bizgebura.AppPackage) *v1.AppPackage {
	var pV1AppPackage *v1.AppPackage
	if source != nil {
		var v1AppPackage v1.AppPackage
		v1AppPackage.Id = c.bizgeburaAppPackageToPV1InternalID((*source))
		v1AppPackage.Source = ToPBAppPackageSource((*source).Source)
		v1AppPackage.SourceId = c.bizgeburaAppPackageToPV1InternalID((*source))
		v1AppPackage.SourcePackageId = (*source).SourcePackageID
		v1AppPackage.Name = (*source).Name
		v1AppPackage.Description = (*source).Description
		v1AppPackage.Binary = c.ToPBAppPackageBinary((*source).Binary)
		pV1AppPackage = &v1AppPackage
	}
	return pV1AppPackage
}
func (c *toPBConverterImpl) ToPBAppPackageBinary(source *bizgebura.AppPackageBinary) *v1.AppPackageBinary {
	var pV1AppPackageBinary *v1.AppPackageBinary
	if source != nil {
		var v1AppPackageBinary v1.AppPackageBinary
		v1AppPackageBinary.Name = (*source).Name
		v1AppPackageBinary.Size = (*source).Size
		v1AppPackageBinary.PublicUrl = (*source).PublicURL
		pV1AppPackageBinary = &v1AppPackageBinary
	}
	return pV1AppPackageBinary
}
func (c *toPBConverterImpl) ToPBAppPackageInternalID(source bizgebura.AppPackage) v1.InternalID {
	var v1InternalID v1.InternalID
	v1InternalID.Id = source.InternalID
	return v1InternalID
}
func (c *toPBConverterImpl) ToPBAppPackageList(source []*bizgebura.AppPackage) []*v1.AppPackage {
	var pV1AppPackageList []*v1.AppPackage
	if source != nil {
		pV1AppPackageList = make([]*v1.AppPackage, len(source))
		for i := 0; i < len(source); i++ {
			pV1AppPackageList[i] = c.ToPBAppPackage(source[i])
		}
	}
	return pV1AppPackageList
}
func (c *toPBConverterImpl) ToPBUser(source biztiphereth.User) v11.User {
	var v1User v11.User
	v1User.Id = c.biztipherethUserToPV1InternalID(source)
	v1User.Username = source.UserName
	v1User.Type = ToPBUserType(source.Type)
	v1User.Status = ToPBUserStatus(source.Status)
	return v1User
}
func (c *toPBConverterImpl) ToPBUserInternalID(source biztiphereth.User) v1.InternalID {
	var v1InternalID v1.InternalID
	v1InternalID.Id = source.InternalID
	return v1InternalID
}
func (c *toPBConverterImpl) ToPBUserList(source []*biztiphereth.User) []*v11.User {
	var pV1UserList []*v11.User
	if source != nil {
		pV1UserList = make([]*v11.User, len(source))
		for i := 0; i < len(source); i++ {
			pV1UserList[i] = c.pBiztipherethUserToPV1User(source[i])
		}
	}
	return pV1UserList
}
func (c *toPBConverterImpl) bizgeburaAppPackageToPV1InternalID(source bizgebura.AppPackage) *v1.InternalID {
	v1InternalID := c.ToPBAppPackageInternalID(source)
	return &v1InternalID
}
func (c *toPBConverterImpl) bizgeburaAppToPV1InternalID(source bizgebura.App) *v1.InternalID {
	v1InternalID := c.ToPBAppInternalID(source)
	return &v1InternalID
}
func (c *toPBConverterImpl) biztipherethAccountToPV1InternalID(source biztiphereth.Account) *v1.InternalID {
	v1InternalID := c.ToPBAccountInternalID(source)
	return &v1InternalID
}
func (c *toPBConverterImpl) biztipherethUserToPV1InternalID(source biztiphereth.User) *v1.InternalID {
	v1InternalID := c.ToPBUserInternalID(source)
	return &v1InternalID
}
func (c *toPBConverterImpl) pBizgeburaAppDetailsToPV1AppDetails(source *bizgebura.AppDetails) *v1.AppDetails {
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
func (c *toPBConverterImpl) pBizgeburaAppToPV1App(source *bizgebura.App) *v1.App {
	var pV1App *v1.App
	if source != nil {
		v1App := c.ToPBApp((*source))
		pV1App = &v1App
	}
	return pV1App
}
func (c *toPBConverterImpl) pBiztipherethAccountToPV1Account(source *biztiphereth.Account) *v1.Account {
	var pV1Account *v1.Account
	if source != nil {
		v1Account := c.ToPBAccount((*source))
		pV1Account = &v1Account
	}
	return pV1Account
}
func (c *toPBConverterImpl) pBiztipherethUserToPV1User(source *biztiphereth.User) *v11.User {
	var pV1User *v11.User
	if source != nil {
		v1User := c.ToPBUser((*source))
		pV1User = &v1User
	}
	return pV1User
}