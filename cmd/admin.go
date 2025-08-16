package cmd

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/urfave/cli/v3"
)

const (
	cmdAdminFlagUsername = "username"
	cmdAdminFlagPassword = "password"
	cmdAdminFlagAdmin    = "admin"
)

func newCmdAdmin() *cli.Command {
	return &cli.Command{
		Name:  "admin",
		Usage: "Administrative commands",
		Commands: []*cli.Command{
			{
				Name:  "user",
				Usage: "User management commands",
				Commands: []*cli.Command{
					{
						Name:        "create",
						Usage:       "Create a new user",
						Description: "Create a new user with specified credentials",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     cmdAdminFlagUsername,
								Aliases:  []string{"u"},
								Required: true,
								Usage:    "Username for the new user",
							},
							&cli.StringFlag{
								Name:     cmdAdminFlagPassword,
								Aliases:  []string{"p"},
								Required: true,
								Usage:    "Password for the new user",
							},
							&cli.BoolFlag{
								Name: cmdAdminFlagAdmin,
							},
						},
						Action: runCmdAdminCreateUser,
					},
				},
			},
		},
	}
}

func runCmdAdminCreateUser(ctx context.Context, cmd *cli.Command) error {
	stdLogger := libzap.NewStdout(libzap.InfoLevel).Sugar()
	appSettings, err := libapp.NewAppSettings(
		id,
		name,
		version,
		protoVersion,
		date,
		cmd.String(cmdServeFlagConfig),
		cmd.String(cmdServeFlagData),
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
	app, cleanup, err := wireAdmin(
		digests,
		bc,
		appSettings,
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}
	defer cleanup()
	username := cmd.String(cmdAdminFlagUsername)
	password := cmd.String(cmdAdminFlagPassword)
	isAdmin := cmd.Bool(cmdAdminFlagAdmin)
	if username == "" || password == "" {
		stdLogger.Fatalf("Username and password are required")
	}
	err = app.CliCreateUser(ctx, username, password, isAdmin)
	if err != nil {
		stdLogger.Fatalf("Create user failed: %v", err)
	}
	if isAdmin {
		stdLogger.Infof("User %s created successfully with admin privileges", username)
	} else {
		stdLogger.Infof("User %s created successfully", username)
	}
	return nil
}
