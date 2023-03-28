//go:build release

package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

func GetInherentSettings() Settings {
	return Settings{
		EnablePanicRecovery: true,
		LogLevel:            libzap.ErrorLevel,
	}
}
