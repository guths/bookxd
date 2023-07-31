package migrations

import "database/sql"

func init() {
	migrator.AddMigrations(&Migration{
		Version: "20230730184904",
		Up:      mig_20230730184904_init_schema_up,
		Down:    mig_20230730184904_init_schema_down,
	})
}

func mig_20230730184904_init_schema_up(tx *sql.Tx) error {
	return nil
}

func mig_20230730184904_init_schema_down(tx *sql.Tx) error {
	return nil
}
