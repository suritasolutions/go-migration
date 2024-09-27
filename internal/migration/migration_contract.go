package migration

type Migration interface {
	Migrate()
	Rollback()
}
