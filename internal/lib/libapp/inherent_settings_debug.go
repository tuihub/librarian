//go:build debug

package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

func getInherentSettings() InherentSettings {
	return InherentSettings{
		EnablePanicRecovery: false,
		LogLevel:            libzap.DebugLevel,
		DefaultConfPath:     "./configs/config.toml",
	}
}
