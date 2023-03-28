//go:build debug

package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

func GetInherentSettings() Settings {
	return Settings{
		EnablePanicRecovery: false,
		LogLevel:            libzap.DebugLevel,
	}
}
