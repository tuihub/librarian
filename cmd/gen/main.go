package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var models = []interface{}{ //nolint:gochecknoglobals // required by gen
	&model.User{},
	&model.Account{},
	&model.Session{},
	&model.Device{},
	&model.Tag{},
	&model.KV{},

	&modelgebura.App{},
	&modelgebura.AppCategory{},
	&modelgebura.AppAppCategory{},
	&modelgebura.AppInfo{},
	&modelgebura.AppRunTime{},
	&modelgebura.Sentinel{},
	&modelgebura.SentinelAppBinary{},
	&modelgebura.SentinelAppBinaryFile{},
	&modelgebura.SentinelLibrary{},
	&modelgebura.SentinelSession{},
	&modelgebura.StoreApp{},
	&modelgebura.StoreAppBinary{},

	&modelfeed.Feed{},
	&modelfeed.Item{},

	&modelyesod.FeedConfig{},
	&modelyesod.FeedActionSet{},
	&modelyesod.FeedConfigAction{},
	&modelyesod.FeedItemCollection{},
	// &modelyesod.FeedItemCollectionFeedItem{}, // Assuming this is needed if we want to query the join table directly, but GORM usually handles many2many. If we added it to struct, include it.
	// I added FeedItemCollectionFeedItem in modelyesod.go?
	// Yes: type FeedItemCollectionFeedItem struct
	&modelyesod.FeedItemCollectionFeedItem{},

	&modelnetzach.NotifyFlow{},
	&modelnetzach.NotifyFlowSource{},
	&modelnetzach.NotifyFlowTarget{},
	&modelnetzach.NotifySource{},
	&modelnetzach.NotifyTarget{},
	&modelnetzach.SystemNotification{},

	&modelsupervisor.PorterInstance{},
	&modelsupervisor.PorterContext{},

	&modelchesed.Image{},
	&modelchesed.File{},
}

func main() {
	outPath := flag.String("out", "./internal/data/orm/query", "output path for generated query code")
	flag.Parse()

	// 1. Generate Query Code
	g := gen.NewGenerator(gen.Config{
		OutPath:       *outPath,
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.ApplyBasic(models...)
	g.Execute()

	// 2. Generate SQL Schema for SQLite
	generateSchema("sqlite", "file::memory:?cache=shared", "schema/sqlite/schema.sql")

	// 3. Generate SQL Schema for Postgres
	// Use a dummy DSN for DryRun
	generateSchema(
		"postgres",
		"host=localhost user=user password=pass dbname=db port=5432 sslmode=disable TimeZone=UTC",
		"schema/postgres/schema.sql",
	)
}

type SQLLogger struct {
	SQLs []string
}

func (l *SQLLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *SQLLogger) Info(context.Context, string, ...interface{})  {}
func (l *SQLLogger) Warn(context.Context, string, ...interface{})  {}
func (l *SQLLogger) Error(context.Context, string, ...interface{}) {}
func (l *SQLLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	if sql != "" {
		l.SQLs = append(l.SQLs, sql+";")
	}
}

func generateSchema(driverName string, dsn string, outputPath string) {
	sqlLogger := &SQLLogger{}

	var dialector gorm.Dialector
	switch driverName {
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		log.Fatalf("unsupported driver: %s", driverName)
	}

	// Use DryRun mode to capture SQL
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:                                   sqlLogger,
		DryRun:                                   true,
		DisableForeignKeyConstraintWhenMigrating: true, // Optional: avoid FK issues during creation if order matters, but AutoMigrate usually handles it.
	})
	if err != nil {
		log.Printf("failed to open db for %s: %v", driverName, err)
		return
	}

	if err = db.AutoMigrate(models...); err != nil {
		log.Printf("failed to auto migrate %s: %v", driverName, err)
	}

	// Write SQL to file
	if len(sqlLogger.SQLs) > 0 {
		dir := filepath.Dir(outputPath)
		if err = os.MkdirAll(dir, 0755); err != nil { //nolint:gosec // directory permission
			log.Printf("failed to create dir %s: %v", dir, err)
			return
		}

		content := strings.Join(sqlLogger.SQLs, "\n")
		if err = os.WriteFile(outputPath, []byte(content), 0644); err != nil { //nolint:gosec // file permission
			log.Printf("failed to write schema to %s: %v", outputPath, err)
		} else {
			log.Printf("Successfully generated schema for %s at %s", driverName, outputPath)
		}
	} else {
		log.Printf("No SQL generated for %s", driverName)
	}
}
