package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

type Settings struct {
	EnablePanicRecovery bool
	LogLevel            libzap.Level
}
