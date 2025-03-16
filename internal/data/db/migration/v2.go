package migration

import "time"

type v2 struct {
	v1
}

func (v *v2) migrate() error {
	return v.migrateWrapper(2, v.v1.migrate, func() error {
		type User struct {
			_Model
			Username string
			Password string
			Status   string
			Type     string
		}
		if err := v.tx.Migrator().CreateTable(new(User)); err != nil {
			return err
		}
		type Session struct {
			_Model
			UserID   _ID
			DeviceID _ID
			Token    string
			ExpireAt time.Time
		}
		if err := v.tx.Migrator().CreateTable(new(Session)); err != nil {
			return err
		}
		type Device struct {
			_Model
			UserID                  _ID
			ClientLocalID           string
			Name                    string
			SystemType              string
			SystemVersion           string
			ClientName              string
			ClientSourceCodeAddress string
			ClientVersion           string
		}
		if err := v.tx.Migrator().CreateTable(new(Device)); err != nil {
			return err
		}
		return nil
	})
}
