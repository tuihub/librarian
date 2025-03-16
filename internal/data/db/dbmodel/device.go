package dbmodel

type SystemType string

const (
	SystemTypeUnspecified SystemType = "unspecified"
	SystemTypeWindows     SystemType = "windows"
	SystemTypeMacOS       SystemType = "macos"
	SystemTypeLinux       SystemType = "linux"
	SystemTypeWeb         SystemType = "web"
	SystemTypeAndroid     SystemType = "android"
	SystemTypeIOS         SystemType = "ios"
)

type Device struct {
	Model
	UserID                  ID
	ClientLocalID           *ID
	Name                    string
	SystemType              SystemType
	SystemVersion           string
	ClientName              string
	ClientSourceCodeAddress string
	ClientVersion           string
}
