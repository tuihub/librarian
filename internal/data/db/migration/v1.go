package migration

import (
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/model"

	"gorm.io/gorm"
)

type v1 struct {
	useTxForMigration *gorm.DB
	tx                *gorm.DB
	targetVersion     *uint
	curVersion        *uint
}

func (v *v1) init(db *gorm.DB) {
	v.useTxForMigration = db
}

func (v *v1) migrate() error {
	version := new(Version)
	version.Version = 1
	if !v.tx.Migrator().HasTable(version) {
		err := v.tx.Migrator().CreateTable(version)
		if err != nil {
			return err
		}
	}
	err := v.tx.Create(version).Error
	if err != nil {
		return err
	}
	v.curVersion = &version.Version
	return nil
}

func (v *v1) migrateWrapper(migrateVersion uint, doPreMigration func() error, doMigration func() error) error {
	var err error
	if v.targetVersion == nil {
		v.targetVersion = &migrateVersion
		v.tx = v.useTxForMigration.Begin()
		defer func() {
			if e := recover(); e != nil {
				v.tx.Rollback()
				panic(e)
			}
			v.targetVersion = nil
			v.curVersion = nil
			v.tx = nil
			if err != nil {
				v.tx.Rollback()
			} else {
				err = v.tx.Commit().Error
			}
		}()
	}

	if v.curVersion == nil {
		version := new(Version)
		// check if the version table exists
		if !v.tx.Migrator().HasTable(version) {
			var cv uint = 0
			v.curVersion = &cv
		} else {
			err = v.tx.Order("version desc").First(version).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			v.curVersion = &version.Version
		}
	}

	if *v.curVersion > migrateVersion {
		return errors.New("database version is newer than the migration version")
	}

	if *v.curVersion == migrateVersion {
		return err
	}

	if *v.curVersion < migrateVersion-1 {
		err = doPreMigration()
		if err != nil {
			return err
		}
	}

	if *v.curVersion != migrateVersion-1 {
		return errors.New("invalid migration sequence")
	}

	err = doMigration()
	if err != nil {
		return err
	}
	err = v.tx.Create(&Version{Version: migrateVersion}).Error
	if err != nil {
		return err
	}
	v.curVersion = &migrateVersion

	return err
}

type Version struct {
	Version   uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Model struct {
	ID        model.InternalID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
