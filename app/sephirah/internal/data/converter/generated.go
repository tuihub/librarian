// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package converter

import (
	bizgebura "github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	biztiphereth "github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	ent "github.com/tuihub/librarian/app/sephirah/internal/ent"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ToBizApp(source *ent.App) *bizgebura.App {
	var pBizgeburaApp *bizgebura.App
	if source != nil {
		var bizgeburaApp bizgebura.App
		bizgeburaApp.InternalID = (*source).InternalID
		bizgeburaApp.Source = ToBizAppSource((*source).Source)
		bizgeburaApp.SourceAppID = (*source).SourceAppID
		bizgeburaApp.SourceURL = (*source).SourceURL
		bizgeburaApp.Name = (*source).Name
		bizgeburaApp.Type = ToBizAppType((*source).Type)
		bizgeburaApp.ShortDescription = (*source).ShortDescription
		bizgeburaApp.ImageURL = (*source).ImageURL
		bizgeburaApp.Details = c.entAppToPBizgeburaAppDetails((*source))
		pBizgeburaApp = &bizgeburaApp
	}
	return pBizgeburaApp
}
func (c *ConverterImpl) ToBizAppPacakgeBinary(source ent.AppPackage) bizgebura.AppPackageBinary {
	var bizgeburaAppPackageBinary bizgebura.AppPackageBinary
	bizgeburaAppPackageBinary.Name = source.BinaryName
	bizgeburaAppPackageBinary.Size = source.BinarySize
	bizgeburaAppPackageBinary.PublicURL = source.BinaryPublicURL
	return bizgeburaAppPackageBinary
}
func (c *ConverterImpl) ToBizAppPackage(source *ent.AppPackage) *bizgebura.AppPackage {
	var pBizgeburaAppPackage *bizgebura.AppPackage
	if source != nil {
		var bizgeburaAppPackage bizgebura.AppPackage
		bizgeburaAppPackage.InternalID = (*source).InternalID
		bizgeburaAppPackage.Source = ToBizAppPackageSource((*source).Source)
		bizgeburaAppPackage.SourceID = (*source).SourceID
		bizgeburaAppPackage.SourcePackageID = (*source).SourcePackageID
		bizgeburaAppPackage.Name = (*source).Name
		bizgeburaAppPackage.Description = (*source).Description
		bizgeburaAppPackage.Binary = c.entAppPackageToPBizgeburaAppPackageBinary((*source))
		pBizgeburaAppPackage = &bizgeburaAppPackage
	}
	return pBizgeburaAppPackage
}
func (c *ConverterImpl) ToBizAppPackageList(source []*ent.AppPackage) []*bizgebura.AppPackage {
	var pBizgeburaAppPackageList []*bizgebura.AppPackage
	if source != nil {
		pBizgeburaAppPackageList = make([]*bizgebura.AppPackage, len(source))
		for i := 0; i < len(source); i++ {
			pBizgeburaAppPackageList[i] = c.ToBizAppPackage(source[i])
		}
	}
	return pBizgeburaAppPackageList
}
func (c *ConverterImpl) ToBizUser(source *ent.User) *biztiphereth.User {
	var pBiztipherethUser *biztiphereth.User
	if source != nil {
		var biztipherethUser biztiphereth.User
		biztipherethUser.InternalID = (*source).InternalID
		biztipherethUser.UserName = (*source).Username
		biztipherethUser.PassWord = (*source).Password
		biztipherethUser.Type = ToLibAuthUserType((*source).Type)
		biztipherethUser.Status = ToBizUserStatus((*source).Status)
		pBiztipherethUser = &biztipherethUser
	}
	return pBiztipherethUser
}
func (c *ConverterImpl) entAppPackageToPBizgeburaAppPackageBinary(source ent.AppPackage) *bizgebura.AppPackageBinary {
	bizgeburaAppPackageBinary := c.ToBizAppPacakgeBinary(source)
	return &bizgeburaAppPackageBinary
}
func (c *ConverterImpl) entAppToBizgeburaAppDetails(source ent.App) bizgebura.AppDetails {
	var bizgeburaAppDetails bizgebura.AppDetails
	bizgeburaAppDetails.Description = source.Description
	bizgeburaAppDetails.ReleaseDate = source.ReleaseDate
	bizgeburaAppDetails.Developer = source.Developer
	bizgeburaAppDetails.Publisher = source.Publisher
	bizgeburaAppDetails.Version = source.Version
	return bizgeburaAppDetails
}
func (c *ConverterImpl) entAppToPBizgeburaAppDetails(source ent.App) *bizgebura.AppDetails {
	bizgeburaAppDetails := c.entAppToBizgeburaAppDetails(source)
	return &bizgeburaAppDetails
}
