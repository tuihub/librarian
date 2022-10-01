package biztiphereth

import librarian "github.com/tuihub/protos/pkg/librarian/v1"

func ToLibrarianAccountPlatform(p AccountPlatform) librarian.AccountPlatform {
	switch p {
	case AccountPlatformSteam:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM
	default:
		return librarian.AccountPlatform_ACCOUNT_PLATFORM_UNSPECIFIED
	}
}
