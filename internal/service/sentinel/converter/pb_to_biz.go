package converter

import (
	"github.com/tuihub/librarian/internal/model/modelgebura"
	sentinel "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/internal/service/sentinel/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend PtrToString
type toBizConverter interface { //nolint:unused // used by generator
	// goverter:ignore ID
	// goverter:map UrlAlternatives AlternativeUrls
	ToBizSentinelInfo(*sentinel.ReportSentinelInformationRequest) *modelgebura.SentinelInfo
	// goverter:map Id ReportedID
	// goverter:ignore AppBinaries
	ToBizSentinelLibrary(*sentinel.SentinelLibrary) *modelgebura.SentinelLibrary
	// goverter:map SentinelGeneratedId GeneratedID
	ToBizSentinelAppBinary(*sentinel.SentinelLibraryAppBinary) *modelgebura.SentinelAppBinary
	ToBizSentinelAppBinaryList([]*sentinel.SentinelLibraryAppBinary) []*modelgebura.SentinelAppBinary
	ToBizSentinelAppBinaryFile(*sentinel.SentinelLibraryAppBinaryFile) *modelgebura.SentinelAppBinaryFile
}

func PtrToString(u *string) string {
	if u == nil {
		return ""
	}
	return *u
}
