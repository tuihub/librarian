package cmd

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libzap"

	"github.com/urfave/cli/v2"
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
		Subcommands: []*cli.Command{
			{
				Name:  "user",
				Usage: "User management commands",
				Subcommands: []*cli.Command{
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

func runCmdAdminCreateUser(ctx *cli.Context) error {
	stdLogger := libzap.NewStdout(libzap.InfoLevel).Sugar()
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

	var bc conf.Librarian
	err = appSettings.LoadConfig(&bc)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	digests := genConfigDigest(&bc)
	app, cleanup, err := wireAdmin(
		digests,
		bc.GetEnableServiceDiscovery(),
		bc.GetServer(),
		bc.GetDatabase(),
		bc.GetS3(),
		bc.GetPorter(),
		bc.GetMiner().GetData(),
		bc.GetAuth(),
		bc.GetMq(),
		bc.GetCache(),
		bc.GetConsul(),
		bc.GetSearch(),
		appSettings,
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}
	defer cleanup()
	username := ctx.String(cmdAdminFlagUsername)
	password := ctx.String(cmdAdminFlagPassword)
	isAdmin := ctx.Bool(cmdAdminFlagAdmin)
	if username == "" || password == "" {
		stdLogger.Fatalf("Username and password are required")
	}
	err = app.CliCreateUser(ctx.Context, username, password, isAdmin)
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
