package libapp

import "github.com/tuihub/librarian/internal/lib/libzap"

type InherentSettings struct {
	EnablePanicRecovery bool
	LogLevel            libzap.Level
	DefaultConfPath     string
}
