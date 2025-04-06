//go:build release

package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

func getInherentSettings() InherentSettings {
	return InherentSettings{
		EnablePanicRecovery: true,
		LogLevel:            libzap.ErrorLevel,
		DefaultConfPath:     "",
		BuildType:           BuildTypeRelease,
	}
}
