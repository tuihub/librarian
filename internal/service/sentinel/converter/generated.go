// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package converter

import (
	modelgebura "github.com/tuihub/librarian/internal/model/modelgebura"
	sentinel "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"
)

func ToBizSentinel(source *sentinel.ReportSentinelInformationRequest) *modelgebura.Sentinel {
	var pModelgeburaSentinel *modelgebura.Sentinel
	if source != nil {
		var modelgeburaSentinel modelgebura.Sentinel
		modelgeburaSentinel.URL = (*source).Url
		if (*source).UrlAlternatives != nil {
			modelgeburaSentinel.AlternativeUrls = make([]string, len((*source).UrlAlternatives))
			for i := 0; i < len((*source).UrlAlternatives); i++ {
				modelgeburaSentinel.AlternativeUrls[i] = (*source).UrlAlternatives[i]
			}
		}
		modelgeburaSentinel.GetTokenPath = (*source).GetTokenPath
		modelgeburaSentinel.DownloadFileBasePath = (*source).DownloadFileBasePath
		if (*source).Libraries != nil {
			modelgeburaSentinel.Libraries = make([]*modelgebura.SentinelLibrary, len((*source).Libraries))
			for j := 0; j < len((*source).Libraries); j++ {
				modelgeburaSentinel.Libraries[j] = ToBizSentinelLibrary((*source).Libraries[j])
			}
		}
		pModelgeburaSentinel = &modelgeburaSentinel
	}
	return pModelgeburaSentinel
}
func ToBizSentinelAppBinary(source *sentinel.SentinelLibraryAppBinary) *modelgebura.SentinelAppBinary {
	var pModelgeburaSentinelAppBinary *modelgebura.SentinelAppBinary
	if source != nil {
		var modelgeburaSentinelAppBinary modelgebura.SentinelAppBinary
		modelgeburaSentinelAppBinary.SentinelLibraryID = (*source).SentinelLibraryId
		modelgeburaSentinelAppBinary.GeneratedID = (*source).SentinelGeneratedId
		modelgeburaSentinelAppBinary.SizeBytes = (*source).SizeBytes
		modelgeburaSentinelAppBinary.NeedToken = (*source).NeedToken
		if (*source).Files != nil {
			modelgeburaSentinelAppBinary.Files = make([]*modelgebura.SentinelAppBinaryFile, len((*source).Files))
			for i := 0; i < len((*source).Files); i++ {
				modelgeburaSentinelAppBinary.Files[i] = ToBizSentinelAppBinaryFile((*source).Files[i])
			}
		}
		modelgeburaSentinelAppBinary.Name = (*source).Name
		modelgeburaSentinelAppBinary.Version = (*source).Version
		modelgeburaSentinelAppBinary.Developer = (*source).Developer
		modelgeburaSentinelAppBinary.Publisher = (*source).Publisher
		pModelgeburaSentinelAppBinary = &modelgeburaSentinelAppBinary
	}
	return pModelgeburaSentinelAppBinary
}
func ToBizSentinelAppBinaryFile(source *sentinel.SentinelLibraryAppBinaryFile) *modelgebura.SentinelAppBinaryFile {
	var pModelgeburaSentinelAppBinaryFile *modelgebura.SentinelAppBinaryFile
	if source != nil {
		var modelgeburaSentinelAppBinaryFile modelgebura.SentinelAppBinaryFile
		modelgeburaSentinelAppBinaryFile.Name = (*source).Name
		modelgeburaSentinelAppBinaryFile.SizeBytes = (*source).SizeBytes
		if (*source).Sha256 != nil {
			modelgeburaSentinelAppBinaryFile.Sha256 = make([]uint8, len((*source).Sha256))
			for i := 0; i < len((*source).Sha256); i++ {
				modelgeburaSentinelAppBinaryFile.Sha256[i] = (*source).Sha256[i]
			}
		}
		modelgeburaSentinelAppBinaryFile.ServerFilePath = (*source).ServerFilePath
		modelgeburaSentinelAppBinaryFile.ChunksInfo = PtrToString((*source).ChunksInfo)
		pModelgeburaSentinelAppBinaryFile = &modelgeburaSentinelAppBinaryFile
	}
	return pModelgeburaSentinelAppBinaryFile
}
func ToBizSentinelAppBinaryList(source []*sentinel.SentinelLibraryAppBinary) []*modelgebura.SentinelAppBinary {
	var pModelgeburaSentinelAppBinaryList []*modelgebura.SentinelAppBinary
	if source != nil {
		pModelgeburaSentinelAppBinaryList = make([]*modelgebura.SentinelAppBinary, len(source))
		for i := 0; i < len(source); i++ {
			pModelgeburaSentinelAppBinaryList[i] = ToBizSentinelAppBinary(source[i])
		}
	}
	return pModelgeburaSentinelAppBinaryList
}
func ToBizSentinelLibrary(source *sentinel.SentinelLibrary) *modelgebura.SentinelLibrary {
	var pModelgeburaSentinelLibrary *modelgebura.SentinelLibrary
	if source != nil {
		var modelgeburaSentinelLibrary modelgebura.SentinelLibrary
		modelgeburaSentinelLibrary.ReportedID = (*source).Id
		modelgeburaSentinelLibrary.DownloadBasePath = (*source).DownloadBasePath
		pModelgeburaSentinelLibrary = &modelgeburaSentinelLibrary
	}
	return pModelgeburaSentinelLibrary
}
