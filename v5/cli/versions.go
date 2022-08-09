package cli

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/eopenio/idump/v5/log"
)

var (
	// ReleaseVersion is the current program version.
	ReleaseVersion = "Unknown"
	// BuildTimestamp is the UTC date time when the program is compiled.
	BuildTimestamp = "Unknown"
	// GitHash is the git commit hash when the program is compiled.
	GitHash = "Unknown"
	// GitBranch is the active git branch when the program is compiled.
	GitBranch = "Unknown"
	// GoVersion is the Go compiler version used to compile this program.
	GoVersion = "Unknown"
)

// LongVersion returns the version information of this program as a string.
func LongVersion() string {
	return fmt.Sprintf(
		"Release version: %s\n"+
			"Git commit hash: %s\n"+
			"Git branch:      %s\n"+
			"Build timestamp: %sZ\n"+
			"Go version:      %s\n",
		ReleaseVersion,
		GitHash,
		GitBranch,
		BuildTimestamp,
		GoVersion,
	)
}

// LogLongVersion logs the version information of this program to the logger.
func LogLongVersion(logger log.Logger) {
	logger.Info("Welcome to dumpling",
		zap.String("Release Version", ReleaseVersion),
		zap.String("Git Commit Hash", GitHash),
		zap.String("Git Branch", GitBranch),
		zap.String("Build timestamp", BuildTimestamp),
		zap.String("Go Version", GoVersion))
}
