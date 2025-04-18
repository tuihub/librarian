package cmd

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/urfave/cli/v2"
)

func newCmdConfig() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configuration commands",
		Subcommands: []*cli.Command{
			{
				Name:        "check",
				Usage:       "Validate configuration file",
				Description: "Check if the configuration file is valid",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "Path to the configuration file",
						Value:   "config.toml",
					},
				},
				Action: runCmdConfigCheck,
			},
		},
	}
}

func runCmdConfigCheck(ctx *cli.Context) error {
	stdLogger := libzap.NewStdout(libzap.InfoLevel).Sugar()
	stdLogger.Infof("=== Configuring ===")
	stdLogger.Infof("[Service\t] Name: %s", name)
	stdLogger.Infof("[Service\t] Version: %s", version)
	appSettings, err := libapp.NewAppSettings(
		id,
		name,
		version,
		protoVersion,
		date,
		ctx.String(cmdServeFlagConfig),
		ctx.String(cmdServeFlagData),
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}

	bc, err := conf.Load(appSettings.ConfPath)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	bc, err = conf.ApplyDeployMode(bc, stdLogger)
	if err != nil {
		stdLogger.Fatalf("Apply deploy mode failed: %v", err)
	}
	digests := conf.GenConfigDigest(bc)
	logConfigDigest(digests, stdLogger)
	stdLogger.Infof("=== Configuration Check Completed ===")
	return nil
}
