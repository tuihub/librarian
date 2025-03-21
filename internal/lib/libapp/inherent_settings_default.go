//go:build !debug && !release

package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

func getInherentSettings() InherentSettings {
	return InherentSettings{
		EnablePanicRecovery: true,
		LogLevel:            libzap.InfoLevel,
		DefaultConfPath:     "./configs/config.yaml",
	}
}
