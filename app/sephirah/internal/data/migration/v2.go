package migration

type v2 struct {
	v1
}

func (v *v2) migrate() error {
	return v.migrateWrapper(2, v.v1.migrate, func() error {
		type User struct {
			Model
			Username string
			Password string
			Status   string
			Type     string
		}
		if err := v.tx.Migrator().CreateTable(new(User)); err != nil {
			return err
		}
		return nil
	})
}
